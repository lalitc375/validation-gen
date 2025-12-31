# Validation: PodDisruptionBudget

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | PodDisruptionBudget name is required and must be a DNS subdomain. |
| `spec.minAvailable` | `*intstr.IntOrString` | `+k8s:optional` | Optional minimum number of pods that must be available after eviction. Mutually exclusive with `maxUnavailable`. |
| `spec.maxUnavailable` | `*intstr.IntOrString` | `+k8s:optional` | Optional maximum number of pods that can be unavailable after eviction. Mutually exclusive with `minAvailable`. |
| `spec.selector` | `*metav1.LabelSelector` | `+k8s:optional` | Optional label query over pods whose evictions are managed by the disruption budget. |
| `spec.unhealthyPodEvictionPolicy` | `*UnhealthyPodEvictionPolicyType` | `+k8s:optional`<br>`+k8s:enum=["IfHealthyBudget", "AlwaysAllow"]` | Optional policy for when unhealthy pods should be considered for eviction. Defaults to `IfHealthyBudget`. |

### Validation Rules

- `minAvailable` and `maxUnavailable` cannot be both specified.
- If specified as a percentage, the value must be between 0% and 100% inclusive.
- If specified as an integer, the value must be non-negative.
