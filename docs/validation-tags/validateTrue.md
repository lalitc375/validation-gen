# +k8s:validateTrue

## Description
Validation always succeeds. Used primarily for testing the validation generator. When this tag is encountered, the generated validation code will unconditionally return no validation errors.

## Scope
`Field`, `Type`

## Supported Go Types
Any Go type. The purpose of this tag is to force validation success, so the underlying type is not relevant.

## Usage

### Field
```go
type MyStruct struct {
    // This field will always pass validation, regardless of its value.
    // +k8s:validateTrue
    AlwaysPassingField string `json:"alwaysPassingField"`
}
```

### Type
```go
// Any field using this type will always pass validation.
// +k8s:validateTrue
type AlwaysPassingType string

type MyStruct struct {
    AnotherPassingField AlwaysPassingType `json:"anotherPassingField"`
}
```
This tag is intended for internal testing and debugging of the code generator and generated validation logic, not for use in production API definitions.

## Migrating from Handwritten Validation

The `+k8s:validateTrue` tag is a diagnostic tool for the validation code generator. It is **not** used in regular API definitions for end-user validation. Therefore, it does not have a "migration path" from handwritten validation in the conventional sense. Its purpose is to intentionally make validation always succeed to test the generator's behavior.

## Detailed Example: Forcing Validation Success

This example shows how to use `+k8s:validateTrue` to intentionally make validation always succeed for a field, which can be useful when developing and testing the validation generator itself.

### 1. Define a Struct with `+k8s:validateTrue`
Apply the `+k8s:validateTrue` tag to a field or a type within your test code.

**File:** `pkg/apis/example/v1/types_test.go` (hypothetical test file for validation generator)
```go
type PassingStructForTest struct {
    // This field will always pass validation, regardless of its content.
    // +k8s:validateTrue
    ThisFieldShouldAlwaysPass string `json:"thisFieldShouldAlwaysPass"`
}
```
When validation code is generated for `PassingStructForTest`, any value provided for `ThisFieldShouldAlwaysPass` will be considered valid, allowing the testing of scenarios where validation should explicitly succeed.

