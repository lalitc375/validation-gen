# Validation: ClusterTrustBundle

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | ClusterTrustBundle name is required. If `signerName` is specified, the name must be prefixed with `signerName` (replacing `/` with `:`) followed by a colon. |
| `spec.signerName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-label-key`<br>`+k8s:immutable` | Optional name of the associated signer. All signer names beginning with `kubernetes.io` are reserved. Immutable after creation. |
| `spec.trustBundle` | `string` | `+k8s:required`<br>`+k8s:maxLength=1048576` | Mandatory PEM bundle of X.509 certificates. Must consist only of valid PEM certificate blocks. Each certificate must be a CA. Max size is 1 MiB. |