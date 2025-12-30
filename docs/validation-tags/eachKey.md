# +k8s:eachKey

## Description
Applies a validation tag to every **key** in a map.

## Scope
`Field`, `Type`

## Supported Go Types
`map[K]V`, `*map[K]V` (and any alias of these types), where `K` is the type of the key being validated.

## Payload
`+<validation-tag>`

## Stability
**Alpha**

## Usage

### Field
```go
type MyStruct struct {
    // Validates that each key in the map has a max length of 32
    // +k8s:eachKey=+k8s:maxLength=32
    Labels map[string]string `json:"labels"`
}
```

### Type
```go
// +k8s:eachKey=+k8s:maxLength=10
type MyMapKey string

type MyMap map[MyMapKey]string

type MyStruct struct {
    Data MyMap `json:"data"`
}
```

## Migrating from Handwritten Validation

When applying `+k8s:eachKey` to a map field that previously had handwritten validation for its keys, update the handwritten logic to mark the specific key validation errors as covered by declarative validation.

## Detailed Example: Validating Map Keys for Max Length

This example demonstrates applying `+k8s:eachKey` with `+k8s:maxLength` to validate the keys of a map.

### 1. Define the Tag in `types.go`
Apply the `+k8s:eachKey` tag with a nested `+k8s:maxLength` to the map field.

**File:** `pkg/apis/example/v1/types.go`
```go
type MyResourceSpec struct {
    // Validates that each key in the map is a valid label key,
    // and specifically, has a maximum length of 63 characters.
    // +k8s:eachKey=+k8s:maxLength=63
    // +k8s:eachKey=+k8s:format=k8s-label-key
    Labels map[string]string `json:"labels,omitempty"`
}
```

### 2. Update Handwritten Validation
In the corresponding handwritten validation function, mark the `TooLong` error for map keys as covered by declarative validation.

**File:** `pkg/apis/example/validation/validation.go`
```go
func ValidateMyResourceSpec(spec *MyResourceSpec, fldPath *field.Path) field.ErrorList {
    allErrs := field.ErrorList{}

    for key, _ := range spec.Labels {
        // Original handwritten validation for key length:
        // if len(key) > 63 {
        //     allErrs = append(allErrs, field.TooLong(fldPath.Child("labels").Key(key), key, 63))
        // }

        // After adding +k8s:eachKey=+k8s:maxLength=63, mark the error as covered:
        if len(key) > 63 {
            allErrs = append(allErrs, field.TooLong(fldPath.Child("labels").Key(key), key, 63).MarkCoveredByDeclarative())
        }
        // ... other handwritten validation for label keys (e.g., format checks) ...
    }

    return allErrs
}
```
By marking the specific length validation error as covered, you ensure that duplicate errors are avoided when both declarative and handwritten validations are active.
