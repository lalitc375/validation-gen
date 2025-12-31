# Validation: Pod

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Pod name is required and must be a DNS subdomain. |
| `spec.containers` | `[]Container` | `+k8s:required`<br>`+k8s:minItems=1` | Mandatory list of containers belonging to the pod. |
| `spec.initContainers` | `[]Container` | `+k8s:optional` | Optional list of initialization containers. |
| `spec.ephemeralContainers` | `[]EphemeralContainer` | `+k8s:optional` | Optional list of ephemeral containers. Forbidden on creation. |
| `spec.restartPolicy` | `RestartPolicy` | `+k8s:optional`<br>`+k8s:enum=["Always", "OnFailure", "Never"]` | Optional restart policy for all containers. Defaults to `Always`. |
| `spec.terminationGracePeriodSeconds` | `*int64` | `+k8s:required`<br>`+k8s:minimum=0` | Mandatory duration in seconds for graceful termination. |
| `spec.activeDeadlineSeconds` | `*int64` | `+k8s:optional`<br>`+k8s:minimum=1` | Optional duration in seconds the pod may be active. |
| `spec.dnsPolicy` | `DNSPolicy` | `+k8s:optional`<br>`+k8s:enum=["ClusterFirstWithHostNet", "ClusterFirst", "Default", "None"]` | Optional DNS policy. Defaults to `ClusterFirst`. |
| `spec.nodeSelector` | `map[string]string` | `+k8s:optional` | Optional selector which must be true for the pod to fit on a node. |
| `spec.serviceAccountName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional name of the ServiceAccount to use. |
| `spec.nodeName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional node on which this pod is scheduled. Immutable once set. |
| `spec.affinity` | `*Affinity` | `+k8s:optional` | Optional scheduling constraints. |
| `spec.tolerations` | `[]Toleration` | `+k8s:optional` | Optional list of tolerations. |
| `spec.priorityClassName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional priority class name. |
| `spec.runtimeClassName` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional runtime class name. |
