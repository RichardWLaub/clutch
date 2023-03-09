import React, { useState } from "react";
import type { clutch as IClutch } from "@clutch-sh/api";
import {
  Button,
  ButtonGroup,
  client,
  Confirmation,
  MetadataTable,
  Resolver,
  Select,
  useWizardContext,
} from "@clutch-sh/core";
import { useDataLayout } from "@clutch-sh/data-layout";
import type { WizardChild } from "@clutch-sh/wizard";
import { Wizard, WizardStep } from "@clutch-sh/wizard";

import type { ConfirmChild, ResolverChild, WorkflowProps } from ".";

const DeploymentIdentifier: React.FC<ResolverChild> = ({ resolverType }) => {
  const { onSubmit } = useWizardContext();
  const resourceData = useDataLayout("resourceData");
  const resolverInput = useDataLayout("resolverInput");

  const onResolve = ({ results, input }) => {
    // Decide how to process results.
    resourceData.assign(results[0]);
    resolverInput.assign(input);
    onSubmit();
  };

  return <Resolver type={resolverType} searchLimit={1} onResolve={onResolve} />;
};

const DeploymentDetails: React.FC<WizardChild> = () => {
  const { onSubmit, onBack } = useWizardContext();
  const resourceData = useDataLayout("resourceData");
  const deployment = resourceData.displayValue() as IClutch.k8s.v1.Deployment;
  const update = (key: string, value: boolean) => {
    resourceData.updateData(key, value);
  };
  const [containerName, setContainerName] = useState(
    deployment.deploymentSpec.template.spec.containers[0].name
  );
  const index = deployment.deploymentSpec.template.spec.containers.findIndex(
    container => container.name === containerName
  );
  return (
    <WizardStep error={resourceData.error} isLoading={resourceData.isLoading}>
      <strong>Deployment Details</strong>
      <MetadataTable
        onUpdate={update}
        data={[
          { name: "Name", value: deployment.name },
          { name: "Namespace", value: deployment.namespace },
          { name: "Replicas", value: deployment.deploymentStatus.replicas },
          {
            name: "Container Name",
            value: (
              <Select
                defaultOption=""
                helperText=""
                label="Container Name"
                name="containerName"
                onChange={value => {
                  setContainerName(value);
                  resourceData.updateData("containerName", value);
                }}
                options={deployment.deploymentSpec.template.spec.containers.map(container => {
                  return { label: container.name };
                })}
              />
            ),
          },
          {
            name: "CPU Limit",
            value: deployment.deploymentSpec.template.spec.containers.find(
              container => container.name === containerName
            ).resources.limits.cpu,
            textFieldLabels: {
              disabledField: "Current Limit",
              updatedField: "New limit",
            },
            input: {
              type: "string",
              key: `deploymentSpec.template.spec.containers[${index}].resources.limits.cpu`,
            },
          },
          {
            name: "CPU Request",
            value: deployment.deploymentSpec.template.spec.containers.find(
              container => container.name === containerName
            ).resources.requests.cpu,
            textFieldLabels: {
              disabledField: "Current Request",
              updatedField: "New Request",
            },
            input: {
              type: "string",
              key: `deploymentSpec.template.spec.containers[${index}].resources.requests.cpu`,
            },
          },
          {
            name: "Memory Limit",
            value: deployment.deploymentSpec.template.spec.containers.find(
              container => container.name === containerName
            ).resources.limits.memory,
            textFieldLabels: {
              disabledField: "Current Limit",
              updatedField: "New limit",
            },
            input: {
              type: "string",
              key: `deploymentSpec.template.spec.containers[${index}].resources.limits.memory`,
            },
          },
          {
            name: "Memory Request",
            value: deployment.deploymentSpec.template.spec.containers.find(
              container => container.name === containerName
            ).resources.requests.memory,
            textFieldLabels: {
              disabledField: "Current Request",
              updatedField: "New Request",
            },
            input: {
              type: "string",
              key: `deploymentSpec.template.spec.containers[${index}].resources.requests.memory`,
            },
          },
        ]}
      />
      <ButtonGroup>
        <Button text="Back" variant="neutral" onClick={() => onBack()} />
        <Button text="Update" variant="destructive" onClick={onSubmit} />
      </ButtonGroup>
    </WizardStep>
  );
};

const Confirm: React.FC<ConfirmChild> = () => {
  const deployment = useDataLayout("resourceData").displayValue() as IClutch.k8s.v1.Deployment;
  const updateData = useDataLayout("updateData");
  return (
    <WizardStep error={updateData.error} isLoading={updateData.isLoading}>
      <Confirmation action="Update" />
      <MetadataTable
        data={[
          { name: "Name", value: deployment.name },
          { name: "Namespace", value: deployment.namespace },
          { name: "Cluster", value: deployment.cluster },
        ]}
      />
    </WizardStep>
  );
};

const ScaleResources: React.FC<WorkflowProps> = ({ heading, resolverType }) => {
  const dataLayout = {
    resolverInput: {},
    resourceData: {},
    updateData: {
      deps: ["resourceData", "resolverInput"],
      hydrator: (
        resourceData: {
          cluster: string;
          containerName: string;
          deploymentSpec: IClutch.k8s.v1.Deployment.DeploymentSpec;
          name: string;
          namespace: string;
        },
        resolverInput: { clientset: string }
      ) => {
        const clientset = resolverInput.clientset ?? "undefined";
        const limits: { [key: string]: string } = {
          cpu: resourceData.deploymentSpec.template.spec.containers.find(
            container => container.name === resourceData.containerName
          ).resources.limits.cpu,
          memory: resourceData.deploymentSpec.template.spec.containers.find(
            container => container.name === resourceData.containerName
          ).resources.limits.memory,
        };
        const requests: { [key: string]: string } = {
          cpu: resourceData.deploymentSpec.template.spec.containers.find(
            container => container.name === resourceData.containerName
          ).resources.requests.cpu,
          memory: resourceData.deploymentSpec.template.spec.containers.find(
            container => container.name === resourceData.containerName
          ).resources.requests.memory,
        };
        return client.post("/v1/k8s/updateDeployment", {
          clientset,
          cluster: resourceData.cluster,
          namespace: resourceData.namespace,
          name: resourceData.name,
          fields: {
            containerResources: [
              {
                containerName: resourceData.containerName,
                resources: {
                  limits,
                  requests,
                },
              },
            ],
          },
        } as IClutch.k8s.v1.UpdateDeploymentRequest);
      },
    },
  };

  return (
    <Wizard dataLayout={dataLayout} heading={heading}>
      <DeploymentIdentifier name="Lookup" resolverType={resolverType} />
      <DeploymentDetails name="Modify" />
      <Confirm name="Confirmation" />
    </Wizard>
  );
};

export default ScaleResources;
