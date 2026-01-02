# Validation: Job

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Job name is required and must be a DNS subdomain. |
| `spec.parallelism` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional maximum desired number of pods the job should run at any given time. |
| `spec.completions` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional desired number of successfully finished pods. Required when `completionMode` is `Indexed`. |
| `spec.activeDeadlineSeconds` | `*int64` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional duration in seconds relative to the startTime that the job may be continuously active before termination. |
| `spec.backoffLimit` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional number of retries before marking this job failed. Defaults to 6. |
| `spec.selector` | `*metav1.LabelSelector` | `+k8s:required`<br>`+k8s:immutable` | Mandatory label query over pods. Immutable after creation. |
| `spec.template` | `corev1.PodTemplateSpec` | `+k8s:required` | Mandatory pod template. `restartPolicy` must be `Never` or `OnFailure`. |
| `spec.completionMode` | `*CompletionMode` | `+k8s:optional`<br>`+k8s:enum=["NonIndexed", "Indexed"]` | Optional completion mode. Defaults to `NonIndexed`. |
| `spec.suspend` | `*bool` | `+k8s:optional` | Optional flag to suspend subsequent executions. |
| `spec.ttlSecondsAfterFinished` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional lifetime of a Job that has finished execution. |
| `spec.backoffLimitPerIndex` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0`<br>`+k8s:immutable` | Optional limit for retries within an index. Only for `Indexed` completion mode. |
| `spec.maxFailedIndexes` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional maximal number of failed indexes before marking the Job as failed. Requires `backoffLimitPerIndex`. |
| `spec.podFailurePolicy` | `*PodFailurePolicy` | `+k8s:optional` | Optional policy for handling failed pods. |
| `spec.podFailurePolicy.rules` | `[]PodFailurePolicyRule` | `+k8s:required`<br/>`+k8s:maxItems=20` | Mandatory list of failure policy rules. |
| `spec.podFailurePolicy.rules[].onPodConditions` | `[]PodFailurePolicyOnPodConditionsPattern` | `+k8s:required`<br/>`+k8s:maxItems=20` | Mandatory list of pod conditions. |
| `spec.podFailurePolicy.rules[].onExitCodes.values` | `[]int32` | `+k8s:required`<br/>`+k8s:maxItems=255` | Mandatory list of exit codes. |
| `spec.successPolicy` | `*SuccessPolicy` | `+k8s:optional`<br>`+k8s:immutable` | Optional policy for declaring Job success. Only for `Indexed` Jobs. |
| `spec.successPolicy.rules` | `[]SuccessPolicyRule` | `+k8s:required`<br/>`+k8s:maxItems=20` | Mandatory list of success policy rules. |
| `spec.managedBy` | `*string` | `+k8s:optional`<br>`+k8s:maxLength=63` | Optional name of the controller that manages this job. |
