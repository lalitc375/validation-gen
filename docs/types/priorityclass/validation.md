# Validation: PriorityClass

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | PriorityClass name is required and must be a DNS subdomain. |
| `value` | `int32` | `+k8s:required`<br>`+k8s:minimum=-2147483648`<br>`+k8s:maximum=1000000000`<br>`+k8s:immutable` | Mandatory integer value of the priority. User-defined priority classes must have values <= 1 billion. Immutable after creation. |
| `globalDefault` | `bool` | `+k8s:optional` | Optional flag indicating if this priority should be used for pods without any priority class. |
| `description` | `string` | `+k8s:optional` | Optional guidelines on when this priority class should be used. |
| `preemptionPolicy` | `*PreemptionPolicy` | `+k8s:optional`<br>`+k8s:enum=["Never", "PreemptLowerPriority"]`<br>`+k8s:immutable` | Optional policy for preempting pods with lower priority. Defaults to `PreemptLowerPriority`. Immutable after creation. |

### Reserved Values

- Priority values greater than 1 billion are reserved for system-critical components.
- Names prefixed with `system-` are reserved for system use.