# Validation: NetworkPolicy

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name`<br>`+k8s:subfield(name=labels)=+k8s:eachKey=+k8s:format=k8s-label-key`<br>`+k8s:subfield(name=annotations)=+k8s:eachKey=+k8s:format=k8s-annotation-key` | NetworkPolicy name is required and must be a DNS subdomain. |
| `spec.podSelector` | `metav1.LabelSelector` | `+k8s:required` | Mandatory selector for the pods to which this NetworkPolicy applies. |
| `spec.ingress` | `[]NetworkPolicyIngressRule` | `+k8s:optional`<br>`+k8s:listType=atomic`<br>`+k8s:eachVal=+k8s:subfield(name=ports)=+k8s:eachVal=+k8s:subfield(name=port)=+k8s:minimum=1`<br>`+k8s:eachVal=+k8s:subfield(name=ports)=+k8s:eachVal=+k8s:subfield(name=port)=+k8s:maximum=65535` | Optional list of ingress rules to be applied to the selected pods. |
| `spec.egress` | `[]NetworkPolicyEgressRule` | `+k8s:optional`<br>`+k8s:listType=atomic`<br>`+k8s:eachVal=+k8s:subfield(name=ports)=+k8s:eachVal=+k8s:subfield(name=port)=+k8s:minimum=1`<br>`+k8s:eachVal=+k8s:subfield(name=ports)=+k8s:eachVal=+k8s:subfield(name=port)=+k8s:maximum=65535` | Optional list of egress rules to be applied to the selected pods. |
| `spec.policyTypes` | `[]PolicyType` | `+k8s:optional`<br>`+k8s:listType=atomic`<br>`+k8s:maxItems=2`<br>`+k8s:eachVal=` | Optional list of rule types that the NetworkPolicy relates to. |

### NetworkPolicyIngressRule / NetworkPolicyEgressRule

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `ports` | `[]NetworkPolicyPort` | `+k8s:optional` | Optional list of ports which should be made accessible. |
| `ports[].protocol` | `*api.Protocol` | `+k8s:optional`<br> | Mandatory protocol if set. Defaults to `TCP`. |
| `ports[].port` | `*intstr.IntOrString` | `+k8s:optional` | Optional port number or name. |
| `ports[].endPort` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=1`<br>`+k8s:maximum=65535` | Optional end port for a range. Must be >= `port`. |
| `from` / `to` | `[]NetworkPolicyPeer` | `+k8s:optional` | Optional list of sources/destinations. |

### NetworkPolicyPeer

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `podSelector` | `*metav1.LabelSelector` | `+k8s:unionMember(union="PeerType")`<br/>`+k8s:optional` | Optional selector for pods. |
| `namespaceSelector` | `*metav1.LabelSelector` | `+k8s:unionMember(union="PeerType")`<br/>`+k8s:optional` | Optional selector for namespaces. |
| `ipBlock` | `*IPBlock` | `+k8s:unionMember(union="PeerType")`<br/>`+k8s:optional` | Optional IP range. Mutually exclusive with `podSelector` and `namespaceSelector`. |
| `ipBlock.cidr` | `string` | `+k8s:required`<br>`+k8s:format=k8s-cidr` | Mandatory CIDR range. |
| `ipBlock.except` | `[]string` | `+k8s:optional`<br>`+k8s:eachVal=+k8s:format=k8s-cidr` | Optional list of CIDRs to exclude. Must be subsets of `cidr`. |


## Types

| Go Type | Validation Tags |
| :--- | :--- |
| `PolicyType` | `+k8s:enum` |
| `Protocol` | `+k8s:enum` |
