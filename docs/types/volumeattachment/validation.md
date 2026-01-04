# Validation: VolumeAttachment

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | VolumeAttachment name is required and must be a DNS subdomain. |
| `spec.attacher` | `string` | `+k8s:required`<br>`+k8s:immutable` | Mandatory name of the CSI volume driver that must handle this request. Immutable after creation. |
| `spec.source` | `VolumeAttachmentSource` | `+k8s:required`<br>`+k8s:immutable` | Mandatory volume that should be attached. Immutable after creation. |
| `spec.source.persistentVolumeName` | `*string` | `+k8s:unionMember`<br/>`+k8s:optional` | Optional name of the PersistentVolume being attached. |
| `spec.source.inlineVolumeSpec` | `*PersistentVolumeSpec` | `+k8s:unionMember`<br/>`+k8s:optional` | Optional inline specification of a volume. |
| `spec.nodeName` | `string` | `+k8s:required`<br>`+k8s:format=k8s-long-name`<br>`+k8s:immutable` | Mandatory name of the node that the volume should be attached to. Immutable after creation. |
