# Validation: LeaseCandidate

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | LeaseCandidate name is required. |
| `spec.leaseName` | `string` | `+k8s:required`<br>`+k8s:immutable` | Mandatory name of the lease for which this candidate is contending. Immutable after creation. |
| `spec.binaryVersion` | `string` | `+k8s:required` | Mandatory binary version in semver format (without leading `v`). |
| `spec.emulationVersion` | `string` | `+k8s:optional` | Optional emulation version in semver format. Must be <= `binaryVersion`. |
| `spec.strategy` | `CoordinatedLeaseStrategy` | `+k8s:required`<br>`+k8s:enum=["OldestEmulationVersion"]` | Mandatory strategy for picking the leader. |
| `spec.pingTime` | `*metav1.MicroTime` | `+k8s:optional` | Optional last time that the server has requested the candidate to renew. |
| `spec.renewTime` | `*metav1.MicroTime` | `+k8s:optional` | Optional time that the candidate was last updated. |
