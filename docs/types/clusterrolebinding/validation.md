# Validation: ClusterRoleBinding

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | ClusterRoleBinding name is required and must be a DNS subdomain. |
| `subjects` | `[]Subject` | `+k8s:optional` | Optional list of subjects (users, groups, or service accounts) that the role applies to. |
| `roleRef` | `RoleRef` | `+k8s:required`<br>`+k8s:immutable` | Mandatory reference to the ClusterRole being used. Immutable after creation. |
| `roleRef.apiGroup` | `string` | `+k8s:required`<br> | Mandatory API group for the role being referenced. |
| `roleRef.kind` | `string` | `+k8s:required`<br> | Mandatory type of role being referenced. Must be `ClusterRole`. |
| `roleRef.name` | `string` | `+k8s:required`<br>`+k8s:format=k8s-path-segment-name` | Mandatory name of the ClusterRole being referenced. |