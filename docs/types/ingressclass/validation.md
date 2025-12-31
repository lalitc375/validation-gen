# Validation: IngressClass

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | IngressClass name is required and must be a DNS subdomain. |
| `spec.controller` | `string` | `+k8s:required`<br>`+k8s:immutable`<br>`+k8s:maxLength=250` | Mandatory name of the controller that should handle this class. Immutable after creation. |
| `spec.parameters` | `*IngressClassParametersReference` | `+k8s:optional` | Optional link to a custom resource containing additional configuration. |

### IngressClassParametersReference

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `apiGroup` | `*string` | `+k8s:optional` | Optional API group of the resource. |
| `kind` | `string` | `+k8s:required` | Mandatory kind of the resource. |
| `name` | `string` | `+k8s:required` | Mandatory name of the resource. |
| `scope` | `*string` | `+k8s:optional`<br>`+k8s:enum=["Cluster", "Namespace"]` | Optional scope. Defaults to `Cluster`. |
| `namespace` | `*string` | `+k8s:optional` | Required if `scope` is `Namespace`. Must be empty if `scope` is `Cluster`. |
