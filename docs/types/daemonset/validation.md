# Validation: DaemonSet

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | DaemonSet name is required and must be a DNS subdomain. |
| `spec.selector` | `*metav1.LabelSelector` | `+k8s:required`<br>`+k8s:immutable` | Mandatory label selector. Immutable after creation. Must match template labels. |
| `spec.template` | `corev1.PodTemplateSpec` | `+k8s:required` | Mandatory pod template. `restartPolicy` must be `Always`. `activeDeadlineSeconds` is forbidden. |
| `spec.updateStrategy.type` | `DaemonSetUpdateStrategyType` | `+k8s:optional`<br> | Strategy type for updates. Defaults to `RollingUpdate`. |
| `spec.updateStrategy.rollingUpdate.maxUnavailable` | `intstr.IntOrString` | `+k8s:optional` | Max pods that can be unavailable. Defaults to 1. Exactly one of `maxSurge` or `maxUnavailable` must be non-zero. |
| `spec.updateStrategy.rollingUpdate.maxSurge` | `intstr.IntOrString` | `+k8s:optional` | Max pods that can be created above desired count. Defaults to 0. |
| `spec.minReadySeconds` | `int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Min seconds for which a pod should be ready. |
| `spec.revisionHistoryLimit` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Max old history to retain. Defaults to 10. |
| `status.conditions` | `[]DaemonSetCondition` | `+k8s:optional`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=type` | Represents the latest available observations of a DaemonSet's current state. |


## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `DaemonSetUpdateStrategyType` | `+k8s:enum` |
