# Validation: Service

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name`<br>`+k8s:subfield(name=labels)=+k8s:eachKey=+k8s:format=k8s-label-key`<br>`+k8s:subfield(name=annotations)=+k8s:eachKey=+k8s:format=k8s-annotation-key` | Service name is required and must be a DNS subdomain. |
| `spec.ports` | `[]ServicePort` | `+k8s:optional`<br>`+k8s:listType=map`<br>`+k8s:listMapKey=port`<br>`+k8s:listMapKey=protocol` | Optional list of ports. |
| `spec.selector` | `map[string]string` | `+k8s:optional`<br>`+k8s:eachKey=+k8s:format=k8s-label-key` | Optional label selector for pods. |
| `spec.clusterIP` | `string` | `+k8s:optional`<br>`+k8s:immutable` | Optional cluster IP address. Immutable after it is set. |
| `spec.type` | `ServiceType` | `+k8s:optional`<br> | Optional type. Defaults to `ClusterIP`. |
| `spec.externalName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional external reference (CNAME). |
| `spec.sessionAffinity` | `ServiceAffinity` | `+k8s:optional`<br> | Optional session affinity. Defaults to `None`. |
| `spec.externalTrafficPolicy` | `ServiceExternalTrafficPolicy` | `+k8s:optional`<br> | Optional external traffic policy. |
| `spec.internalTrafficPolicy` | `*ServiceInternalTrafficPolicy` | `+k8s:optional`<br> | Optional internal traffic policy. Defaults to `Cluster`. |
| `spec.sessionAffinityConfig.clientIP` | `*ClientIPConfig` | `+k8s:optional` | Optional client IP affinity configuration. |

### ServicePort

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `name` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional port name. |
| `protocol` | `Protocol` | `+k8s:optional`<br> | Mandatory protocol if set. Defaults to `TCP`. |
| `port` | `int32` | `+k8s:required`<br>`+k8s:minimum=1`<br>`+k8s:maximum=65535` | Mandatory port number. |
| `targetPort` | `intstr.IntOrString` | `+k8s:optional` | Optional port to access on the pods. |
| `nodePort` | `int32` | `+k8s:optional`<br>`+k8s:minimum=1`<br>`+k8s:maximum=65535` | Optional node port. |
| `appProtocol` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-label-key` | Optional application protocol. |

### ClientIPConfig

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `timeoutSeconds` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=1`<br>`+k8s:maximum=86400` | Optional session affinity timeout. |

## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `Protocol` | `+k8s:enum` |
| `ServiceAffinity` | `+k8s:enum` |
| `ServiceExternalTrafficPolicy` | `+k8s:enum` |
| `ServiceInternalTrafficPolicy` | `+k8s:enum` |
| `ServiceType` | `+k8s:enum` |