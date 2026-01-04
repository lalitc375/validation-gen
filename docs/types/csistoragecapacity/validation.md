# Validation: CSIStorageCapacity

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | CSIStorageCapacity name is required and must be a DNS subdomain. |
| `nodeTopology` | `*metav1.LabelSelector` | `+k8s:optional`<br>`+k8s:immutable` | Defines nodes with access to the storage. Immutable after creation. |
| `storageClassName` | `string` | `+k8s:required`<br>`+k8s:format=k8s-long-name`<br>`+k8s:immutable` | Mandatory name of the StorageClass. Immutable after creation. |
| `capacity` | `*resource.Quantity` | `+k8s:optional`<br>`+k8s:minimum=0` | Available capacity in bytes. Must be non-negative if set. |
| `maximumVolumeSize` | `*resource.Quantity` | `+k8s:optional`<br>`+k8s:minimum=0` | Largest size for a single volume. Must be non-negative if set. |
