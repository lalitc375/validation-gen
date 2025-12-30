# +k8s:zeroOrOneOfMember

## Description
Defines a "loose" union where **at most one** member can be set (all unset is valid). This is useful for fields that represent alternative choices, where having zero choices or exactly one choice is valid.

## Scope
`Field`

## Supported Go Types
Any Go type. The tag indicates that this field is part of a group where at most one field can be set.

## Arguments
`union=<name>` (optional): Specifies the name of the union if a struct contains multiple independent "zero or one of" groups.

## Stability
**Stable**

## Usage

### Field
```go
type PodSchedulingGate struct {
    // This field is part of a "zero or one of" group. If set, no other fields in this group can be set.
    // +k8s:zeroOrOneOfMember
    // +k8s:optional
    Name string `json:"name,omitempty"`

    // This field is also part of the same group. If Name is set, this must be unset, and vice versa.
    // +k8s:zeroOrOneOfMember
    // +k8s:optional
    CustomGate *CustomGateConfig `json:"customGate,omitempty"`
}

type CustomGateConfig struct { /* ... */ }
```
In this example, either `Name` can be set, or `CustomGate` can be set, or neither can be set. However, both cannot be set simultaneously.

## Migrating from Handwritten Validation

The `+k8s:zeroOrOneOfMember` tag declaratively enforces an "at most one" constraint among a set of fields, including the case where all fields are unset. This directly replaces handwritten validation logic that would manually check for mutual exclusivity and generate `field.Forbidden` or `field.Invalid` errors if multiple fields were set.

This tag is particularly powerful when used in conjunction with `+k8s:item` to apply such a constraint to specific elements within a list.

## Detailed Example: Mutually Exclusive Certificate Conditions

This example demonstrates how `+k8s:zeroOrOneOfMember` is used with `+k8s:item` to ensure that a `CertificateSigningRequest` can have at most one of an "Approved" or "Denied" condition.

### 1. Define the Tags in `types.go`
Apply the `+k8s:item(type="...")=+k8s:zeroOrOneOfMember` to the `Conditions` field in `CertificateSigningRequestStatus`.

**File:** `staging/src/k8s.io/api/certificates/v1/types.go`
```go
type CertificateSigningRequestStatus struct {
    // ...
    // conditions applied to the request. Known conditions are "Approved", "Denied", and "Failed".
    // Approved and Denied conditions are mutually exclusive.
    // +listType=map
    // +k8s:listMapKey=type
    // +k8s:item(type: "Approved")=+k8s:zeroOrOneOfMember
    // +k8s:item(type: "Denied")=+k8s:zeroOrOneOfMember
    Conditions []CertificateSigningRequestCondition `json:"conditions,omitempty" protobuf:"bytes,1,rep,name=conditions"`
    // ...
}
```

### 2. Update Handwritten Validation
In the corresponding handwritten validation function for `CertificateSigningRequestStatus`, remove or mark as covered the explicit checks for mutual exclusivity between "Approved" and "Denied" conditions.

**File:** `pkg/apis/certificates/validation/validation.go`
```go
func validateConditions(fldPath *field.Path, csr *certificates.CertificateSigningRequest, opts certificateValidationOptions) field.ErrorList {
    allErrs := field.ErrorList{}
    // ...
    hasApproved := false
    hasDenied := false

    for i, c := range csr.Status.Conditions {
        // ... other validations ...

        // Original handwritten mutual exclusivity check (example):
        // if !opts.allowBothApprovedAndDenied { // 'opts' might be for compatibility
        //     switch c.Type {
        //     case certificates.CertificateApproved:
        //         hasApproved = true
        //         if hasDenied {
        //             allErrs = append(allErrs, field.Invalid(fldPath, c.Type, "Approved and Denied conditions are mutually exclusive"))
        //         }
        //     case certificates.CertificateDenied:
        //         hasDenied = true
        //         if hasApproved {
        //             allErrs = append(allErrs, field.Invalid(fldPath, c.Type, "Approved and Denied conditions are mutually exclusive"))
        //         }
        //     }
        // }

        // With +k8s:item(type="...")=+k8s:zeroOrOneOfMember, this check is generated automatically.
        // If still present for backward compatibility during migration, mark the error as covered:
        if !opts.allowBothApprovedAndDenied {
            switch c.Type {
            case certificates.CertificateApproved:
                hasApproved = true
                if hasDenied {
                    allErrs = append(allErrs, field.Invalid(fldPath, c.Type, "Approved and Denied conditions are mutually exclusive").WithOrigin("zeroOrOneOf").MarkCoveredByDeclarative())
                }
            case certificates.CertificateDenied:
                hasDenied = true
                if hasApproved {
                    allErrs = append(allErrs, field.Invalid(fldPath, c.Type, "Approved and Denied conditions are mutually exclusive").WithOrigin("zeroOrOneOf").MarkCoveredByDeclarative())
                }
            }
        }
    }

    return allErrs
}
```
The `+k8s:zeroOrOneOfMember` tag, especially when combined with `+k8s:item`, allows for precise and declarative definition of mutually exclusive choices, simplifying validation logic.
