# Validation: Ingress

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name`<br>`+k8s:subfield(name=labels)=+k8s:eachKey=+k8s:format=k8s-label-key`<br>`+k8s:subfield(name=annotations)=+k8s:eachKey=+k8s:format=k8s-annotation-key` | Ingress name is required and must be a DNS subdomain. |
| `spec.ingressClassName` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional name of the IngressClass resource. |
| `spec.defaultBackend` | `*IngressBackend` | `+k8s:optional` | Optional default backend. Required if `rules` is not specified. |
| `spec.rules` | `[]IngressRule` | `+k8s:optional` | Optional list of host rules. |
| `spec.rules[].host` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional host name. Can be a precise DNS name or a wildcard (`*.foo.com`). IPs are not allowed. |
| `spec.rules[].http.paths` | `[]HTTPIngressPath` | `+k8s:required`<br>`+k8s:minItems=1` | Mandatory collection of paths that map requests to backends. |
| `spec.rules[].http.paths[].path` | `string` | `+k8s:optional` | Optional path match. Must begin with `/` if `pathType` is `Exact` or `Prefix`. |
| `spec.rules[].http.paths[].pathType` | `*PathType` | `+k8s:required`<br>`+k8s:enum=["Exact", "Prefix", "ImplementationSpecific"]` | Mandatory interpretation of the path matching. |
| `spec.rules[].http.paths[].backend` | `IngressBackend` | `+k8s:required` | Mandatory backend for the path. |
| `spec.tls` | `[]IngressTLS` | `+k8s:optional` | Optional TLS configuration. |
| `spec.tls[].hosts` | `[]string` | `+k8s:optional`<br>`+k8s:eachVal=+k8s:format=hostname` | Optional list of hosts included in the TLS certificate. |
| `spec.tls[].secretName` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional name of the secret used to terminate TLS traffic. |

### IngressBackend

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `service` | `*IngressServiceBackend` | `+k8s:unionMember`<br/>`+k8s:optional` | References a service as a backend. Mutually exclusive with `resource`. |
| `service.name` | `string` | `+k8s:required`<br>`+k8s:format=k8s-long-name` | Mandatory name of the referenced service. |
| `service.port.name` | `string` | `+k8s:unionMember`<br/>`+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional name of the port on the service. Mutually exclusive with `number`. |
| `service.port.number` | `int32` | `+k8s:unionMember`<br/>`+k8s:optional`<br>`+k8s:minimum=1`<br>`+k8s:maximum=65535` | Optional numerical port. Mutually exclusive with `name`. |
| `resource` | `*api.TypedLocalObjectReference` | `+k8s:unionMember`<br/>`+k8s:optional` | References another Kubernetes resource. Mutually exclusive with `service`. |
