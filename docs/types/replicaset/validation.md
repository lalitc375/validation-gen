# Validation: ReplicaSet

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | ReplicaSet name is required and must be a DNS subdomain. |
| `spec.replicas` | `int32` | `+k8s:required`<br>`+k8s:minimum=0` | Mandatory number of desired replicas. |
| `spec.selector` | `*metav1.LabelSelector` | `+k8s:required`<br>`+k8s:immutable` | Mandatory label selector. Must match pod template labels. Immutable after creation. |
| `spec.template` | `corev1.PodTemplateSpec` | `+k8s:required` | Mandatory template for the pods that will be created. `restartPolicy` must be `Always`. `activeDeadlineSeconds` is forbidden. |
| `spec.minReadySeconds` | `int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional minimum number of seconds for which a newly created pod should be ready without any of its container crashing. |
