# Validation: ComponentStatus (Deprecated)

`ComponentStatus` is a deprecated API (since v1.19) and is read-only. It provides status information for control plane components.

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required` | Name of the component. |
| `conditions` | `[]ComponentCondition` | `+k8s:optional` | List of current conditions for the component. |

### ComponentCondition

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `type` | `ComponentConditionType` | `+k8s:required`<br>`+k8s:enum=["Healthy"]` | Mandatory condition type. |
| `status` | `ConditionStatus` | `+k8s:required`<br>`+k8s:enum=["True", "False", "Unknown"]` | Mandatory status of the condition. |
| `message` | `string` | `+k8s:optional` | Optional human-readable message. |
| `error` | `string` | `+k8s:optional` | Optional error message. |