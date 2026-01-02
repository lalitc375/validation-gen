# Validation: StatefulSet

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | StatefulSet name is required and must be a DNS subdomain. |
| `spec.replicas` | `int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional number of desired replicas. Defaults to 1. |
| `spec.selector` | `*metav1.LabelSelector` | `+k8s:required`<br>`+k8s:immutable` | Mandatory label selector. Must match pod template labels. Immutable after creation. |
| `spec.template` | `corev1.PodTemplateSpec` | `+k8s:required` | Mandatory template for the pods that will be created. `restartPolicy` must be `Always`. `activeDeadlineSeconds` is forbidden. |
| `spec.serviceName` | `string` | `+k8s:required`<br>`+k8s:format=k8s-long-name` | Mandatory name of the service that governs this StatefulSet. |
| `spec.podManagementPolicy` | `PodManagementPolicyType` | `+k8s:optional`<br>`+k8s:enum=["OrderedReady", "Parallel"]` | Optional policy for creating pods. Defaults to `OrderedReady`. |
| `spec.updateStrategy.type` | `StatefulSetUpdateStrategyType` | `+k8s:optional`<br>`+k8s:enum=["RollingUpdate", "OnDelete"]` | Optional strategy for updates. Defaults to `RollingUpdate`. |
| `spec.updateStrategy.rollingUpdate.partition` | `int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional ordinal at which the set should be partitioned for updates. |
| `spec.volumeClaimTemplates` | `[]PersistentVolumeClaim` | `+k8s:optional` | Optional list of claims that pods are allowed to reference. |
| `spec.ordinals.start` | `int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional start ordinal. Defaults to 0. |
| `spec.minReadySeconds` | `int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional minimum seconds for which a pod should be ready. |
| `spec.revisionHistoryLimit` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=0` | Optional max revisions to maintain. Defaults to 10. |
