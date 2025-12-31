# Validation: ValidatingWebhookConfiguration

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Configuration name is required and must be a DNS subdomain. |
| `webhooks` | `[]ValidatingWebhook` | `+k8s:optional` | List of webhooks and the affected resources and operations. |
| `webhooks[].name` | `string` | `+k8s:required` | Mandatory fully qualified name of the admission webhook. |
| `webhooks[].clientConfig` | `WebhookClientConfig` | `+k8s:required` | Mandatory instructions for how to communicate with the hook. Exactly one of `url` or `service` must be set. |
| `webhooks[].rules` | `[]RuleWithOperations` | `+k8s:optional` | Optional list of operations on what resources/subresources the webhook cares about. |
| `webhooks[].failurePolicy` | `*FailurePolicyType` | `+k8s:optional`<br>`+k8s:enum=["Ignore", "Fail"]` | Optional failure policy. Defaults to `Ignore`. |
| `webhooks[].matchPolicy` | `*MatchPolicyType` | `+k8s:optional`<br>`+k8s:enum=["Exact", "Equivalent"]` | Optional match policy. Defaults to `Exact`. |
| `webhooks[].namespaceSelector` | `*metav1.LabelSelector` | `+k8s:optional` | Optional selector to run the webhook only on objects in matching namespaces. |
| `webhooks[].objectSelector` | `*metav1.LabelSelector` | `+k8s:optional` | Optional selector to run the webhook only on matching objects. |
| `webhooks[].sideEffects` | `*SideEffectClass` | `+k8s:required`<br>`+k8s:enum=["None", "NoneOnDryRun"]` | Mandatory statement on side effects. `Unknown` and `Some` are typically disallowed for modern webhooks. |
| `webhooks[].timeoutSeconds` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=1`<br>`+k8s:maximum=30` | Optional timeout in seconds. Defaults to 10. |
| `webhooks[].admissionReviewVersions` | `[]string` | `+k8s:required`<br>`+k8s:minItems=1`<br>`+k8s:eachVal=+k8s:enum=["v1", "v1beta1"]` | Mandatory list of `AdmissionReview` versions the webhook accepts. |
| `webhooks[].matchConditions` | `[]MatchCondition` | `+k8s:optional`<br>`+k8s:maxItems=64` | Optional list of CEL conditions that must be met for the webhook to be called. |
