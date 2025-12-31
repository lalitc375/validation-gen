# Validation: ReplicationController

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | ReplicationController name is required and must be a DNS subdomain. |
| `spec.replicas` | `*int32` | `+k8s:required`<br>`+k8s:minimum=0` | Mandatory number of desired replicas. |
| `spec.selector` | `map[string]string` | `+k8s:required` | Mandatory label query over pods. Must match pod template labels. |
| `spec.template` | `*PodTemplateSpec` | `+k8s:required` | Mandatory template for the pods that will be created. `restartPolicy` must be `Always`. `activeDeadlineSeconds` is forbidden. |
| `spec.minReadySeconds` | `int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional minimum number of seconds for which a newly created pod should be ready without any of its container crashing. |
