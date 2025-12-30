# +k8s:maxItems

## Description
Limits the maximum number of items in a list or array.

## Scope
`Field`, `Type`

## Supported Go Types
`[]any`, `*[]any` (and any alias of these types)

## Payload
`<non-negative integer>`

## Stability
**Stable**

## Usage

### Field
```go
type PodSpec struct {
    // +k8s:maxItems=1000
    HostAliases []HostAlias `json:"hostAliases,omitempty"`
}
```

### Type
```go
// +k8s:maxItems=10
type AliasSlice []string

type MyStruct struct {
    Aliases AliasSlice `json:"aliases,omitempty"`
}
```

## Migrating from Handwritten Validation

When adding `+k8s:maxItems` to a field that already has handwritten validation, follow this pattern:

1.  **Add the Tag**: Add the `+k8s:maxItems=<value>` tag to the struct field.
2.  **Mark Errors**: Update the handwritten validation logic to mark the item count check error as covered by declarative validation using `.MarkCoveredByDeclarative()`.

## Detailed Example: Validating HostAliases

This example demonstrates how to apply `maxItems` validation to the `HostAliases` field of a `PodSpec` struct.

### 1. Define the Tag in `types.go`
Add the `+k8s:maxItems=1000` tag to the `HostAliases` field in the `PodSpec` struct definition.

**File:** `staging/src/k8s.io/api/core/v1/types.go`
```go
type PodSpec struct {
    // +k8s:maxItems=1000
    HostAliases []HostAlias `json:"hostAliases,omitempty" patchStrategy:"merge" patchMergeKey:"ip" protobuf:"bytes,23,rep,name=hostAliases"`
    // ...
}
```

### 2. Update Handwritten Validation
In the existing validation logic for `PodSpec`, mark the `TooMany` error for the `hostAliases` field as covered by declarative validation.

**File:** `pkg/apis/core/validation/validation.go`

```go
func ValidatePodSpec(spec *core.PodSpec, fldPath *field.Path, opts PodValidationOptions) field.ErrorList {
    allErrs := field.ErrorList{}

    // ... other validation ...

    if len(spec.HostAliases) > 1000 {
        allErrs = append(allErrs, field.TooMany(fldPath.Child("hostAliases"), len(spec.HostAliases), 1000).MarkCoveredByDeclarative())
    }

    // ... other validation ...

    return allErrs
}
```
This ensures that the declarative validation tooling can verify the `maxItems` constraint while maintaining compatibility with the existing handwritten validation.
