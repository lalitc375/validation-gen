# Validation: RangeAllocation

`RangeAllocation` is used to track the allocation of a range of resources, such as IP addresses or ports.

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required` | Name of the range allocation. |
| `range` | `string` | `+k8s:required` | Mandatory string representing a unique label for a range of resources (e.g., a CIDR or port range). |
| `data` | `[]byte` | `+k8s:required` | Mandatory byte array representing the serialized state of the range allocation. |
