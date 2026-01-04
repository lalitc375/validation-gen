# Validation: PersistentVolume

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name`<br>`+k8s:subfield(name=labels)=+k8s:eachKey=+k8s:format=k8s-label-key`<br>`+k8s:subfield(name=annotations)=+k8s:eachKey=+k8s:format=k8s-annotation-key` | PersistentVolume name is required and must be a DNS subdomain. |
| `spec.capacity` | `ResourceList` | `+k8s:required` | Mandatory map of resource names to quantities. Only `storage` is typically allowed. |
| `spec.accessModes` | `[]PersistentVolumeAccessMode` | `+k8s:required`<br>`+k8s:minItems=1`<br>`+k8s:eachVal=` | Mandatory ways the volume can be mounted. `ReadWriteOncePod` cannot be combined with other modes. |
| `spec.persistentVolumeReclaimPolicy` | `PersistentVolumeReclaimPolicy` | `+k8s:optional`<br> | Optional policy for maintenance after release. Defaults to `Retain`. |
| `spec.storageClassName` | `string` | `+k8s:optional` | Optional name of the StorageClass to which this volume belongs. |
| `spec.mountOptions` | `[]string` | `+k8s:optional` | Optional list of mount options. |
| `spec.volumeMode` | `*PersistentVolumeMode` | `+k8s:optional`<br> | Optional volume mode. Defaults to `Filesystem`. |
| `spec.nodeAffinity` | `*VolumeNodeAffinity` | `+k8s:optional` | Optional constraints that limit what nodes this volume can be accessed from. |
| `spec.claimRef` | `*ObjectReference` | `+k8s:optional` | Optional reference to the PersistentVolumeClaim. |
| `spec.gcePersistentDisk` | `*GCEPersistentDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | GCE Persistent Disk. |
| `spec.awsElasticBlockStore` | `*AWSElasticBlockStoreVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | AWS Elastic Block Store. |
| `spec.hostPath` | `*HostPathVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Host Path. |
| `spec.glusterfs` | `*GlusterfsPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | GlusterFS. |
| `spec.nfs` | `*NFSVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | NFS. |
| `spec.rbd` | `*RBDPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Rados Block Device. |
| `spec.iscsi` | `*ISCSIPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | ISCSI. |
| `spec.cinder` | `*CinderPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Cinder. |
| `spec.cephfs` | `*CephFSPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | CephFS. |
| `spec.fc` | `*FCVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Fibre Channel. |
| `spec.flocker` | `*FlockerVolumeSource" | `+k8s:unionMember`<br/>`+k8s:optional` | Flocker. |
| `spec.flexVolume` | `*FlexPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | FlexVolume. |
| `spec.azureFile` | `*AzureFilePersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Azure File. |
| `spec.vsphereVolume` | `*VsphereVirtualDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | vSphere Volume. |
| `spec.quobyte` | `*QuobyteVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Quobyte. |
| `spec.azureDisk` | `*AzureDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Azure Disk. |
| `spec.photonPersistentDisk` | `*PhotonPersistentDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Photon Persistent Disk. |
| `spec.portworxVolume` | `*PortworxVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Portworx Volume. |
| `spec.scaleIO` | `*ScaleIOPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | ScaleIO. |
| `spec.local` | `*LocalVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Local Volume. |
| `spec.storageos` | `*StorageOSPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | StorageOS. |
| `spec.csi` | `*CSIPersistentVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | CSI. |
| `spec.flexVolume.options` | `map[string]string` | `+k8s:eachKey=+k8s:format=k8s-flex-volume-option-key`<br/>`+k8s:optional` | Optional extra command options. |


## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `PersistentVolumeAccessMode` | `+k8s:enum` |
| `PersistentVolumeMode` | `+k8s:enum` |
| `PersistentVolumeReclaimPolicy` | `+k8s:enum` |
