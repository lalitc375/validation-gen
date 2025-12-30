# +k8s:validation:validateFalse

## Description
Specifies that the value of the field must be `false`. This tag is used to enforce that a boolean field is always set to `false`.

## Scope
`Field`

## Supported Go Types
`bool`

## Stability
**Alpha**

## Usage

### Field
```go
type MyObject struct {
    // This feature is disabled and must be set to false.
    // +k8s:validation:validateFalse
    FeatureEnabled bool `json:"featureEnabled"`
}
```
This tag is useful for fields that are intended to be disabled or are not yet implemented, ensuring that they cannot be set to `true`.

## Migrating from Handwritten Validation

If you have existing handwritten validation that checks if a boolean field is `false`, you can replace it with the `+k8s:validation:validateFalse` tag.

For example, if your old validation logic was:
```go
if obj.FeatureEnabled {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("featureEnabled"), obj.FeatureEnabled, "must be false"))
}
```
You can remove this code and use the declarative tag instead:
```go
// +k8s:validation:validateFalse
FeatureEnabled bool `json:"featureEnabled"`
```
This simplifies the validation logic and makes the intent clear in the type definition.

## Test Coverage

To test the `+k8s:validation:validateFalse` tag, you should create test cases that check both the valid (`false`) and invalid (`true`) states of the boolean field.

### Example: Enforcing a Feature Flag to be Disabled

**File:** `pkg/apis/example/v1/types.go`
```go
type FeatureFlags struct {
    // This feature is currently experimental and must be disabled.
    // +k8s:validation:validateFalse
    EnableAlphaFeature bool `json:"enableAlphaFeature"`
}
```

**Test Cases:**
```go
// 1. Test with the field set to false (Valid)
validObj := &example.FeatureFlags{EnableAlphaFeature: false}
// expected: no validation error

// 2. Test with the field set to true (Invalid)
invalidObj := &example.FeatureFlags{EnableAlphaFeature: true}
// expected: field.Invalid(..., "must be false")
```
These tests verify that the `+k8s:validation:validateFalse` tag correctly enforces that the boolean field is always `false`.