# Validation: PersistentVolumeClaim

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name`<br>`+k8s:subfield(name=labels)=+k8s:eachKey=+k8s:format=k8s-label-key`<br>`+k8s:subfield(name=annotations)=+k8s:eachKey=+k8s:format=k8s-label-key` | PersistentVolumeClaim name is required and must be a DNS subdomain. |
| `spec.accessModes` | `[]PersistentVolumeAccessMode` | `+k8s:required` | Mandatory ways the volume can be mounted. `ReadWriteOncePod` cannot be combined with other modes. |
| `spec.resources.requests` | `ResourceList` | `+k8s:required` | Mandatory minimum resources required. `storage` is required and must be a positive quantity. |
| `spec.selector` | `*metav1.LabelSelector` | `+k8s:optional` | Optional label selector. |
| `spec.volumeName` | `string` | `+k8s:optional`<br>`+k8s:immutable` | Optional binding reference to a PersistentVolume. Immutable after it is set. |
| `spec.storageClassName` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional name of the StorageClass. |
| `spec.volumeMode` | `*PersistentVolumeMode` | `+k8s:optional`<br> | Optional volume mode. Defaults to `Filesystem`. |
| `spec.dataSource` | `*TypedLocalObjectReference` | `+k8s:optional` | Optional data source. |
| `spec.dataSourceRef` | `*TypedObjectReference` | `+k8s:optional` | Optional data source reference. |
| `spec.volumeAttributesClassName` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional VolumeAttributesClass name. |
| `status.allocatedResources` | `ResourceList` |  | Resources allocated for the claim. |
| `status.allocatedResourceStatuses` | `map[ResourceName]ClaimResourceStatus` |  | Status of resource allocation. |

### TypedLocalObjectReference

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `apiGroup` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional API group. |
| `kind` | `string` | `+k8s:required` | Mandatory kind. |
| `name` | `string` | `+k8s:required` | Mandatory name. |

### TypedObjectReference

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `apiGroup` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional API group. |
| `kind` | `string` | `+k8s:required` | Mandatory kind. |
| `name` | `string` | `+k8s:required` | Mandatory name. |
| `namespace` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional namespace. |

## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `ClaimResourceStatus` | `+k8s:enum` |
| `PersistentVolumeAccessMode` | `+k8s:enum` |
| `PersistentVolumeMode` | `+k8s:enum` |