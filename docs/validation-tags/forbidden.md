# +k8s:forbidden

## Description
Indicates that a field must NOT be specified. Useful for fields that are deprecated or reserved for future use.

## Scope
`Field`

## Supported Go Types
Any Go type. The tag indicates that the field itself should not be present, regardless of its type.

## Stability
**Alpha**

## Usage

### Field
```go
type Status struct {
    // +k8s:forbidden
    // This field is deprecated and should not be used.
    LegacyStatus string `json:"legacyStatus,omitempty"`
}
```
If a user attempts to set `LegacyStatus` in the YAML/JSON for `Status`, validation will fail.

## Migrating from Handwritten Validation

The `+k8s:forbidden` tag is a direct declarative equivalent to reporting `field.Forbidden` errors in handwritten validation. When migrating, you can remove the explicit `field.Forbidden` check for the field and rely on the tag. If the field needs to be conditionally forbidden, combine `+k8s:forbidden` with conditional tags like `+k8s:ifEnabled` or `+k8s:ifDisabled`.

## Detailed Example: Forbidding a Deprecated Field

This example demonstrates how to use `+k8s:forbidden` to disallow a field that is no longer supported.

### 1. Define the Tag in `types.go`
Apply the `+k8s:forbidden` tag to the deprecated field in your API struct.

**File:** `pkg/apis/example/v1/types.go`
```go
type PodSpec struct {
    // ...
    // +k8s:forbidden
    // This field is deprecated and will be removed in future versions.
    DeprecatedField string `json:"deprecatedField,omitempty"`
    // ...
}
```

### 2. Update Handwritten Validation
In the corresponding handwritten validation function, remove or mark as covered any explicit `field.Forbidden` checks for `DeprecatedField`.

**File:** `pkg/apis/core/validation/validation.go`
```go
func ValidatePodSpec(spec *core.PodSpec, fldPath *field.Path, opts PodValidationOptions) field.ErrorList {
    allErrs := field.ErrorList{}

    // ... other validations ...

    // Original handwritten validation for DeprecatedField:
    // if len(spec.DeprecatedField) > 0 {
    //     allErrs = append(allErrs, field.Forbidden(fldPath.Child("deprecatedField"), "deprecatedField is no longer supported"))
    // }

    // With +k8s:forbidden, the generated code handles this, so the handwritten check can be removed or marked covered.
    // If still present for backward compatibility during migration:
    if len(spec.DeprecatedField) > 0 {
        allErrs = append(allErrs, field.Forbidden(fldPath.Child("deprecatedField"), "deprecatedField is no longer supported").MarkCoveredByDeclarative())
    }

    // ... other validations ...
    return allErrs
}
```
The `+k8s:forbidden` tag will automatically generate the appropriate validation logic, ensuring that any attempt to set `DeprecatedField` will result in a validation error.
