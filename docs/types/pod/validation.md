# Validation: Pod

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name`<br>`+k8s:subfield(name=labels)=+k8s:eachKey=+k8s:format=k8s-label-key`<br>`+k8s:subfield(name=annotations)=+k8s:eachKey=+k8s:format=k8s-annotation-key` | Pod name is required and must be a DNS subdomain. |
| `spec.containers` | `[]Container` | `+k8s:required`<br>`+k8s:minItems=1`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=name` | Mandatory list of containers belonging to the pod. |
| `spec.initContainers` | `[]Container` | `+k8s:optional`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=name` | Optional list of initialization containers. |
| `spec.ephemeralContainers` | `[]EphemeralContainer` | `+k8s:optional`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=name` | Optional list of ephemeral containers. Forbidden on creation. |
| `spec.volumes` | `[]Volume` | `+k8s:optional`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=name` | Optional list of volumes. |
| `spec.restartPolicy` | `RestartPolicy` | `+k8s:optional`<br> | Optional restart policy. Defaults to `Always`. |
| `spec.terminationGracePeriodSeconds` | `*int64` | `+k8s:required`<br>`+k8s:minimum=0` | Mandatory duration for graceful termination. |
| `spec.activeDeadlineSeconds` | `*int64` | `+k8s:optional`<br>`+k8s:minimum=1` | Optional duration the pod may be active. |
| `spec.dnsPolicy` | `DNSPolicy` | `+k8s:optional`<br> | Optional DNS policy. Defaults to `ClusterFirst`. |
| `spec.nodeSelector` | `map[string]string` | `+k8s:optional`<br>`+k8s:eachKey=+k8s:format=k8s-label-key` | Optional node selector. |
| `spec.serviceAccountName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional name of the ServiceAccount. |
| `spec.nodeName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional node name. Immutable once set. |
| `spec.affinity` | `*Affinity` | `+k8s:optional` | Optional scheduling constraints. |
| `spec.tolerations` | `[]Toleration` | `+k8s:optional`<br>`+k8s:listType=atomic` | Optional list of tolerations. |
| `spec.hostAliases` | `[]HostAlias` | `+k8s:optional`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=ip` | Optional list of host aliases. |
| `spec.priorityClassName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional priority class name. |
| `spec.runtimeClassName` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional runtime class name. |

### Container

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `name` | `string` | `+k8s:required`<br>`+k8s:format=k8s-short-name` | Mandatory name of the container. |
| `ports` | `[]ContainerPort` | `+k8s:optional`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=containerPort`<br>`+k8s:listMapKey=protocol` | List of network ports. |
| `env` | `[]EnvVar` | `+k8s:optional`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=name` | Environment variables. |
| `envFrom` | `[]EnvFromSource` | `+k8s:optional`<br>`+k8s:listType=atomic` | Sources to populate environment variables. |
| `volumeMounts` | `[]VolumeMount` | `+k8s:optional`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=mountPath` | Pod volumes to mount. |
| `volumeDevices` | `[]VolumeDevice` | `+k8s:optional`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=devicePath` | Block devices to use. |
| `resources.limits` | `ResourceList` | `+k8s:eachKey=+k8s:format=k8s-container-resource-name` | Resource limits. |
| `resources.requests` | `ResourceList` | `+k8s:eachKey=+k8s:format=k8s-container-resource-name` | Resource requests. |
| `livenessProbe` / `readinessProbe` / `startupProbe` | `*Probe` | `+k8s:optional` | Health check probes. |

### EphemeralContainer

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `name` | `string` | `+k8s:required`<br>`+k8s:format=k8s-short-name` | Mandatory name. |
| `targetContainerName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Target container name. |

