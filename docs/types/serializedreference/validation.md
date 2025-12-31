# Validation: SerializedReference

`SerializedReference` is a legacy type used to wrap an `ObjectReference`. It is largely deprecated in favor of direct use of `ObjectReference`.

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `reference` | `ObjectReference` | `+k8s:optional` | Optional reference to an object. |
