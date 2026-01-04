# Validation: IPAddress

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-ip` | IPAddress name is required and must be a valid IP address in canonical format. |
| `spec.parentRef` | `*ParentReference` | `+k8s:required`<br>`+k8s:immutable` | Mandatory reference to the resource that this IPAddress is attached to. Immutable after creation. |

### ParentReference

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `group` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional API group of the referent. |
| `resource` | `string` | `+k8s:required`<br>`+k8s:format=k8s-path-segment-name` | Mandatory resource of the referent. |
| `namespace` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-path-segment-name` | Optional namespace of the referent. |
| `name` | `string` | `+k8s:required`<br>`+k8s:format=k8s-path-segment-name` | Mandatory name of the referent. |
