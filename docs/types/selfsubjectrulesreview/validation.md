# Validation: SelfSubjectRulesReview

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:empty` | Metadata must be empty for SelfSubjectRulesReview. |
| `spec` | `SelfSubjectRulesReviewSpec` | `+k8s:required` | Mandatory specification for the rules review. |
| `spec.namespace` | `string` | `+k8s:optional` | Optional namespace to evaluate rules for. If not specified, rules are evaluated for all namespaces. |
