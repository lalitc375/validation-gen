# +k8s:immutable

## Description
The field cannot be changed after creation.

## Scope
`Field`, `Type`

## Supported Go Types
Any Go type. The tag indicates that the field's value cannot be modified once set.

## Stability
**Alpha**

## Usage

### Field
```go
type Spec struct {
    // This field cannot be changed after creation.
    // +k8s:immutable
    ClusterIP string `json:"clusterIP"`
}
```

### Type
```go
// This type, and thus any field using it, cannot be changed after creation.
// +k8s:immutable
type ImmutableConfig struct {
    Value string `json:"value"`
}

type MyResourceSpec struct {
    Config ImmutableConfig `json:"config"`
}
```
If a field is a pointer type (e.g., `*string`), `+k8s:immutable` prevents changes to the pointed-to value, but it can still transition between `nil` and a non-`nil` value (unless combined with `+k8s:update=NoSet` or `+k8s:update=NoUnset`).

## Migrating from Handwritten Validation

The `+k8s:immutable` tag declaratively enforces that a field cannot be changed after creation. This directly replaces handwritten validation logic that would compare the new value of a field against its old value during an update operation and report an error if they differ.

## Detailed Example: Immutable ResourceClaim Specification

This example demonstrates applying `+k8s:immutable` to the `Spec` field of a `ResourceClaim` object, ensuring that its specification cannot be altered after creation.

### 1. Define the Tag in `types.go`
Apply the `+k8s:immutable` tag to the `Spec` field in the `ResourceClaim` struct definition.

**File:** `staging/src/k8s.io/api/resource/v1/types.go`
```go
type ResourceClaim struct {
    // ...
    // Spec describes what is being requested and how to configure it.
    // The spec is immutable.
    // +k8s:immutable
    Spec ResourceClaimSpec `json:"spec" protobuf:"bytes,2,name=spec"`
    // ...
}
```

### 2. Update Handwritten Validation
In the corresponding handwritten validation function for `ResourceClaim` updates, remove or mark as covered the explicit checks for `Spec` immutability.

**File:** `pkg/apis/resource/validation/validation.go`
```go
func ValidateResourceClaimUpdate(newClaim, oldClaim *resource.ResourceClaim) field.ErrorList {
    allErrs := field.ErrorList{}

    // ... other validations ...

    // Original handwritten validation for immutable Spec:
    // allErrs = append(allErrs, apivalidation.ValidateImmutableField(newClaim.Spec, oldClaim.Spec, field.NewPath("spec"))...)

    // After adding +k8s:immutable, the generated code handles this.
    // If still present for backward compatibility during migration:
    allErrs = append(allErrs, apivalidation.ValidateImmutableField(newClaim.Spec, oldClaim.Spec, field.NewPath("spec")).MarkCoveredByDeclarative()...)

    // ... other validations ...
    return allErrs
}
```
By using `+k8s:immutable`, the API server automatically rejects any attempts to modify the `Spec` field after the initial creation, enforcing the immutability rule declaratively.
