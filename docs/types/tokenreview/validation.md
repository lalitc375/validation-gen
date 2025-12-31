# Validation: TokenReview

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:empty` | Metadata must be empty for TokenReview. |
| `spec` | `TokenReviewSpec` | `+k8s:required` | Mandatory specification of the token authentication request. |
| `spec.token` | `string` | `+k8s:required` | Mandatory opaque bearer token to authenticate. |
| `spec.audiences` | `[]string` | `+k8s:optional` | Optional list of audiences the token must be valid for. |
