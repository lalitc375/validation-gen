# +k8s:maximum

## Description
Specifies the maximum allowed value for an integer field.

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
type PriorityClass struct {
    // +k8s:maximum=1000000000
    Value int32 `json:"value"`
}
```

### Type
```go
// +k8s:maximum=65535
type Port int32

type ServicePort struct {
    Port Port `json:"port"`
}
```

### Map & Slice
To validate items in a map or slice, compose with `+k8s:eachVal` or `+k8s:eachKey`.

```go
type MyStruct struct {
    // Validates that each int in the slice is at most 100
    // +k8s:eachVal=+k8s:maximum=100
    Values []int `json:"values,omitempty"`
}
```

## Migrating from Handwritten Validation

When adding `+k8s:maximum` to a field that already has handwritten validation, follow this pattern:

1.  **Add the Tag**: Add the `+k8s:maximum=<value>` tag to the struct field.
2.  **Mark Errors**: Update the handwritten validation logic to mark the maximum value check error as covered by declarative validation using `.MarkCoveredByDeclarative()`.

## Detailed Example: Validating Port Range

This example demonstrates how to apply `maximum` validation to the `Port` field of a `ServicePort` struct.

### 1. Define the Tag in `types.go`
Add the `+k8s:maximum=65535` tag to the `Port` field in the `ServicePort` struct definition.

**File:** `staging/src/k8s.io/api/core/v1/types.go`
```go
type ServicePort struct {
    // ...
    // +k8s:maximum=65535
    Port int32 `json:"port" protobuf:"varint,3,opt,name=port"`
    // ...
}
```

### 2. Update Handwritten Validation
In the existing validation logic, mark the `Invalid` error for the `port` field as covered by declarative validation when the value exceeds the maximum.

**File:** `pkg/apis/core/validation/validation.go`

```go
func validateServicePort(sp *core.ServicePort, fldPath *field.Path) field.ErrorList {
    allErrs := field.ErrorList{}

    if sp.Port > 65535 {
        allErrs = append(allErrs, field.Invalid(fldPath.Child("port"), sp.Port, "must be less than or equal to 65535").MarkCoveredByDeclarative())
    }
    // ... other validation ...
    return allErrs
}
```
This ensures that the declarative validation tooling can verify the `maximum` constraint while maintaining compatibility with the existing handwritten validation.

## Test Coverage

When using `+k8s:maximum`, you should add declarative validation tests to verify that values larger than the specified maximum are rejected.

### Example

Suppose you have a field that must be at most 10.

**File:** `pkg/apis/example/v1/types.go`
```go
type MyStruct struct {
    // +k8s:maximum=10
    Count int32 `json:"count,omitempty"`
}
```

Your `declarative_validation_test.go` should include test cases for values that are less than, equal to, and greater than the maximum.

**File:** `pkg/apis/example/validation/declarative_validation_test.go`
```go
func TestDeclarativeValidateMaximum(t *testing.T) {
    // ...
    testCases := map[string]struct {
        input        example.MyStruct
        expectedErrs field.ErrorList
    }{
        "value less than maximum": {
            input: mkMyStruct(func(obj *example.MyStruct) {
                obj.Count = 5
            }),
            expectedErrs: field.ErrorList{},
        },
        "value equal to maximum": {
            input: mkMyStruct(func(obj *example.MyStruct) {
                obj.Count = 10
            }),
            expectedErrs: field.ErrorList{},
        },
        "value greater than maximum": {
            input: mkMyStruct(func(obj *example.MyStruct) {
                obj.Count = 11
            }),
            expectedErrs: field.ErrorList{
                field.Invalid(field.NewPath("spec", "count"), 11, "must be less than or equal to 10").WithOrigin("maximum"),
            },
        },
    }
    // ...
}
```

In this example:
1.  We test values that are less than, equal to, and greater than the specified `maximum`.
2.  For the value that is greater than the maximum, we expect a `field.Invalid` error.
3.  The error origin is `maximum`, corresponding to the `+k8s:maximum` tag.
