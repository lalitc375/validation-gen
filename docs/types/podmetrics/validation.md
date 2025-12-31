# Validation: PodMetrics

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Pod name is required and must be a DNS subdomain. |
| `timestamp` | `metav1.Time` | `+k8s:required` | Mandatory time when metrics were collected. |
| `window` | `metav1.Duration` | `+k8s:required` | Mandatory duration of the time interval from which metrics were collected. |
| `containers` | `[]ContainerMetrics` | `+k8s:required`<br>`+k8s:minItems=1` | Mandatory list of metrics for all containers in the pod. |
| `containers[].name` | `string` | `+k8s:required` | Mandatory name of the container. |
| `containers[].usage` | `corev1.ResourceList` | `+k8s:required` | Mandatory resource usage metrics of the container. |
