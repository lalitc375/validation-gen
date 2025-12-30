# +k8s:update

## Description
Constrains how a field can be updated. This tag applies validation rules to field transitions during update operations.

## Scope
`Field`

## Supported Go Types
Any Go type. The tag applies constraints to changes in the field's value.

## Payload
`NoSet` | `NoUnset` | `NoModify`
*   `NoSet`: Cannot transition from a nil/zero value to a set (non-nil/non-zero) value.
*   `NoUnset`: Cannot transition from a set (non-nil/non-zero) value to a nil/zero value.
*   `NoModify`: The value of the field cannot change. However, it can transition between set/unset unless combined with `NoSet` or `NoUnset`.

## Stability
**Alpha**

## Usage

### Field
```go
type ResourceClaimStatus struct {
    // Once Allocation is set, its value cannot be modified.
    // +k8s:update=NoModify
    Allocation *AllocationResult `json:"allocation,omitempty"`
}
```
This tag provides fine-grained control over how specific fields are allowed to change throughout the lifecycle of an API object.

## Migrating from Handwritten Validation

The `+k8s:update` tag provides a declarative way to enforce rules about how fields can change during update operations, directly replacing handwritten validation logic that typically compares `old` and `new` versions of an object. This often involves using `apiequality.Semantic.DeepEqual` or similar comparisons for immutability, or explicit checks for transitions between unset/set states.

When migrating to `+k8s:update`, the corresponding handwritten logic can be removed or marked as covered.

## Detailed Example: Preventing Modification of ResourceClaim Allocation

This example demonstrates how `+k8s:update=NoModify` is applied to the `Allocation` field within `ResourceClaimStatus`, ensuring that once an allocation is made, its details cannot be changed.

### 1. Define the Tag in `types.go`
Apply the `+k8s:update=NoModify` tag to the `Allocation` field in the `ResourceClaimStatus` struct definition.

**File:** `staging/src/k8s.io/api/resource/v1/types.go`
```go
type ResourceClaimStatus {
    // Allocation is set once the claim has been allocated successfully.
    // Its value cannot be modified after it has been set.
    // +k8s:update=NoModify
    Allocation *AllocationResult `json:"allocation,omitempty" protobuf:"bytes,1,opt,name=allocation"`
    // ...
}
```

### 2. Update Handwritten Validation
In the corresponding handwritten validation function for `ResourceClaimStatus` updates, remove or mark as covered the explicit checks for `Allocation` immutability.

**File:** `pkg/apis/resource/validation/validation.go`
```go
func ValidateResourceClaimStatusUpdate(newClaim, oldClaim *resource.ResourceClaim, fldPath *field.Path) field.ErrorList {
    allErrs := field.ErrorList{}

    // ... other validations ...

    // Original handwritten validation for Allocation immutability (example):
    // if oldClaim.Status.Allocation != nil && newClaim.Status.Allocation != nil &&
    //    !apiequality.Semantic.DeepEqual(newClaim.Status.Allocation, oldClaim.Status.Allocation) {
    //    allErrs = append(allErrs, field.Forbidden(fldPath.Child("allocation"), "field is immutable after it has been set"))
    // }

    // After adding +k8s:update=NoModify, the generated code handles this.
    // If still present for backward compatibility during migration, mark the error as covered:
    if oldClaim.Status.Allocation != nil && newClaim.Status.Allocation != nil &&
       !apiequality.Semantic.DeepEqual(newClaim.Status.Allocation, oldClaim.Status.Allocation) {
       allErrs = append(allErrs, field.Forbidden(fldPath.Child("allocation"), "field is immutable after it has been set").MarkCoveredByDeclarative())
    }

    // ... other validations ...
    return allErrs
}
```
The `+k8s:update=NoModify` tag ensures that the `Allocation` field, once populated, cannot be changed on subsequent updates, providing clear API behavior and reducing custom validation code.
