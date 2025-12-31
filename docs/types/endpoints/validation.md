# Validation: Endpoints

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Endpoints name is required and must be a DNS subdomain. |
| `subsets` | `[]EndpointSubset` | `+k8s:optional` | Sets of IP addresses and ports. |
| `subsets[].addresses` | `[]EndpointAddress` | `+k8s:optional` | IP addresses which offer the related ports that are marked as ready. |
| `subsets[].addresses[].ip` | `string` | `+k8s:required`<br>`+k8s:format=k8s-ip` | Mandatory valid IP address. |
| `subsets[].addresses[].hostname` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional hostname for the endpoint. |
| `subsets[].addresses[].nodeName` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional node hosting the endpoint. |
| `subsets[].notReadyAddresses` | `[]EndpointAddress` | `+k8s:optional` | IP addresses which offer the related ports but are NOT yet marked as ready. |
| `subsets[].ports` | `[]EndpointPort` | `+k8s:optional` | Port numbers and protocols of the endpoints. |
| `subsets[].ports[].name` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional name of this port. Required if more than one port is specified. |
| `subsets[].ports[].port` | `int32` | `+k8s:required`<br>`+k8s:minimum=1`<br>`+k8s:maximum=65535` | Mandatory port number. |
| `subsets[].ports[].protocol` | `Protocol` | `+k8s:required`<br>`+k8s:enum=["TCP", "UDP", "SCTP"]` | Mandatory protocol. Defaults to `TCP`. |
| `subsets[].ports[].appProtocol` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-label-key` | Optional application protocol hint. |
