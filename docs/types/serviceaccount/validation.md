# Validation: ServiceAccount

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | ServiceAccount name is required and must be a DNS subdomain. |
| `secrets` | `[]ObjectReference` | `+k8s:optional` | Optional list of secrets allowed to be used by pods running with this ServiceAccount. |
| `imagePullSecrets` | `[]LocalObjectReference` | `+k8s:optional` | Optional list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. |
| `automountServiceAccountToken` | `*bool` | `+k8s:optional` | Optional flag indicating whether a service account token should be automatically mounted. |
