# +k8s:validateFalse

## Description
Validation always fails. Used primarily for testing the validation generator. When this tag is encountered, the generated validation code will unconditionally return a validation error.

## Scope
`Field`, `Type`

## Supported Go Types
Any Go type. The purpose of this tag is to force a validation failure, so the underlying type is not relevant.

## Usage

### Field
```go
type MyStruct struct {
    // This field will always cause a validation error when set.
    // +k8s:validateFalse
    AlwaysFailingField string `json:"alwaysFailingField"`
}
```

### Type
```go
// Any field using this type will always cause a validation error.
// +k8s:validateFalse
type AlwaysFailingType string

type MyStruct struct {
    AnotherFailingField AlwaysFailingType `json:"anotherFailingField"`
}
```
This tag is intended for internal testing and debugging of the code generator and generated validation logic, not for use in production API definitions.

## Migrating from Handwritten Validation

The `+k8s:validateFalse` tag is a diagnostic tool for the validation code generator. It is **not** used in regular API definitions for end-user validation. Therefore, it does not have a "migration path" from handwritten validation in the conventional sense. Its purpose is to intentionally make validation fail to test the generator's behavior.

## Detailed Example: Forcing a Validation Failure

This example shows how to use `+k8s:validateFalse` to intentionally make validation fail for a field, which can be useful when developing and testing the validation generator itself.

### 1. Define a Struct with `+k8s:validateFalse`
Apply the `+k8s:validateFalse` tag to a field or a type within your test code.

**File:** `pkg/apis/example/v1/types_test.go` (hypothetical test file for validation generator)
```go
type FailingStructForTest struct {
    // Setting this field will always result in a validation error.
    // +k8s:validateFalse
    ThisFieldShouldAlwaysFail string `json:"thisFieldShouldAlwaysFail"`
}
```
When validation code is generated for `FailingStructForTest`, any attempt to set `ThisFieldShouldAlwaysFail` to any value will result in a validation error, allowing the testing of error handling paths.
