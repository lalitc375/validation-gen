# Validation: StorageClass

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name`<br>`+k8s:subfield(name=labels)=+k8s:eachKey=+k8s:format=k8s-label-key`<br>`+k8s:subfield(name=annotations)=+k8s:eachKey=+k8s:format=k8s-annotation-key` | StorageClass name is required and must be a DNS subdomain. |
| `provisioner` | `string` | `+k8s:required`<br>`+k8s:format=k8s-label-key`<br>`+k8s:immutable` | Mandatory name of the driver expected to handle this StorageClass. Immutable after creation. |
| `parameters` | `map[string]string` | `+k8s:optional`<br>`+k8s:eachKey=+k8s:minLength=1`<br>`+k8s:immutable` | Optional parameters for the provisioner. Maximum 512 parameters, cumulative size <= 256 KiB. Immutable after creation. |
| `reclaimPolicy` | `*PersistentVolumeReclaimPolicy` | `+k8s:optional`<br>`+k8s:immutable` | Optional policy for maintenance after release. Defaults to `Delete`. Immutable after creation. |
| `mountOptions` | `[]string` | `+k8s:optional`<br>`+k8s:listType=atomic` | Optional list of mount options. |
| `allowVolumeExpansion` | `*bool` | `+k8s:optional` | Optional flag indicating if the storage class allows volume expansion. |
| `volumeBindingMode` | `*VolumeBindingMode` | `+k8s:required`<br>`+k8s:immutable` | Mandatory mode indicating how PVCs should be provisioned and bound. Immutable after creation. |
| `allowedTopologies` | `[]TopologySelectorTerm` | `+k8s:optional`<br>`+k8s:listType=atomic` | Optional list of node topologies where volumes can be dynamically provisioned. |


## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `PersistentVolumeReclaimPolicy` | `+k8s:enum` |
| `VolumeBindingMode` | `+k8s:enum` |
