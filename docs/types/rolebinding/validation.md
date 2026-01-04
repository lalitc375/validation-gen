# Validation: RoleBinding

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | RoleBinding name is required and must be a DNS subdomain. |
| `subjects` | `[]Subject` | `+k8s:optional`<br>`+k8s:listType=atomic` | Optional list of subjects. |
| `roleRef` | `RoleRef` | `+k8s:required`<br>`+k8s:immutable` | Mandatory reference to the role being used. Immutable after creation. |

### Subject

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `kind` | `string` | `+k8s:required`<br> | Mandatory type of subject being referenced. |
| `apiGroup` | `string` | `+k8s:optional` | Mandatory for `User` and `Group` (must be `rbac.authorization.k8s.io`). Must be empty for `ServiceAccount`. |
| `name` | `string` | `+k8s:required` | Mandatory name of the subject. |
| `namespace` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Required for `ServiceAccount` if the subject is not in the same namespace as the RoleBinding. |

### RoleRef

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `apiGroup` | `string` | `+k8s:required` | Mandatory API group for the role being referenced. |
| `kind` | `string` | `+k8s:required` | Mandatory type of role being referenced. |
| `name` | `string` | `+k8s:required`<br>`+k8s:format=k8s-path-segment-name` | Mandatory name of the role being referenced. |