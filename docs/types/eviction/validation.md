# Validation: Eviction

`Eviction` is a subresource of `Pod` used to cause pod eviction subject to disruption budgets.

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required` | Name of the pod to be evicted. |
| `deleteOptions` | `*metav1.DeleteOptions` | `+k8s:optional` | Optional options for pod deletion during eviction. |
