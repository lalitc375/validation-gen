# Validation: Namespace

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Namespace name is required and must be a DNS subdomain. |
| `spec.finalizers` | `[]FinalizerName` | `+k8s:optional` | Optional list of namespaced-scoped resources that must be empty before the namespace is deleted. |

### NamespaceStatus

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `phase` | `NamespacePhase` | `+k8s:optional`<br> | The current lifecycle phase of the namespace. |


## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `NamespacePhase` | `+k8s:enum` |
