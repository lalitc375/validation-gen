# Validation: Pod

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name`<br>`+k8s:subfield(name=labels)=+k8s:eachKey=+k8s:format=k8s-label-key`<br>`+k8s:subfield(name=annotations)=+k8s:eachKey=+k8s:format=k8s-annotation-key` | Pod name is required and must be a DNS subdomain. |
| `spec.containers` | `[]Container` | `+k8s:required`<br>`+k8s:minItems=1` | Mandatory list of containers belonging to the pod. |
| `spec.containers[].resources.limits` | `ResourceList` | `+k8s:eachKey=+k8s:format=k8s-container-resource-name` | Container resource limits. |
| `spec.containers[].resources.requests` | `ResourceList` | `+k8s:eachKey=+k8s:format=k8s-container-resource-name` | Container resource requests. |
| `spec.initContainers` | `[]Container` | `+k8s:optional` | Optional list of initialization containers. |
| `spec.initContainers[].resources.limits` | `ResourceList` | `+k8s:eachKey=+k8s:format=k8s-container-resource-name` | Container resource limits. |
| `spec.initContainers[].resources.requests` | `ResourceList` | `+k8s:eachKey=+k8s:format=k8s-container-resource-name` | Container resource requests. |
| `spec.ephemeralContainers` | `[]EphemeralContainer` | `+k8s:optional` | Optional list of ephemeral containers. Forbidden on creation. |
| `spec.restartPolicy` | `RestartPolicy` | `+k8s:optional`<br>`+k8s:enum=["Always", "OnFailure", "Never"]` | Optional restart policy for all containers. Defaults to `Always`. |
| `spec.terminationGracePeriodSeconds` | `*int64` | `+k8s:required`<br>`+k8s:minimum=0` | Mandatory duration in seconds for graceful termination. |
| `spec.activeDeadlineSeconds` | `*int64` | `+k8s:optional`<br>`+k8s:minimum=1` | Optional duration in seconds the pod may be active. |
| `spec.dnsPolicy` | `DNSPolicy` | `+k8s:optional`<br>`+k8s:enum=["ClusterFirstWithHostNet", "ClusterFirst", "Default", "None"]` | Optional DNS policy. Defaults to `ClusterFirst`. |
| `spec.nodeSelector` | `map[string]string` | `+k8s:optional`<br>`+k8s:eachKey=+k8s:format=k8s-label-key` | Optional selector which must be true for the pod to fit on a node. |
| `spec.serviceAccountName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional name of the ServiceAccount to use. |
| `spec.nodeName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional node on which this pod is scheduled. Immutable once set. |
| `spec.affinity` | `*Affinity` | `+k8s:optional` | Optional scheduling constraints. |
| `spec.tolerations` | `[]Toleration` | `+k8s:optional` | Optional list of tolerations. |
| `spec.hostAliases` | `[]HostAlias` | `+k8s:optional` | Optional list of host aliases. |
| `spec.hostAliases[].hostnames` | `[]string` | `+k8s:eachVal=+k8s:format=k8s-dns-subdomain` | Hostnames for the host alias. |
| `spec.priorityClassName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional priority class name. |
| `spec.runtimeClassName` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional runtime class name. |
| `spec.volumes` | `[]Volume` | `+k8s:optional` | Optional list of volumes. |
| `spec.volumes[].name` | `string` | `+k8s:required`<br>`+k8s:format=k8s-dns-label` | Volume name. |
| `spec.volumes[].hostPath` | `*HostPathVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Host Path. |
| `spec.volumes[].emptyDir` | `*EmptyDirVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Empty Dir. |
| `spec.volumes[].gcePersistentDisk` | `*GCEPersistentDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | GCE Persistent Disk. |
| `spec.volumes[].awsElasticBlockStore` | `*AWSElasticBlockStoreVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | AWS Elastic Block Store. |
| `spec.volumes[].gitRepo` | `*GitRepoVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Git Repo (Deprecated). |
| `spec.volumes[].secret` | `*SecretVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Secret. |
| `spec.volumes[].nfs` | `*NFSVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | NFS. |
| `spec.volumes[].iscsi` | `*ISCSIVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | ISCSI. |
| `spec.volumes[].glusterfs` | `*GlusterfsVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | GlusterFS. |
| `spec.volumes[].persistentVolumeClaim` | `*PersistentVolumeClaimVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Persistent Volume Claim. |
| `spec.volumes[].rbd` | `*RBDVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Rados Block Device. |
| `spec.volumes[].flexVolume` | `*FlexVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | FlexVolume. |
| `spec.volumes[].cinder` | `*CinderVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Cinder. |
| `spec.volumes[].cephfs` | `*CephFSVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | CephFS. |
| `spec.volumes[].flocker` | `*FlockerVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Flocker. |
| `spec.volumes[].downwardAPI` | `*DownwardAPIVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Downward API. |
| `spec.volumes[].fc` | `*FCVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Fibre Channel. |
| `spec.volumes[].azureFile` | `*AzureFileVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Azure File. |
| `spec.volumes[].configMap` | `*ConfigMapVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | ConfigMap. |
| `spec.volumes[].vsphereVolume` | `*VsphereVirtualDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | vSphere Volume. |
| `spec.volumes[].quobyte` | `*QuobyteVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Quobyte. |
| `spec.volumes[].azureDisk` | `*AzureDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Azure Disk. |
| `spec.volumes[].photonPersistentDisk` | `*PhotonPersistentDiskVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Photon Persistent Disk. |
| `spec.volumes[].projected` | `*ProjectedVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Projected Volume. |
| `spec.volumes[].portworxVolume` | `*PortworxVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Portworx Volume. |
| `spec.volumes[].scaleIO` | `*ScaleIOVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | ScaleIO. |
| `spec.volumes[].storageos` | `*StorageOSVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | StorageOS. |
| `spec.volumes[].csi` | `*CSIVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | CSI. |
| `spec.volumes[].ephemeral` | `*EphemeralVolumeSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Ephemeral. |
| `spec.volumes[].projected.sources[].secret` | `*SecretProjection` | `+k8s:unionMember`<br/>`+k8s:optional` | Secret projection. |
| `spec.volumes[].projected.sources[].downwardAPI` | `*DownwardAPIProjection` | `+k8s:unionMember`<br/>`+k8s:optional` | Downward API projection. |
| `spec.volumes[].projected.sources[].configMap` | `*ConfigMapProjection` | `+k8s:unionMember`<br/>`+k8s:optional` | ConfigMap projection. |
| `spec.volumes[].projected.sources[].serviceAccountToken` | `*ServiceAccountTokenProjection` | `+k8s:unionMember`<br/>`+k8s:optional` | ServiceAccountToken projection. |
| `spec.volumes[].projected.sources[].clusterTrustBundle` | `*ClusterTrustBundleProjection` | `+k8s:unionMember`<br/>`+k8s:optional` | ClusterTrustBundle projection. |
| `spec.volumes[].projected.sources[].podCertificate` | `*PodCertificateProjection` | `+k8s:unionMember`<br/>`+k8s:optional` | PodCertificate projection. |
| `spec.volumes[].downwardAPI.items[].fieldRef` | `*ObjectFieldSelector` | `+k8s:unionMember`<br/>`+k8s:optional` | Field Ref. |
| `spec.volumes[].downwardAPI.items[].resourceFieldRef` | `*ResourceFieldSelector` | `+k8s:unionMember`<br/>`+k8s:optional` | Resource Field Ref. |
| `spec.containers[].env` | `[]EnvVar` | `+k8s:optional` | Environment variables. |
| `spec.containers[].env[].value` | `string` | `+k8s:unionMember`<br/>`+k8s:optional` | Value. |
| `spec.containers[].env[].valueFrom` | `*EnvVarSource` | `+k8s:unionMember`<br/>`+k8s:optional` | Value From. |
| `spec.containers[].env[].valueFrom.fieldRef` | `*ObjectFieldSelector` | `+k8s:unionMember`<br/>`+k8s:optional` | Field Ref. |
| `spec.containers[].env[].valueFrom.resourceFieldRef` | `*ResourceFieldSelector` | `+k8s:unionMember`<br/>`+k8s:optional` | Resource Field Ref. |
| `spec.containers[].env[].valueFrom.configMapKeyRef` | `*ConfigMapKeySelector` | `+k8s:unionMember`<br/>`+k8s:optional` | ConfigMap Key Ref. |
| `spec.containers[].env[].valueFrom.secretKeyRef` | `*SecretKeySelector` | `+k8s:unionMember`<br/>`+k8s:optional` | Secret Key Ref. |
| `spec.containers[].env[].valueFrom.fileKeyRef` | `*FileKeySelector` | `+k8s:unionMember`<br/>`+k8s:optional` | File Key Ref. |
| `spec.containers[].livenessProbe.exec` | `*ExecAction` | `+k8s:unionMember`<br/>`+k8s:optional` | Exec action. |
| `spec.containers[].livenessProbe.initialDelaySeconds` | `int32` | `+k8s:optional`<br/>`+k8s:minimum=0` | Number of seconds after the container has started before liveness probes are initiated. |
| `spec.containers[].livenessProbe.timeoutSeconds` | `int32` | `+k8s:optional`<br/>`+k8s:minimum=1` | Number of seconds after which the probe times out. |
| `spec.containers[].livenessProbe.periodSeconds` | `int32` | `+k8s:optional`<br/>`+k8s:minimum=1` | How often (in seconds) to perform the probe. |
| `spec.containers[].livenessProbe.successThreshold` | `int32` | `+k8s:optional`<br/>`+k8s:minimum=1` | Minimum consecutive successes for the probe to be considered successful after having failed. |
| `spec.containers[].livenessProbe.failureThreshold` | `int32` | `+k8s:optional`<br/>`+k8s:minimum=1` | Minimum consecutive failures for the probe to be considered failed after having succeeded. |
| `spec.containers[].livenessProbe.httpGet` | `*HTTPGetAction` | `+k8s:unionMember`<br/>`+k8s:optional` | HTTP Get action. |
| `spec.containers[].livenessProbe.tcpSocket` | `*TCPSocketAction` | `+k8s:unionMember`<br/>`+k8s:optional` | TCP Socket action. |
| `spec.containers[].livenessProbe.grpc` | `*GRPCAction` | `+k8s:unionMember`<br/>`+k8s:optional` | GRPC action. |
| `spec.containers[].lifecycle.postStart.exec` | `*ExecAction` | `+k8s:unionMember`<br/>`+k8s:optional` | Exec action. |
| `spec.containers[].lifecycle.postStart.httpGet` | `*HTTPGetAction` | `+k8s:unionMember`<br/>`+k8s:optional` | HTTP Get action. |
| `spec.containers[].lifecycle.postStart.tcpSocket` | `*TCPSocketAction` | `+k8s:unionMember`<br/>`+k8s:optional` | TCP Socket action. |
| `spec.containers[].lifecycle.preStop.exec` | `*ExecAction` | `+k8s:unionMember`<br/>`+k8s:optional` | Exec action. |
| `spec.containers[].lifecycle.preStop.httpGet` | `*HTTPGetAction` | `+k8s:unionMember`<br/>`+k8s:optional` | HTTP Get action. |
| `spec.containers[].lifecycle.preStop.tcpSocket` | `*TCPSocketAction` | `+k8s:unionMember`<br/>`+k8s:optional` | TCP Socket action. |
| `spec.securityContext.seccompProfile.type` | `SeccompProfileType` | `+k8s:unionDiscriminator` | Seccomp profile type. |
| `spec.securityContext.seccompProfile.localhostProfile` | `*string` | `+k8s:unionMember(memberName="Localhost")`<br/>`+k8s:optional` | Localhost profile. |
| `spec.containers[].securityContext.seccompProfile.type` | `SeccompProfileType` | `+k8s:unionDiscriminator` | Seccomp profile type. |
| `spec.containers[].securityContext.seccompProfile.localhostProfile` | `*string` | `+k8s:unionMember(memberName="Localhost")`<br/>`+k8s:optional` | Localhost profile. |
