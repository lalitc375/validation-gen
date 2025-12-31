# Validation: PersistentVolume

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | PersistentVolume name is required and must be a DNS subdomain. |
| `spec.capacity` | `ResourceList` | `+k8s:required` | Mandatory map of resource names to quantities. Only `storage` is typically allowed. |
| `spec.accessModes` | `[]PersistentVolumeAccessMode` | `+k8s:required`<br>`+k8s:minItems=1`<br>`+k8s:eachVal=+k8s:enum=["ReadWriteOnce", "ReadOnlyMany", "ReadWriteMany", "ReadWriteOncePod"]` | Mandatory ways the volume can be mounted. `ReadWriteOncePod` cannot be combined with other modes. |
| `spec.persistentVolumeReclaimPolicy` | `PersistentVolumeReclaimPolicy` | `+k8s:optional`<br>`+k8s:enum=["Recycle", "Delete", "Retain"]` | Optional policy for maintenance after release. Defaults to `Retain`. |
| `spec.storageClassName` | `string` | `+k8s:optional` | Optional name of the StorageClass to which this volume belongs. |
| `spec.mountOptions` | `[]string` | `+k8s:optional` | Optional list of mount options. |
| `spec.volumeMode` | `*PersistentVolumeMode` | `+k8s:optional`<br>`+k8s:enum=["Filesystem", "Block"]` | Optional volume mode. Defaults to `Filesystem`. |
| `spec.nodeAffinity` | `*VolumeNodeAffinity` | `+k8s:optional` | Optional constraints that limit what nodes this volume can be accessed from. |
| `spec.claimRef` | `*ObjectReference` | `+k8s:optional` | Optional reference to the PersistentVolumeClaim. |

### Volume Sources

Exactly one of the following must be specified:
- `awsElasticBlockStore`, `azureDisk`, `azureFile`, `cephfs`, `cinder`, `csi`, `fc`, `flexVolume`, `flocker`, `gcePersistentDisk`, `glusterfs`, `hostPath`, `iscsi`, `local`, `nfs`, `photonPersistentDisk`, `portworxVolume`, `quobyte`, `rbd`, `scaleIO`, `storageos`, `vsphereVolume`.
