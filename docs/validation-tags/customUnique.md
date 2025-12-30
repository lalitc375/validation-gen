# +k8s:customUnique

## Description
Disables generated uniqueness validation, deferring to custom handwritten validation logic.

## Scope
`Field`, `Type`

## Supported Go Types
`[]any`, `*[]any` (and any alias of these types)

## Stability
**Alpha**

## Usage

The `+k8s:customUnique` tag is applied to a field of type slice, and must be used in conjunction with `+k8s:listType`.

### Field
```go
type CertificateSigningRequestStatus struct {
    // +k8s:listType=map
    // +k8s:listMapKey=type
    // +k8s:customUnique
    Conditions []CertificateSigningRequestCondition `json:"conditions,omitempty"`
}
```

### Type
```go
// +k8s:listType=set
// +k8s:customUnique
type MyCustomUniqueList []string

type MyStruct struct {
    List MyCustomUniqueList `json:"list,omitempty"`
}
```

## Migrating from Handwritten Validation

The `+k8s:customUnique` tag explicitly disables generated uniqueness validation, signaling that custom handwritten logic will handle uniqueness checks. When migrating, ensure your handwritten validation correctly implements the desired uniqueness rules and, if necessary, marks errors as covered using `.MarkCoveredByDeclarative()`.

## Detailed Example: Custom Unique Conditions

This example demonstrates how `+k8s:customUnique` is used on a slice field, where uniqueness is handled by custom validation logic.

### 1. Define the Tag in `types.go`
Add the `+k8s:customUnique` tag to the `Conditions` field in the `CertificateSigningRequestStatus` struct definition, along with `+k8s:listType=map` and `+k8s:listMapKey=type`.

**File:** `staging/src/k8s.io/api/certificates/v1/types.go`
```go
type CertificateSigningRequestStatus struct {
    // ...
    // +k8s:listType=map
    // +k8s:listMapKey=type
    // +k8s:customUnique
    Conditions []CertificateSigningRequestCondition `json:"conditions,omitempty" protobuf:"bytes,1,rep,name=conditions"`
}
```

### 2. Implement Handwritten Uniqueness Validation
Implement or update the handwritten validation function to check for unique condition types. This function will be responsible for reporting duplicate errors.

**File:** `pkg/apis/certificates/validation/validation.go`

```go
func validateConditions(fldPath *field.Path, csr *certificates.CertificateSigningRequest, opts certificateValidationOptions) field.ErrorList {
    allErrs := field.ErrorList{}
    seenTypes := map[certificates.RequestConditionType]bool{}

    for i, c := range csr.Status.Conditions {
        // ... other validations for condition fields ...

        // This custom logic checks for duplicate condition types
        // The +k8s:customUnique tag ensures that the code generator
        // does not add its own uniqueness check for this field.
        if !opts.allowDuplicateConditionTypes { // This option is for backward compatibility
            if seenTypes[c.Type] {
                allErrs = append(allErrs, field.Duplicate(fldPath.Index(i).Child("type"), c.Type))
            }
            seenTypes[c.Type] = true
        }
    }
    return allErrs
}
```
In this example, the `+k8s:customUnique` tag ensures that the generated code skips uniqueness validation for the `Conditions` field, allowing the `validateConditions` function to manage it entirely.
