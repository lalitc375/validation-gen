# Validation: PersistentVolumeClaim

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name`<br>`+k8s:subfield(name=labels)=+k8s:eachKey=+k8s:format=k8s-label-key`<br>`+k8s:subfield(name=annotations)=+k8s:eachKey=+k8s:format=k8s-annotation-key` | PersistentVolumeClaim name is required and must be a DNS subdomain. |
| `spec.accessModes` | `[]PersistentVolumeAccessMode` | `+k8s:required`<br>`+k8s:minItems=1`<br>`+k8s:eachVal=` | Mandatory ways the volume can be mounted. `ReadWriteOncePod` cannot be combined with other modes. |
| `spec.resources.requests` | `ResourceList` | `+k8s:required`<br>`+k8s:eachKey=+k8s:format=k8s-container-resource-name` | Mandatory minimum resources required. `storage` is required and must be a positive quantity. |
| `spec.selector` | `*metav1.LabelSelector` | `+k8s:optional` | Optional label query over volumes to consider for binding. |
| `spec.volumeName` | `string` | `+k8s:optional`<br>`+k8s:immutable` | Optional binding reference to a PersistentVolume. Immutable after it is set. |
| `spec.storageClassName` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional name of the StorageClass required by the claim. |
| `spec.volumeMode` | `*PersistentVolumeMode` | `+k8s:optional`<br> | Optional volume mode required by the claim. Defaults to `Filesystem`. |
| `spec.dataSource` | `*TypedLocalObjectReference` | `+k8s:optional` | Optional object from which to populate the volume (e.g., a VolumeSnapshot). |
| `spec.dataSource.apiGroup` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional API group. |
| `spec.dataSource.kind` | `string` | `+k8s:required` | Mandatory kind. |
| `spec.dataSource.name` | `string` | `+k8s:required` | Mandatory name. |
| `spec.dataSourceRef` | `*TypedObjectReference` | `+k8s:optional` | Optional object from which to populate the volume. More flexible than `dataSource`. |
| `spec.dataSourceRef.apiGroup` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional API group. |
| `spec.dataSourceRef.kind` | `string` | `+k8s:required` | Mandatory kind. |
| `spec.dataSourceRef.name` | `string` | `+k8s:required` | Mandatory name. |
| `spec.dataSourceRef.namespace` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional namespace. |
| `spec.volumeAttributesClassName` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional name of the VolumeAttributesClass used by the claim. |
| `status.allocatedResources` | `ResourceList` | `+k8s:eachKey=+k8s:format=k8s-pvc-resource-key` | Resources allocated for the claim. |
| `status.allocatedResourceStatuses` | `map[ResourceName]ClaimResourceStatus` | `+k8s:eachKey=+k8s:format=k8s-pvc-resource-key` | Status of resource allocation. |


## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `ClaimResourceStatus` | `+k8s:enum` |
| `PersistentVolumeAccessMode` | `+k8s:enum` |
| `PersistentVolumeMode` | `+k8s:enum` |
