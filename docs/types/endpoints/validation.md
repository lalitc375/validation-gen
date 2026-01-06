# Validation: Endpoints

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name`<br>`+k8s:subfield(name=labels)=+k8s:eachKey=+k8s:format=k8s-label-key`<br>`+k8s:subfield(name=annotations)=+k8s:eachKey=+k8s:format=k8s-label-key` | Endpoints name is required and must be a DNS subdomain. |
| `subsets` | `[]EndpointSubset` | `+k8s:optional`<br>`+k8s:listType=atomic` | Sets of IP addresses and ports. |

### EndpointSubset

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `addresses` | `[]EndpointAddress` | `+k8s:optional`<br>`+k8s:listType=atomic` | IP addresses which offer the related ports that are marked as ready. |
| `notReadyAddresses` | `[]EndpointAddress` | `+k8s:optional`<br>`+k8s:listType=atomic` | IP addresses which offer the related ports but are NOT yet marked as ready. |
| `ports` | `[]EndpointPort` | `+k8s:optional`<br>`+k8s:listType=atomic` | Port numbers and protocols of the endpoints. |

### EndpointAddress

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `ip` | `string` | `+k8s:required`<br>`+k8s:format=k8s-ip` | Mandatory valid IP address. |
| `hostname` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional hostname for the endpoint. |
| `nodeName` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-long-name` | Optional node hosting the endpoint. |

### EndpointPort

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `name` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-short-name` | Optional name of this port. Required if more than one port is specified. |
| `port` | `int32` | `+k8s:required`<br>`+k8s:minimum=1`<br>`+k8s:maximum=65535` | Mandatory port number. |
| `protocol` | `Protocol` | `+k8s:required` | Mandatory protocol. Defaults to `TCP`. |
| `appProtocol` | `*string` | `+k8s:optional`<br>`+k8s:format=k8s-label-key` | Optional application protocol hint. |

## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `Protocol` | `+k8s:enum` |