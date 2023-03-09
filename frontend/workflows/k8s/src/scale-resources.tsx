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
  const deploymentData = useDataLayout("deploymentData");
  const resolverInput = useDataLayout("resolverInput");

  const onResolve = ({ results, input }) => {
    // Decide how to process results.
    deploymentData.assign(results[0]);
    resolverInput.assign(input);
    onSubmit();
  };

  return <Resolver type={resolverType} searchLimit={1} onResolve={onResolve} />;
};

const DeploymentDetails: React.FC<WizardChild> = () => {
  const { onSubmit, onBack } = useWizardContext();
  const deploymentData = useDataLayout("deploymentData");
  const deployment = deploymentData.displayValue() as IClutch.k8s.v1.Deployment;
  const update = (key: string, value: boolean) => {
    deploymentData.updateData(key, value);
  };
  const [containerName, setContainerName] = useState(
    deployment.deploymentSpec.template.spec.containers[0].name
  );
  const index = deployment.deploymentSpec.template.spec.containers.findIndex(
    container => container.name === containerName
  );
  return (
    <WizardStep error={deploymentData.error} isLoading={deploymentData.isLoading}>
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
                  deploymentData.updateData("containerName", value);
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
  const deployment = useDataLayout("deploymentData").displayValue() as IClutch.k8s.v1.Deployment;
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
    deploymentData: {},
    updateData: {
      deps: ["deploymentData", "resolverInput"],
      hydrator: (
        deploymentData: {
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
          cpu: deploymentData.deploymentSpec.template.spec.containers.find(
            container => container.name === deploymentData.containerName
          ).resources.limits.cpu,
          memory: deploymentData.deploymentSpec.template.spec.containers.find(
            container => container.name === deploymentData.containerName
          ).resources.limits.memory,
        };
        const requests: { [key: string]: string } = {
          cpu: deploymentData.deploymentSpec.template.spec.containers.find(
            container => container.name === deploymentData.containerName
          ).resources.requests.cpu,
          memory: deploymentData.deploymentSpec.template.spec.containers.find(
            container => container.name === deploymentData.containerName
          ).resources.requests.memory,
        };
        return client.post("/v1/k8s/updateDeployment", {
          clientset,
          cluster: deploymentData.cluster,
          namespace: deploymentData.namespace,
          name: deploymentData.name,
          fields: {
            containerResources: [
              {
                containerName: deploymentData.containerName,
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
