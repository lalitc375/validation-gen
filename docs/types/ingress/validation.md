# Validation: Ingress

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name`<br>`+k8s:subfield(name=labels)=+k8s:eachKey=+k8s:format=k8s-label-key`<br>`+k8s:subfield(name=annotations)=+k8s:eachKey=+k8s:format=k8s-label-key` | Ingress name is required and must be a DNS subdomain. |
| `spec.ingressClassName` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional name of the IngressClass resource. |
| `spec.defaultBackend` | `*IngressBackend` | `+k8s:optional` | Optional default backend. Required if `rules` is not specified. |
| `spec.rules` | `[]IngressRule` | `+k8s:optional`<br>`+k8s:listType=atomic` | Optional list of host rules. |
| `spec.tls` | `[]IngressTLS` | `+k8s:optional`<br>`+k8s:listType=atomic` | Optional TLS configuration. |

### IngressRule

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `host` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional host name. Can be a precise DNS name or a wildcard (`*.foo.com`). |
| `http.paths` | `[]HTTPIngressPath` | `+k8s:required`<br>`+k8s:listType=atomic` | Mandatory collection of paths. |

### HTTPIngressPath

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `path` | `string` | `+k8s:optional` | Optional path match. |
| `pathType` | `*PathType` | `+k8s:required` | Mandatory path matching interpretation. |
| `backend` | `IngressBackend` | `+k8s:required` | Mandatory backend. |

### IngressBackend

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `service` | `*IngressServiceBackend` | `+k8s:unionMember`<br/>`+k8s:optional` | References a service as a backend. |
| `resource` | `*api.TypedLocalObjectReference` | `+k8s:unionMember`<br/>`+k8s:optional` | References another Kubernetes resource. |

### IngressServiceBackend

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `name` | `string` | `+k8s:required`<br>`+k8s:format=k8s-short-name` | Mandatory name of the referenced service. |
| `port` | `ServiceBackendPort` | `+k8s:required` | Port of the referenced service. |

### ServiceBackendPort

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `name` | `string` | `+k8s:unionMember`<br/>`+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional name of the port. |
| `number` | `int32` | `+k8s:unionMember`<br/>`+k8s:optional`<br>`+k8s:minimum=1`<br>`+k8s:maximum=65535` | Optional numerical port. |

### IngressTLS

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `hosts` | `[]string` | `+k8s:optional`<br>`+k8s:eachVal=+k8s:format=k8s-long-name`<br>`+k8s:listType=atomic` | Optional list of hosts. |
| `secretName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional secret name. |

## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `PathType` | `+k8s:enum` |