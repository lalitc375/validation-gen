# Validation: CSINode

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | CSINode name must be the Kubernetes node name. |
| `spec.drivers` | `[]CSINodeDriver` | `+k8s:required`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=name` | List of CSI drivers running on the node. |
| `spec.drivers[].name` | `string` | `+k8s:required`<br>`+k8s:format=k8s-long-name` | Mandatory CSI driver name. |
| `spec.drivers[].nodeID` | `string` | `+k8s:required`<br>`+k8s:maxLength=192` | Mandatory ID of the node from the driver's perspective. |
| `spec.drivers[].topologyKeys` | `[]string` | `+k8s:optional`<br>`+k8s:eachVal=+k8s:format=k8s-label-key`<br>`+k8s:listType=atomic` | Optional list of topology keys supported by the driver. |
| `spec.drivers[].allocatable` | `*VolumeNodeResources` | `+k8s:optional` | Optional volume resources available for scheduling. |
| `spec.drivers[].allocatable.count` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional maximum number of unique volumes. |
