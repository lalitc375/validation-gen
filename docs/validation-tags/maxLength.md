# +k8s:maxLength

## Description
Specifies the maximum length (in characters) for a string field.

## Scope
`Field`, `Type`, `ListVal`, `MapKey`, `MapVal`

## Supported Go Types
`string`, `*string` (and any alias of these types)

## Payload
`<non-negative integer>`

## Stability
**Stable**

## Usage

### Field
```go
type ObjectMeta struct {
    // +k8s:maxLength=253
    Name string `json:"name,omitempty"`
}
```

### Type
```go
// +k8s:maxLength=63
type LabelValue string

type ObjectMeta struct {
    Labels map[string]LabelValue `json:"labels,omitempty"`
}
```

### Map & Slice
To validate items in a map or slice, compose with `+k8s:eachVal` or `+k8s:eachKey`.

```go
type MyStruct struct {
    // Validates that each string in the slice has a max length of 8
    // +k8s:eachVal=+k8s:maxLength=8
    Values []string `json:"values,omitempty"`
}
```

## Migrating from Handwritten Validation

When adding `+k8s:maxLength` to a field that already has handwritten validation, follow this pattern:

1.  **Add the Tag**: Add the `+k8s:maxLength=<value>` tag to the struct field.
2.  **Mark Errors**: Update the handwritten validation logic to mark the length-check error as covered by declarative validation using `.MarkCoveredByDeclarative()`.

## Detailed Example: Validating a Name Length

This example demonstrates how to apply `maxLength` validation to the `Name` field of an `ObjectMeta` struct.

### 1. Define the Tag in `types.go`
Add the `+k8s:maxLength=253` tag to the `Name` field in the `ObjectMeta` struct definition.

**File:** `staging/src/k8s.io/apimachinery/pkg/apis/meta/v1/types.go`
```go
type ObjectMeta struct {
    // +k8s:maxLength=253
    Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
    // ...
}
```

### 2. Update Handwritten Validation
In the existing validation logic for `ObjectMeta`, mark the `TooLong` error for the `name` field as covered by declarative validation.

**File:** `staging/src/k8s.io/apimachinery/pkg/api/validation/objectmeta.go`

```go
func ValidateObjectMeta(meta *metav1.ObjectMeta, requiresNamespace bool, nameFn ValidateNameFunc, fldPath *field.Path) field.ErrorList {
    allErrs := field.ErrorList{}
    if len(meta.Name) == 0 {
        // ... (required check)
    } else {
        if len(meta.Name) > validation.DNS1123SubdomainMaxLength {
            allErrs = append(allErrs, field.TooLong(fldPath.Child("name"), meta.Name, validation.DNS1123SubdomainMaxLength).MarkCoveredByDeclarative())
        }
        for _, msg := range nameFn(meta.Name, false) {
            allErrs = append(allErrs, field.Invalid(fldPath.Child("name"), meta.Name, msg))
        }
    }
    // ... other validation ...
    return allErrs
}
```
This ensures that the declarative validation tooling can verify the `maxLength` constraint while maintaining compatibility with the existing handwritten validation.
