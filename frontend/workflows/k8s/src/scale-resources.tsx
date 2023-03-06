import React from "react";
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
                options={deployment.deploymentSpec.template.spec.containers.map(container => {
                  return { label: container.name };
                })}
              />
            ),
          },
          {
            name: "CPU Limit",
            value: deployment.deploymentSpec.template.spec.containers[0].resources.limits.cpu,
            textFieldLabels: {
              disabledField: "Current Limit",
              updatedField: "New limit",
            },
            input: {
              type: "string",
              key: "deploymentSpec.template.spec.containers[0].resources.limits.cpu",
            },
          },
          {
            name: "CPU Request",
            value: deployment.deploymentSpec.template.spec.containers[0].resources.requests.cpu,
            textFieldLabels: {
              disabledField: "Current Request",
              updatedField: "New Request",
            },
            input: {
              type: "string",
              key: "deploymentSpec.template.spec.containers[0].resources.requests.cpu",
            },
          },
          {
            name: "Memory Limit",
            value: deployment.deploymentSpec.template.spec.containers[0].resources.limits.memory,
            textFieldLabels: {
              disabledField: "Current Limit",
              updatedField: "New limit",
            },
            input: {
              type: "string",
              key: "deploymentSpec.template.spec.containers[0].resources.limits.memory",
            },
          },
          {
            name: "Memory Request",
            value: deployment.deploymentSpec.template.spec.containers[0].resources.requests.memory,
            textFieldLabels: {
              disabledField: "Current Request",
              updatedField: "New Request",
            },
            input: {
              type: "string",
              key: "deploymentSpec.template.spec.containers[0].resources.requests.memory",
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
  const node = useDataLayout("resourceData").displayValue() as IClutch.k8s.v1.Node;
  const updateData = useDataLayout("updateData");
  return (
    <WizardStep error={updateData.error} isLoading={updateData.isLoading}>
      <Confirmation action="Update" />
      <MetadataTable
        data={[
          { name: "Name", value: node.name },
          { name: "Cluster", value: node.cluster },
          { name: "Unschedulable", value: String(node.unschedulable) },
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
        resourceData: { containerName: string; deployment: IClutch.k8s.v1.Deployment },
        resolverInput: { clientset: string }
      ) => {
        const clientset = resolverInput.clientset ?? "undefined";
        const thename = "crontupisto";
        const limits: { [key: string]: string } = {
          cpu: resourceData.deployment.deploymentSpec.template.spec.containers.find(
            container => container.name === thename
          ).resources.limits.cpu,
          memory: resourceData.deployment.deploymentSpec.template.spec.containers.find(
            container => container.name === thename
          ).resources.limits.memory,
        };
        const requests: { [key: string]: string } = {
          cpu: resourceData.deployment.deploymentSpec.template.spec.containers.find(
            container => container.name === thename
          ).resources.requests.cpu,
          memory: resourceData.deployment.deploymentSpec.template.spec.containers.find(
            container => container.name === thename
          ).resources.requests.memory,
        };
        return client.post("/v1/k8s/updateDeployment", {
          clientset,
          cluster: resourceData.deployment.cluster,
          namespace: resourceData.deployment.namespace,
          name: resourceData.deployment.name,
          fields: {
            containerResources: [
              {
                containerName: "healthcheck-chart",
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
