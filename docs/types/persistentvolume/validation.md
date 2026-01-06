# Validation: PersistentVolume

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name`<br>`+k8s:subfield(name=labels)=+k8s:eachKey=+k8s:format=k8s-label-key`<br>`+k8s:subfield(name=annotations)=+k8s:eachKey=+k8s:format=k8s-label-key` | PersistentVolume name is required and must be a DNS subdomain. |
| `spec.capacity` | `ResourceList` | `+k8s:required` | Mandatory map of resource names to quantities. Only `storage` is typically allowed. |
| `spec.accessModes` | `[]PersistentVolumeAccessMode` | `+k8s:required` | Mandatory ways the volume can be mounted. `ReadWriteOncePod` cannot be combined with other modes. |
| `spec.persistentVolumeReclaimPolicy` | `PersistentVolumeReclaimPolicy` | `+k8s:optional`<br> | Optional policy for maintenance after release. Defaults to `Retain`. |
| `spec.storageClassName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional name of the StorageClass. |
| `spec.mountOptions` | `[]string` | `+k8s:optional`<br>`+k8s:listType=atomic` | Optional list of mount options. |
| `spec.volumeMode` | `*PersistentVolumeMode` | `+k8s:optional`<br> | Optional volume mode. Defaults to `Filesystem`. |
| `spec.nodeAffinity` | `*VolumeNodeAffinity` | `+k8s:optional` | Optional constraints that limit what nodes this volume can be accessed from. |
| `spec.claimRef` | `*ObjectReference` | `+k8s:optional` | Optional reference to the PersistentVolumeClaim. |
| `spec.volumeAttributesClassName` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional VolumeAttributesClass name. |

### VolumeSource

The following union members are available under `spec`.

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `gcePersistentDisk` | `*GCEPersistentDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | GCE Persistent Disk. |
| `awsElasticBlockStore` | `*AWSElasticBlockStoreVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | AWS Elastic Block Store. |
| `hostPath` | `*HostPathVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Host Path. |
| `glusterfs` | `*GlusterfsPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | GlusterFS. |
| `nfs` | `*NFSVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | NFS. |
| `rbd` | `*RBDPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Rados Block Device. |
| `iscsi` | `*ISCSIPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | ISCSI. |
| `cinder` | `*CinderPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Cinder. |
| `cephfs` | `*CephFSPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | CephFS. |
| `fc` | `*FCVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Fibre Channel. |
| `flocker` | `*FlockerVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Flocker. |
| `flexVolume` | `*FlexPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | FlexVolume. |
| `azureFile` | `*AzureFilePersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Azure File. |
| `vsphereVolume` | `*VsphereVirtualDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | vSphere Volume. |
| `quobyte` | `*QuobyteVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Quobyte. |
| `azureDisk` | `*AzureDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Azure Disk. |
| `photonPersistentDisk` | `*PhotonPersistentDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Photon Persistent Disk. |
| `portworxVolume` | `*PortworxVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Portworx Volume. |
| `scaleIO` | `*ScaleIOPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | ScaleIO. |
| `local` | `*LocalVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Local Volume. |
| `storageos` | `*StorageOSPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | StorageOS. |
| `csi` | `*CSIPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | CSI. |

### GlusterfsPersistentVolumeSource

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `endpoints` | `string` | `+k8s:required`<br>`+k8s:format=k8s-long-name` | Mandatory endpoints name. |
| `endpointsNamespace` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional endpoints namespace. |

### QuobyteVolumeSource

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `tenant` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-uuid` | Optional tenant UUID. |

### StorageOSPersistentVolumeSource

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `volumeName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional volume name. |
| `volumeNamespace` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional volume namespace. |

### CSIPersistentVolumeSource

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `driver` | `string` | `+k8s:required`<br>`+k8s:format=k8s-long-name` | Mandatory driver name. |

### FlexPersistentVolumeSource

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `options` | `map[string]string` | `+k8s:optional` | Optional extra command options. |

## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `PersistentVolumeAccessMode` | `+k8s:enum` |
| `PersistentVolumeMode` | `+k8s:enum` |
| `PersistentVolumeReclaimPolicy` | `+k8s:enum` |