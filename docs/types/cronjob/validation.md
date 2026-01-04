# Validation: CronJob

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name`<br>`+k8s:subfield(name=name)=+k8s:maxLength=52` | CronJob name is required and must be a DNS subdomain. Max length is 52 characters to allow for job suffixes. |
| `spec.schedule` | `string` | `+k8s:required` | Mandatory cron-format schedule string. |
| `spec.timeZone` | `*string` | `+k8s:optional` | Optional IANA time zone for the schedule. |
| `spec.startingDeadlineSeconds` | `*int64` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional non-negative deadline for starting missed jobs. |
| `spec.concurrencyPolicy` | `ConcurrencyPolicy` | `+k8s:required`<br> | Mandatory policy for concurrent job executions. Defaults to `Allow`. |
| `spec.suspend` | `*bool` | `+k8s:optional` | Optional flag to suspend subsequent executions. |
| `spec.jobTemplate` | `JobTemplateSpec` | `+k8s:required` | Mandatory template for created jobs. |
| `spec.successfulJobsHistoryLimit` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional limit on successful jobs to retain. |
| `spec.failedJobsHistoryLimit` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional limit on failed jobs to retain. |


## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `ConcurrencyPolicy` | `+k8s:enum` |
