# Validation: Role

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Role name is required and must be a DNS subdomain. |
| `rules` | `[]PolicyRule` | `+k8s:required` | Mandatory list of rules that can be referenced as a unit. |

### PolicyRule

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `verbs` | `[]string` | `+k8s:required` | Mandatory list of verbs (e.g., `get`, `list`, `watch`, `create`, `update`, `patch`, `delete`). `*` represents all verbs. |
| `apiGroups` | `[]string` | `+k8s:optional` | Optional list of API groups. `""` represents the core API group, `*` represents all API groups. Required if `nonResourceURLs` is not set. |
| `resources` | `[]string` | `+k8s:optional` | Optional list of resources. `*` represents all resources. Required if `nonResourceURLs` is not set. |
| `resourceNames` | `[]string` | `+k8s:optional` | Optional white list of names that the rule applies to. |
| `nonResourceURLs` | `[]string` | `+k8s:optional` | Optional list of partial URLs. Forbidden for namespaced `Role` resources. |
