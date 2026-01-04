# Validation: RoleBinding

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | RoleBinding name is required and must be a DNS subdomain. |
| `subjects` | `[]Subject` | `+k8s:optional` | Optional list of subjects (users, groups, or service accounts) that the role applies to. |
| `roleRef` | `RoleRef` | `+k8s:required`<br>`+k8s:immutable` | Mandatory reference to the role being used. Immutable after creation. |
| `roleRef.apiGroup` | `string` | `+k8s:required`<br> | Mandatory API group for the role being referenced. |
| `roleRef.kind` | `string` | `+k8s:required`<br> | Mandatory type of role being referenced. |
| `roleRef.name` | `string` | `+k8s:required`<br>`+k8s:format=k8s-path-segment-name` | Mandatory name of the role being referenced. |

### Subject

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `kind` | `string` | `+k8s:required`<br> | Mandatory type of subject being referenced. |
| `apiGroup` | `string` | `+k8s:optional` | Mandatory for `User` and `Group` (must be `rbac.authorization.k8s.io`). Must be empty for `ServiceAccount`. |
| `name` | `string` | `+k8s:required` | Mandatory name of the subject. |
| `namespace` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Required for `ServiceAccount` if the subject is not in the same namespace as the RoleBinding. |
