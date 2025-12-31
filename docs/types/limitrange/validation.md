# Validation: LimitRange

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | LimitRange name is required and must be a DNS subdomain. |
| `spec.limits` | `[]LimitRangeItem` | `+k8s:required` | Mandatory list of LimitRangeItem objects that are enforced. |
| `spec.limits[].type` | `LimitType` | `+k8s:required`<br>`+k8s:enum=["Pod", "Container", "PersistentVolumeClaim"]` | Mandatory type of resource that this limit applies to. |
| `spec.limits[].max` | `ResourceList` | `+k8s:optional` | Optional maximum usage constraints on this kind by resource name. |
| `spec.limits[].min` | `ResourceList` | `+k8s:optional` | Optional minimum usage constraints on this kind by resource name. |
| `spec.limits[].default` | `ResourceList` | `+k8s:optional` | Optional default resource requirement limit value. Forbidden if `type` is `Pod`. |
| `spec.limits[].defaultRequest` | `ResourceList` | `+k8s:optional` | Optional default resource requirement request value. Forbidden if `type` is `Pod`. |
| `spec.limits[].maxLimitRequestRatio` | `ResourceList` | `+k8s:optional` | Optional maximum burst value for the named resource. Each ratio must be >= 1. |

### Validation Rules

- `min` must be less than or equal to `max` if both are specified.
- `defaultRequest` must be greater than or equal to `min` if `min` is specified.
- `defaultRequest` must be less than or equal to `max` if `max` is specified.
- `defaultRequest` must be less than or equal to `default` if both are specified.
- `default` must be greater than or equal to `min` if `min` is specified.
- `default` must be less than or equal to `max` if `max` is specified.
- For `PersistentVolumeClaim`, either `min` or `max` storage value is required.
