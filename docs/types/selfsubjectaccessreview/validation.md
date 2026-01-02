# Validation: SelfSubjectAccessReview

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:empty` | Metadata must be empty for SelfSubjectAccessReview. |
| `spec` | `SelfSubjectAccessReviewSpec` | `+k8s:required` | Mandatory specification of the access request being evaluated. |
| `spec.resourceAttributes` | `*ResourceAttributes` | `+k8s:unionMember`<br/>`+k8s:optional` | Optional information for a resource access request. Mutually exclusive with `nonResourceAttributes`. |
| `spec.nonResourceAttributes` | `*NonResourceAttributes` | `+k8s:unionMember`<br/>`+k8s:optional` | Optional information for a non-resource access request. Mutually exclusive with `resourceAttributes`. |

### ResourceAttributes

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `namespace` | `string` | `+k8s:optional` | Optional namespace of the action. |
| `verb` | `string` | `+k8s:optional` | Optional verb to check (e.g., `get`, `create`). |
| `group` | `string` | `+k8s:optional` | Optional API group. |
| `version` | `string` | `+k8s:optional` | Optional API version. |
| `resource` | `string` | `+k8s:optional` | Optional resource type. |
| `subresource` | `string` | `+k8s:optional` | Optional subresource. |
| `name` | `string` | `+k8s:optional` | Optional resource name. |

### NonResourceAttributes

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `path` | `string` | `+k8s:optional` | Optional URL path of the request. |
| `verb` | `string` | `+k8s:optional` | Optional standard HTTP verb. |
