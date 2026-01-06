# Validation: ValidatingWebhookConfiguration

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Configuration name is required and must be a DNS subdomain. |
| `webhooks` | `[]ValidatingWebhook` | `+k8s:optional`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=name` | List of webhooks and the affected resources and operations. |
| `webhooks[].name` | `string` | `+k8s:required` | Mandatory fully qualified name of the admission webhook. |
| `webhooks[].clientConfig` | `WebhookClientConfig` | `+k8s:required` | Mandatory instructions for how to communicate with the hook. Exactly one of `url` or `service` must be set. |
| `webhooks[].clientConfig.url` | `*string` | `+k8s:unionMember`<br/>`+k8s:optional` | Optional URL. Mutually exclusive with `service`. |
| `webhooks[].clientConfig.service` | `*ServiceReference` | `+k8s:unionMember`<br/>`+k8s:optional` | Optional service reference. Mutually exclusive with `url`. |
| `webhooks[].rules` | `[]RuleWithOperations` | `+k8s:optional`<br>`+k8s:listType=atomic` | Optional list of operations on what resources/subresources the webhook cares about. |
| `webhooks[].failurePolicy` | `*FailurePolicyType` | `+k8s:optional`<br> | Optional failure policy. Defaults to `Ignore`. |
| `webhooks[].matchPolicy` | `*MatchPolicyType` | `+k8s:optional`<br> | Optional match policy. Defaults to `Exact`. |
| `webhooks[].namespaceSelector` | `*metav1.LabelSelector` | `+k8s:optional` | Optional selector to run the webhook only on objects in matching namespaces. |
| `webhooks[].objectSelector` | `*metav1.LabelSelector` | `+k8s:optional` | Optional selector to run the webhook only on matching objects. |
| `webhooks[].sideEffects` | `*SideEffectClass` | `+k8s:required`<br> | Mandatory statement on side effects. `Unknown` and `Some` are typically disallowed for modern webhooks. |
| `webhooks[].timeoutSeconds` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=1`<br>`+k8s:maximum=30` | Optional timeout in seconds. Defaults to 10. |
| `webhooks[].admissionReviewVersions` | `[]string` | `+k8s:required`<br>`+k8s:listType=atomic` | Mandatory list of `AdmissionReview` versions the webhook accepts. |
| `webhooks[].matchConditions` | `[]MatchCondition` | `+k8s:optional`<br>`+k8s:maxItems=64`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=name` | Optional list of CEL conditions that must be met for the webhook to be called. |


## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `FailurePolicyType` | `+k8s:enum` |
| `MatchPolicyType` | `+k8s:enum` |
| `SideEffectClass` | `+k8s:enum` |
