# Validation: VolumeAttributesClass

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | VolumeAttributesClass name is required and must be a DNS subdomain. |
| `driverName` | `string` | `+k8s:required`<br>`+k8s:immutable` | Mandatory name of the CSI driver. Immutable after creation. |
| `parameters` | `map[string]string` | `+k8s:required`<br>`+k8s:minItems=1`<br>`+k8s:immutable` | Mandatory volume attributes defined by the CSI driver. Must contain at least one key/value pair. Maximum 512 parameters. Immutable after creation. |
