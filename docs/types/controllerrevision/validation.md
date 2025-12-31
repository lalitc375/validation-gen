# Validation: ControllerRevision

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | ControllerRevision name is required and must be a DNS subdomain. |
| `data` | `runtime.RawExtension` | `+k8s:required`<br>`+k8s:immutable` | The state data is mandatory and cannot be changed after creation. |
| `revision` | `int64` | `+k8s:required`<br>`+k8s:minimum=0` | The revision number is mandatory and must be non-negative. |
