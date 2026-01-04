# Validation: Deployment

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Deployment name is required and must be a DNS subdomain. |
| `spec.replicas` | `int32` | `+k8s:required`<br>`+k8s:minimum=0` | Mandatory number of desired pods. |
| `spec.selector` | `*metav1.LabelSelector` | `+k8s:required`<br>`+k8s:immutable` | Mandatory label selector. Immutable after creation. Must match template labels. |
| `spec.template` | `corev1.PodTemplateSpec` | `+k8s:required` | Mandatory pod template. `restartPolicy` must be `Always`. |
| `spec.strategy.type` | `DeploymentStrategyType` | `+k8s:optional`<br> | Strategy type for updates. Defaults to `RollingUpdate`. |
| `spec.strategy.rollingUpdate.maxUnavailable` | `*intstr.IntOrString` | `+k8s:optional` | Max pods that can be unavailable during update. |
| `spec.strategy.rollingUpdate.maxSurge` | `*intstr.IntOrString` | `+k8s:optional` | Max pods that can be scheduled above desired count. |
| `spec.minReadySeconds` | `int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Min seconds for which a pod should be ready. |
| `spec.revisionHistoryLimit` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Max old history to retain. Defaults to 10. |
| `spec.paused` | `bool` | `+k8s:optional` | Indicates if the deployment is paused. |
| `spec.progressDeadlineSeconds` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Max time for deployment to make progress. Must be > `minReadySeconds` if both set. |
| `status.conditions` | `[]DeploymentCondition` | `+k8s:optional`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=type` | Represents the latest available observations of a deployment's current state. |


## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `DeploymentStrategyType` | `+k8s:enum` |
