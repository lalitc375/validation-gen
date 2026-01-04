# Validation: Event

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | Event name is required and must be a DNS subdomain. |
| `eventTime` | `metav1.MicroTime` | `+k8s:required` | Mandatory time when this Event was first observed. |
| `reportingController` | `string` | `+k8s:optional`<br>`+k8s:format=k8s-label-key` | Optional name of the controller that emitted this Event. |
| `reportingInstance` | `string` | `+k8s:required`<br>`+k8s:maxLength=128` | Mandatory ID of the controller instance. |
| `action` | `string` | `+k8s:required`<br>`+k8s:maxLength=128` | Mandatory action taken/failed. |
| `reason` | `string` | `+k8s:required`<br>`+k8s:maxLength=128` | Mandatory machine-readable reason for the action. |
| `regarding` | `corev1.ObjectReference` | `+k8s:optional`<br>`+k8s:immutable` | The object this Event is about. Immutable after creation. |
| `type` | `string` | `+k8s:required`<br> | Mandatory type of the event. |
| `note` | `string` | `+k8s:optional`<br>`+k8s:maxLength=1024` | Optional human-readable description. |
| `series.count` | `int32` | `+k8s:optional`<br>`+k8s:minimum=2` | Mandatory if series is present. Number of occurrences. |
| `series.lastObservedTime` | `metav1.MicroTime` | `+k8s:optional` | Mandatory if series is present. Time of last occurrence. |
