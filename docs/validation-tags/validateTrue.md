# +k8s:validation:validateTrue

## Description
Specifies that the value of the field must be `true`. This tag is used to enforce that a boolean field is always set to `true`.

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
    // This feature must be enabled.
    // +k8s:validation:validateTrue
    LicenseAccepted bool `json:"licenseAccepted"`
}
```
This tag is useful for fields that must be explicitly enabled or acknowledged, ensuring they cannot be set to `false`.

## Migrating from Handwritten Validation

If you have existing handwritten validation that checks if a boolean field is `true`, you can replace it with the `+k8s:validation:validateTrue` tag.

For example, if your old validation logic was:
```go
if !obj.LicenseAccepted {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("licenseAccepted"), obj.LicenseAccepted, "must be true"))
}
```
You can remove this code and use the declarative tag instead:
```go
// +k8s:validation:validateTrue
LicenseAccepted bool `json:"licenseAccepted"`
```
This simplifies the validation logic and makes the intent clear in the type definition.

## Test Coverage

To test the `+k8s:validation:validateTrue` tag, you should create test cases that check both the valid (`true`) and invalid (`false`) states of the boolean field.

### Example: Requiring Acceptance of Terms

**File:** `pkg/apis/example/v1/types.go`
```go
type UserPreferences struct {
    // The user must agree to the terms of service.
    // +k8s:validation:validateTrue
    TermsAgreed bool `json:"termsAgreed"`
}
```

**Test Cases:**
```go
// 1. Test with the field set to true (Valid)
validObj := &example.UserPreferences{TermsAgreed: true}
// expected: no validation error

// 2. Test with the field set to false (Invalid)
invalidObj := &example.UserPreferences{TermsAgreed: false}
// expected: field.Invalid(..., "must be true")
```
These tests verify that the `+k8s:validation:validateTrue` tag correctly enforces that the boolean field is always `true`.