### Volume

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `name` | `string` | `+k8s:required`<br>`+k8s:format=k8s-short-name` | Mandatory volume name. |
| `hostPath` | `*HostPathVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Host Path. |
| `emptyDir` | `*EmptyDirVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Empty Dir. |
| `gcePersistentDisk` | `*GCEPersistentDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | GCE Persistent Disk. |
| `awsElasticBlockStore` | `*AWSElasticBlockStoreVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | AWS Elastic Block Store. |
| `gitRepo` | `*GitRepoVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Git Repo (Deprecated). |
| `secret` | `*SecretVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Secret. |
| `nfs` | `*NFSVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | NFS. |
| `iscsi` | `*ISCSIVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | ISCSI. |
| `glusterfs` | `*GlusterfsVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | GlusterFS. |
| `persistentVolumeClaim` | `*PersistentVolumeClaimVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Persistent Volume Claim. |
| `rbd` | `*RBDVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Rados Block Device. |
| `flexVolume` | `*FlexVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | FlexVolume. |
| `cinder` | `*CinderVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Cinder. |
| `cephfs` | `*CephFSVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | CephFS. |
| `flocker` | `*FlockerVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Flocker. |
| `downwardAPI` | `*DownwardAPIVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Downward API. |
| `fc` | `*FCVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Fibre Channel. |
| `azureFile` | `*AzureFileVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Azure File. |
| `configMap` | `*ConfigMapVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | ConfigMap. |
| `vsphereVolume` | `*VsphereVirtualDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | vSphere Volume. |
| `quobyte` | `*QuobyteVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Quobyte. |
| `azureDisk` | `*AzureDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Azure Disk. |
| `photonPersistentDisk` | `*PhotonPersistentDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Photon Persistent Disk. |
| `projected` | `*ProjectedVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Projected Volume. |
| `portworxVolume` | `*PortworxVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Portworx Volume. |
| `scaleIO` | `*ScaleIOVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | ScaleIO. |
| `storageos` | `*StorageOSVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | StorageOS. |
| `csi` | `*CSIVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | CSI. |
| `ephemeral` | `*EphemeralVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Ephemeral. |

### EnvVar

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `name` | `string` | `+k8s:required`<br>`+k8s:format=k8s-relaxed-env-var-name` | Mandatory name. |
| `value` | `string` | `+k8s:unionMember`<br/>`+k8s:optional` | Value. |
| `valueFrom` | `*EnvVarSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Value From. |

### EnvVarSource

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `fieldRef` | `*ObjectFieldSelector` | `+k8s:unionMember`<br/>`+k8s:optional` | Field Ref. |
| `resourceFieldRef` | `*ResourceFieldSelector` | `+k8s:unionMember`<br/>`+k8s:optional` | Resource Field Ref. |
| `configMapKeyRef` | `*ConfigMapKeySelector` | `+k8s:unionMember`<br/>`+k8s:optional` | ConfigMap Key Ref. |
| `secretKeyRef` | `*SecretKeySelector` | `+k8s:unionMember`<br/>`+k8s:optional` | Secret Key Ref. |
| `fileKeyRef` | `*FileKeySelector` | `+k8s:unionMember`<br/>`+k8s:optional` | File Key Ref. |

### Probe

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `exec` | `*ExecAction` | `+k8s:unionMember`<br/>`+k8s:optional` | Exec action. |
| `httpGet` | `*HTTPGetAction` | `+k8s:unionMember`<br/>`+k8s:optional` | HTTP Get action. |
| `tcpSocket` | `*TCPSocketAction` | `+k8s:unionMember`<br/>`+k8s:optional` | TCP Socket action. |
| `grpc` | `*GRPCAction` | `+k8s:unionMember`<br/>`+k8s:optional` | GRPC action. |
| `initialDelaySeconds` | `int32` | `+k8s:optional`<br/>`+k8s:minimum=0` | Initial delay. |
| `timeoutSeconds` | `int32` | `+k8s:optional`<br/>`+k8s:minimum=1` | Timeout. |
| `periodSeconds` | `int32` | `+k8s:optional`<br/>`+k8s:minimum=1` | Period. |
| `successThreshold` | `int32` | `+k8s:optional`<br/>`+k8s:minimum=1` | Success threshold. |
| `failureThreshold` | `int32` | `+k8s:optional`<br/>`+k8s:minimum=1` | Failure threshold. |

### SeccompProfile

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `type` | `SeccompProfileType` | `+k8s:unionDiscriminator` | Seccomp profile type. |
| `localhostProfile` | `*string` | `+k8s:unionMember(memberName="Localhost")`<br/>`+k8s:optional` | Localhost profile. |

### HostAlias

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `ip` | `string` | `+k8s:required`<br>`+k8s:format=k8s-ip` | Mandatory IP. |
| `hostnames` | `[]string` | `+k8s:eachVal=+k8s:format=k8s-long-name` | List of hostnames. |

## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `DNSPolicy` | `+k8s:enum` |
| `RestartPolicy` | `+k8s:enum` |
| `SeccompProfileType` | `+k8s:enum` |