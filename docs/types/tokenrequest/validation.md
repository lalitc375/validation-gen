# Validation: TokenRequest

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:empty` | Metadata must be empty for TokenRequest. |
| `spec` | `TokenRequestSpec` | `+k8s:required` | Mandatory specification of the token request. |
| `spec.audiences` | `[]string` | `+k8s:required`<br>`+k8s:minItems=1` | Mandatory intended audiences of the token. |
| `spec.expirationSeconds` | `int64` | `+k8s:optional`<br>`+k8s:minimum=600` | Optional requested duration of validity. Must be at least 10 minutes (600 seconds). |
| `spec.boundObjectRef` | `*BoundObjectReference` | `+k8s:optional` | Optional reference to an object that the token will be bound to. |

### BoundObjectReference

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `kind` | `string` | `+k8s:optional`<br> | Optional kind of the referent. |
| `apiVersion` | `string` | `+k8s:optional` | Optional API version of the referent. |
| `name` | `string` | `+k8s:optional` | Optional name of the referent. |
| `uid` | `types.UID` | `+k8s:optional` | Optional UID of the referent. |
