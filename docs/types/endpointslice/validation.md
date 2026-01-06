# Validation: EndpointSlice

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | EndpointSlice name is required and must be a DNS subdomain. |
| `addressType` | `AddressType` | `+k8s:required`<br>`+k8s:immutable`<br> | Mandatory address type. Immutable after creation. |
| `endpoints` | `[]Endpoint` | `+k8s:required`<br>`+k8s:maxItems=1000`<br>`+k8s:listType=atomic` | List of endpoints. Max 1000 per slice. |
| `ports` | `[]EndpointPort` | `+k8s:optional`<br>`+k8s:maxItems=100`<br>`+k8s:listType=atomic` | List of network ports. |

### Endpoint

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `addresses` | `[]string` | `+k8s:required`<br>`+k8s:maxItems=100`<br>`+k8s:listType=set` | Mandatory IP addresses or FQDNs. |
| `hostname` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional DNS label hostname. |
| `nodeName` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional node hosting the endpoint. |
| `zone` | `*string` | `+k8s:optional` | Optional zone name. |
| `hints.forZones` | `[]ForZone` | `+k8s:optional`<br>`+k8s:maxItems=8`<br>`+k8s:listType=atomic` | Optional topology routing hints for zones. |
| `hints.forNodes` | `[]ForNode` | `+k8s:optional`<br>`+k8s:maxItems=8`<br>`+k8s:listType=atomic` | Optional topology routing hints for nodes. |

### EndpointPort

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `name` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional name. Must be unique within the list if set. |
| `port` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=1`<br>`+k8s:maximum=65535` | Optional port number. |
| `protocol` | `*api.Protocol` | `+k8s:optional`<br> | Mandatory protocol if set. Defaults to `TCP`. |
| `appProtocol` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-label-key` | Optional application protocol hint. |

## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `AddressType` | `+k8s:enum` |
| `Protocol` | `+k8s:enum` |