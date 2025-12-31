# Validation: Service

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Service name is required and must be a DNS subdomain. |
| `spec.ports` | `[]ServicePort` | `+k8s:optional` | Optional list of ports. Required unless the Service is headless or of type `ExternalName`. |
| `spec.selector` | `map[string]string` | `+k8s:optional` | Optional label selector for pods. |
| `spec.clusterIP` | `string` | `+k8s:optional`<br>`+k8s:immutable` | Optional cluster IP address. Immutable after it is set. |
| `spec.type` | `ServiceType` | `+k8s:optional`<br>`+k8s:enum=["ClusterIP", "NodePort", "LoadBalancer", "ExternalName"]` | Optional type determines how the Service is exposed. Defaults to `ClusterIP`. |
| `spec.externalName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional external reference (CNAME). Required if type is `ExternalName`. |
| `spec.sessionAffinity` | `ServiceAffinity` | `+k8s:optional`<br>`+k8s:enum=["None", "ClientIP"]` | Optional session affinity. Defaults to `None`. |
| `spec.externalTrafficPolicy` | `ServiceExternalTrafficPolicy` | `+k8s:optional`<br>`+k8s:enum=["Cluster", "Local"]` | Optional external traffic policy. Only for `NodePort` and `LoadBalancer`. |
| `spec.internalTrafficPolicy` | `*ServiceInternalTrafficPolicy` | `+k8s:optional`<br>`+k8s:enum=["Cluster", "Local"]` | Optional internal traffic policy. Defaults to `Cluster`. |

### ServicePort

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `name` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional name of the port. Required if more than one port is specified. |
| `protocol` | `Protocol` | `+k8s:optional`<br>`+k8s:enum=["TCP", "UDP", "SCTP"]` | Mandatory protocol if set. Defaults to `TCP`. |
| `port` | `int32` | `+k8s:required`<br>`+k8s:minimum=1`<br>`+k8s:maximum=65535` | Mandatory port number. |
| `targetPort` | `intstr.IntOrString` | `+k8s:optional` | Optional port to access on the pods. |
| `nodePort` | `int32` | `+k8s:optional`<br>`+k8s:minimum=1`<br>`+k8s:maximum=65535` | Optional node port. |
