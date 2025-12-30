# +k8s:minimum

## Description
Specifies the minimum allowed value for an integer field.

## Scope
`Field`, `Type`, `ListVal`, `MapKey`, `MapVal`

## Supported Go Types
`int`, `int16`, `int32`, `int64`, `uint`, `uint16`, `uint32`, `uint64` and their pointer types (and any alias of these types)

## Payload
`<integer>`

## Stability
**Stable**

## Usage

### Field
```go
type ReplicaSetSpec struct {
    // +k8s:minimum=0
    Replicas *int32 `json:"replicas,omitempty"`
}
```

### Type
```go
// +k8s:minimum=1
type Port int32

type ServicePort struct {
    Port Port `json:"port"`
}
```

### Map & Slice
To validate items in a map or slice, compose with `+k8s:eachVal` or `+k8s:eachKey`.

```go
type MyStruct struct {
    // Validates that each int in the slice is at least 1
    // +k8s:eachVal=+k8s:minimum=1
    Values []int `json:"values,omitempty"`
}
```

## Migrating from Handwritten Validation

When adding `+k8s:minimum` to a field that already has handwritten validation, follow this pattern:

1.  **Add the Tag**: Add the `+k8s:minimum=<value>` tag to the struct field.
2.  **Mark Errors**: Update the handwritten validation logic to mark the minimum value check error as covered by declarative validation using `.MarkCoveredByDeclarative()`.

## Detailed Example: Validating Replicas

This example demonstrates how to apply `minimum` validation to the `Replicas` field of a `ReplicaSetSpec` struct.

### 1. Define the Tag in `types.go`
Add the `+k8s:minimum=0` tag to the `Replicas` field in the `ReplicaSetSpec` struct definition.

**File:** `staging/src/k8s.io/api/apps/v1/types.go`
```go
type ReplicaSetSpec struct {
    // +k8s:minimum=0
    Replicas *int32 `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
    // ...
}
```

### 2. Update Handwritten Validation
In the existing validation logic for `ReplicaSetSpec`, mark the `Invalid` error for the `replicas` field as covered by declarative validation when the value is negative.

**File:** `pkg/apis/apps/validation/validation.go`

```go
func ValidateReplicaSetSpec(spec *apps.ReplicaSetSpec, fldPath *field.Path, opts apivalidation.PodValidationOptions) field.ErrorList {
    allErrs := field.ErrorList{}

    if spec.Replicas != nil && *spec.Replicas < 0 {
        allErrs = append(allErrs, field.Invalid(fldPath.Child("replicas"), *spec.Replicas, "must be non-negative").MarkCoveredByDeclarative())
    }
    // ... other validation ...
    return allErrs
}
```
This ensures that the declarative validation tooling can verify the `minimum` constraint while maintaining compatibility with the existing handwritten validation.
