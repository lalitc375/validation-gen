# Validation: ClusterRole

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | ClusterRole name is required and must be a DNS subdomain. |
| `rules` | `[]PolicyRule` | `+k8s:required`<br>`+k8s:minItems=1` | Mandatory list of rules that apply to resources or non-resource URLs. |
| `aggregationRule` | `*AggregationRule` | `+k8s:optional` | Optional rule that describes how to build the rules for this ClusterRole. |

### AggregationRule

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `clusterRoleSelectors` | `[]metav1.LabelSelector` | `+k8s:required`<br>`+k8s:minItems=1` | Mandatory list of selectors used to find ClusterRoles and create the rules. |