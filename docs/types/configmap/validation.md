# Validation: ConfigMap

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | ConfigMap name is required and must be a DNS subdomain. |
| `immutable` | `*bool` | `+k8s:optional` | Optional; if set to true, prevents further updates to `data` and `binaryData`. |
| `data` | `map[string]string` | `+k8s:optional` | Optional map of configuration data. Keys must follow `IsConfigMapKey` rules. |
| `binaryData` | `map[string][]byte` | `+k8s:optional` | Optional map of binary configuration data. Keys must follow `IsConfigMapKey` rules. |
