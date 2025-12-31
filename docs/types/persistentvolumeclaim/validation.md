# Validation: PersistentVolumeClaim

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | PersistentVolumeClaim name is required and must be a DNS subdomain. |
| `spec.accessModes` | `[]PersistentVolumeAccessMode` | `+k8s:required`<br>`+k8s:minItems=1`<br>`+k8s:eachVal=+k8s:enum=["ReadWriteOnce", "ReadOnlyMany", "ReadWriteMany", "ReadWriteOncePod"]` | Mandatory ways the volume can be mounted. `ReadWriteOncePod` cannot be combined with other modes. |
| `spec.resources.requests` | `ResourceList` | `+k8s:required` | Mandatory minimum resources required. `storage` is required and must be a positive quantity. |
| `spec.selector` | `*metav1.LabelSelector` | `+k8s:optional` | Optional label query over volumes to consider for binding. |
| `spec.volumeName` | `string` | `+k8s:optional`<br>`+k8s:immutable` | Optional binding reference to a PersistentVolume. Immutable after it is set. |
| `spec.storageClassName` | `*string` | `+k8s:optional` | Optional name of the StorageClass required by the claim. |
| `spec.volumeMode` | `*PersistentVolumeMode` | `+k8s:optional`<br>`+k8s:enum=["Filesystem", "Block"]` | Optional volume mode required by the claim. Defaults to `Filesystem`. |
| `spec.dataSource` | `*TypedLocalObjectReference` | `+k8s:optional` | Optional object from which to populate the volume (e.g., a VolumeSnapshot). |
| `spec.dataSourceRef` | `*TypedObjectReference` | `+k8s:optional` | Optional object from which to populate the volume. More flexible than `dataSource`. |
| `spec.volumeAttributesClassName` | `*string` | `+k8s:optional` | Optional name of the VolumeAttributesClass used by this claim. |
