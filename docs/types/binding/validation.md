# Validation: Binding

`Binding` ties one object to another; for example, a pod is bound to a node by a scheduler.

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required` | Name of the object being bound (typically a Pod). |
| `target` | `ObjectReference` | `+k8s:required` | Mandatory reference to the object to bind to. |
| `target.kind` | `string` | `+k8s:optional`<br> | The kind of the target object. Currently only `Node` is supported. |
| `target.name` | `string` | `+k8s:required` | Mandatory name of the target object. |
