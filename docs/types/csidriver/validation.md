# Validation: CSIDriver

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | CSIDriver name is required and must be a DNS subdomain (the CSI driver name). |
| `spec.attachedRequired` | `*bool` | `+k8s:required`<br>`+k8s:immutable` | Indicates if the driver requires an attach operation. Immutable after creation. |
| `spec.podInfoOnMount` | `*bool` | `+k8s:required` | Indicates if pod info is passed on mount. |
| `spec.storageCapacity` | `*bool` | `+k8s:required` | Indicates if the driver produces capacity information. |
| `spec.fsGroupPolicy` | `*FSGroupPolicy` | `+k8s:optional`<br> | Controls if Kubernetes should modify volume ownership and permissions. |
| `spec.volumeLifecycleModes` | `[]VolumeLifecycleMode` | `+k8s:optional`<br>`+k8s:immutable`<br>`+k8s:eachVal=` | Supported volume modes. Immutable after creation. |
| `spec.nodeAllocatableUpdatePeriodSeconds` | `*int64` | `+k8s:optional`<br>`+k8s:minimum=10` | Period for node allocatable updates. Must be at least 10 seconds. |


## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `FSGroupPolicy` | `+k8s:enum` |
| `VolumeLifecycleMode` | `+k8s:enum` |
