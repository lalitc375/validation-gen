# Validation: Lease

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Lease name is required and must be a DNS subdomain. |
| `spec.holderIdentity` | `*string` | `+k8s:optional` | Optional identity of the holder of the current lease. |
| `spec.leaseDurationSeconds` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=1` | Optional duration that candidates for a lease need to wait to force acquire it. |
| `spec.acquireTime` | `*metav1.MicroTime` | `+k8s:optional` | Optional time when the current lease was acquired. |
| `spec.renewTime` | `*metav1.MicroTime` | `+k8s:optional` | Optional time when the current holder of a lease has last updated the lease. |
| `spec.leaseTransitions` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional number of transitions of a lease between holders. |
| `spec.strategy` | `*CoordinatedLeaseStrategy` | `+k8s:optional` | Optional strategy for picking the leader for coordinated leader election. Currently only `OldestEmulationVersion` is supported by Kubernetes. |
| `spec.preferredHolder` | `*string` | `+k8s:optional` | Optional signal to a lease holder that the lease has a more optimal holder. Only if `strategy` is set. |


## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `CoordinatedLeaseStrategy` | `+k8s:enum` |
