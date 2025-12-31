# Validation: ResourceQuota

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | ResourceQuota name is required and must be a DNS subdomain. |
| `spec.hard` | `ResourceList` | `+k8s:optional` | Optional set of desired hard limits for each named resource. Each quantity must be non-negative. |
| `spec.scopes` | `[]ResourceQuotaScope` | `+k8s:optional`<br>`+k8s:immutable`<br>`+k8s:eachVal=+k8s:enum=["Terminating", "NotTerminating", "BestEffort", "NotBestEffort", "PriorityClass", "CrossNamespacePodAffinity"]` | Optional collection of filters that must match each object tracked by a quota. Immutable after creation. |
| `spec.scopeSelector` | `*ScopeSelector` | `+k8s:optional` | Optional collection of filters like `scopes` but expressed using operators. |
| `spec.scopeSelector.matchExpressions` | `[]ScopedResourceSelectorRequirement` | `+k8s:optional` | List of scope selector requirements. |

### ScopedResourceSelectorRequirement

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `scopeName` | `ResourceQuotaScope` | `+k8s:required`<br>`+k8s:enum=["Terminating", "NotTerminating", "BestEffort", "NotBestEffort", "PriorityClass", "CrossNamespacePodAffinity"]` | Mandatory name of the scope. |
| `operator` | `ScopeSelectorOperator` | `+k8s:required`<br>`+k8s:enum=["In", "NotIn", "Exists", "DoesNotExist"]` | Mandatory relationship to a set of values. |
| `values` | `[]string` | `+k8s:optional` | Required and must be non-empty if operator is `In` or `NotIn`. Must be empty if operator is `Exists` or `DoesNotExist`. |
