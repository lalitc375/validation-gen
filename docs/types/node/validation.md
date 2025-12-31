# Validation: Node

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Node name is required and must be a DNS subdomain. |
| `spec.podCIDR` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-cidr`<br>`+k8s:immutable` | Optional pod IP range. Immutable after it is set. |
| `spec.podCIDRs` | `[]string` | `+k8s:optional`<br>`+k8s:eachVal=+k8s:format=k8s-cidr`<br>`+k8s:immutable` | Optional pod IP ranges. Immutable after they are set. |
| `spec.providerID` | `string` | `+k8s:optional`<br>`+k8s:immutable` | Optional ID assigned by the cloud provider. Immutable after it is set. |
| `spec.unschedulable` | `bool` | `+k8s:optional` | Optional flag to disable pod scheduling on the node. |
| `spec.taints` | `[]Taint` | `+k8s:optional` | Optional list of taints that have the "NoSchedule", "PreferNoSchedule" or "NoExecute" effect. |

### NodeStatus

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `capacity` | `ResourceList` | `+k8s:optional` | Total resources of the node. Each value must be non-negative. |
| `allocatable` | `ResourceList` | `+k8s:optional` | Resources available for scheduling. Each value must be non-negative. |
| `addresses` | `[]NodeAddress` | `+k8s:optional` | List of addresses for the node. Must not contain duplicates. |
