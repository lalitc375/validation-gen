# +k8s:validateError

## Description
Causes a code generation error (compile-time). Used primarily for testing the validation generator itself, for example, to verify that a specific tag combination correctly triggers an error during code generation.

## Scope
`Field`, `Type`

## Supported Go Types
Any Go type. The purpose of this tag is to trigger a compile-time error, so the underlying type is not relevant.

## Usage

### Field
```go
type MyStruct struct {
    // This field will cause a code generation error.
    // +k8s:validateError
    ProblemField string `json:"problemField"`
}
```

### Type
```go
// This type definition will cause a code generation error.
// +k8s:validateError
type ProblemType string

type MyStruct struct {
    AnotherProblem ProblemType `json:"anotherProblem"`
}
```
This tag is intended for internal testing of the code generator and should not be used in production API definitions.

## Migrating from Handwritten Validation

The `+k8s:validateError` tag is a diagnostic tool for the validation code generator. It is **not** used in regular API definitions for end-user validation. Therefore, it does not have a "migration path" from handwritten validation in the conventional sense. Its purpose is to intentionally break code generation to test the generator's error handling.

## Detailed Example: Triggering a Generation Error

This example shows how to use `+k8s:validateError` to intentionally cause a code generation failure, which can be useful when developing and testing the validation generator itself.

### 1. Define a Struct with `+k8s:validateError`
Apply the `+k8s:validateError` tag to a field or a type within your test code.

**File:** `pkg/apis/example/v1/types_test.go` (hypothetical test file for validation generator)
```go
type InvalidStructForTest struct {
    // Attempting to generate validation code for this struct will result in a compile-time error.
    // +k8s:validateError
    ThisFieldShouldCauseAnError string `json:"thisFieldShouldCauseAnError"`
}
```
When the validation code generator processes `InvalidStructForTest`, it will encounter `+k8s:validateError` and, as designed, produce a code generation error, confirming its functionality for testing purposes.
