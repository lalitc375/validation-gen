# Validation: NodeMetrics

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Node name is required and must be a DNS subdomain. |
| `timestamp` | `metav1.Time` | `+k8s:required` | Mandatory time when metrics were collected. |
| `window` | `metav1.Duration` | `+k8s:required` | Mandatory duration of the time interval from which metrics were collected. |
| `usage` | `corev1.ResourceList` | `+k8s:required` | Mandatory resource usage metrics of the node. |
