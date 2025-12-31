# Validation: PodTemplate

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | PodTemplate name is required and must be a DNS subdomain. |
| `template` | `PodTemplateSpec` | `+k8s:required` | Mandatory template for creating copies of a predefined pod. |
| `template.metadata` | `metav1.ObjectMeta` | `+k8s:optional` | Optional metadata of the pods created from this template. |
| `template.spec` | `PodSpec` | `+k8s:required` | Mandatory specification of the desired behavior of the pod. |
