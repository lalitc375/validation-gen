# Validation: RuntimeClass

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | RuntimeClass name is required and must be a DNS subdomain. |
| `handler` | `string` | `+k8s:required`<br>`+k8s:format=k8s-short-name`<br>`+k8s:immutable` | Mandatory name of the underlying runtime and configuration to use. Must be a valid DNS label. Immutable after creation. |
| `overhead` | `*Overhead` | `+k8s:optional` | Optional resource overhead associated with running a pod for this RuntimeClass. |
| `scheduling` | `*Scheduling` | `+k8s:optional` | Optional scheduling constraints to ensure pods are scheduled to supporting nodes. |

### Scheduling

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `nodeSelector` | `map[string]string` | `+k8s:optional` | Optional label selector that must match labels on supporting nodes. |
| `tolerations` | `[]Toleration` | `+k8s:optional` | Optional list of tolerations to be appended to pods running with this RuntimeClass. |
