# Validation: LocalSubjectAccessReview

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=namespace)=+k8s:required` | LocalSubjectAccessReview requires a namespace in metadata. Other metadata fields must be empty. |
| `spec.user` | `string` | `+k8s:optional` | At least one of `user` or `groups` must be specified. |
| `spec.groups` | `[]string` | `+k8s:optional` | At least one of `user` or `groups` must be specified. |
| `spec.resourceAttributes` | `*ResourceAttributes` | `+k8s:required` | Mandatory for LocalSubjectAccessReview. `nonResourceAttributes` is disallowed. |
| `spec.resourceAttributes.namespace` | `string` | `+k8s:required` | Must match `metadata.namespace`. |
| `spec.resourceAttributes.verb` | `string` | `+k8s:optional` | Optional verb to check (e.g., `get`, `list`). |
| `spec.resourceAttributes.group` | `string` | `+k8s:optional` | Optional API group of the resource. |
| `spec.resourceAttributes.version` | `string` | `+k8s:optional` | Optional API version of the resource. |
| `spec.resourceAttributes.resource` | `string` | `+k8s:optional` | Optional resource type. |
| `spec.resourceAttributes.subresource` | `string` | `+k8s:optional` | Optional subresource. |
| `spec.resourceAttributes.name` | `string` | `+k8s:optional` | Optional resource name. |
| `spec.resourceAttributes.fieldSelector` | `*FieldSelectorAttributes` | `+k8s:optional` | Optional field selector limitation. |
| `spec.resourceAttributes.labelSelector` | `*LabelSelectorAttributes` | `+k8s:optional` | Optional label selector limitation. |
