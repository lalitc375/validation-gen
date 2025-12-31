# Validation: Secret

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Secret name is required and must be a DNS subdomain. |
| `data` | `map[string][]byte` | `+k8s:optional` | Optional map of secret data. Each value is a base64 encoded string. Total size must not exceed 1 MiB. |
| `immutable` | `*bool` | `+k8s:optional` | Optional flag. If true, ensures that data stored in the Secret cannot be updated (only metadata can be modified). |
| `type` | `SecretType` | `+k8s:optional`<br>`+k8s:immutable` | Optional type used to facilitate programmatic handling of secret data. Immutable after creation. |

### Validation Rules for Specific Types

- **`kubernetes.io/service-account-token`**: Requires `kubernetes.io/service-account.name` annotation.
- **`kubernetes.io/dockercfg`**: Must contain a `.dockercfg` key with a valid JSON value.
- **`kubernetes.io/dockerconfigjson`**: Must contain a `.dockerconfigjson` key with a valid JSON value.
- **`kubernetes.io/basic-auth`**: Must contain either `username` or `password` keys.
- **`kubernetes.io/ssh-auth`**: Must contain an `ssh-privatekey` key.
- **`kubernetes.io/tls`**: Must contain both `tls.crt` and `tls.key` keys.
