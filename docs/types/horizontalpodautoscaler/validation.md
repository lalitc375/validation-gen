# Validation: HorizontalPodAutoscaler

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | HPA name is required and must be a DNS subdomain. |
| `spec.scaleTargetRef` | `CrossVersionObjectReference` | `+k8s:required` | Mandatory reference to the resource to scale. |
| `spec.scaleTargetRef.kind` | `string` | `+k8s:required` | Mandatory kind of the target resource. |
| `spec.scaleTargetRef.name` | `string` | `+k8s:required` | Mandatory name of the target resource. |
| `spec.minReplicas` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=1` | Optional lower limit for replicas. Defaults to 1. |
| `spec.maxReplicas` | `int32` | `+k8s:required`<br>`+k8s:minimum=1` | Mandatory upper limit for replicas. Must be >= `minReplicas`. |
| `spec.metrics` | `[]MetricSpec` | `+k8s:optional` | Optional list of metrics to use for calculation. |
| `spec.behavior.scaleUp.stabilizationWindowSeconds` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0`<br>`+k8s:maximum=3600` | Optional window to consider past recommendations for scaling up. |
| `spec.behavior.scaleDown.stabilizationWindowSeconds` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0`<br>`+k8s:maximum=3600` | Optional window to consider past recommendations for scaling down. |
| `spec.behavior.scaleUp.selectPolicy` | `*ScalingPolicySelect` | `+k8s:optional`<br> | Optional policy selection for scaling up. |
| `spec.behavior.scaleDown.selectPolicy` | `*ScalingPolicySelect` | `+k8s:optional`<br> | Optional policy selection for scaling down. |
| `spec.behavior.scaleUp.policies` | `[]HPAScalingPolicy` | `+k8s:optional`<br>`+k8s:minItems=1` | List of scaling policies for scaling up. |
| `spec.behavior.scaleDown.policies` | `[]HPAScalingPolicy` | `+k8s:optional`<br>`+k8s:minItems=1` | List of scaling policies for scaling down. |
| `spec.behavior.scaleUp.policies[].type` | `HPAScalingPolicyType` | `+k8s:required`<br> | Mandatory policy type. |
| `spec.behavior.scaleUp.policies[].value` | `int32` | `+k8s:required`<br>`+k8s:minimum=1` | Mandatory policy value. |
| `spec.behavior.scaleUp.policies[].periodSeconds` | `int32` | `+k8s:required`<br>`+k8s:minimum=1`<br>`+k8s:maximum=1800` | Mandatory policy period in seconds. |


## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `HPAScalingPolicyType` | `+k8s:enum` |
| `ScalingPolicySelect` | `+k8s:enum` |
