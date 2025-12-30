# +k8s:item

## Description
Validates a specific item in a `listType=map` list. The item is selected by matching values of the keys.

## Scope
`Field`, `ListVal`

## Supported Go Types
`[]struct{...}` (slices of structs), where the struct has fields corresponding to `listMapKey`.

## Arguments
`key=<value>` (Required): A key-value pair where `key` matches one of the `listMapKey` fields, and `value` is the expected value of that field.

## Payload
`+<validation-tag>`: The validation tag(s) to apply to the matching item.

## Stability
**Stable**

## Usage

### Field
```go
type CertificateSigningRequestStatus struct {
    // ...
    // +k8s:listType=map
    // +k8s:listMapKey=type
    // +k8s:item(type="Approved")=+k8s:zeroOrOneOfMember
    // +k8s:item(type="Denied")=+k8s:zeroOrOneOfMember
    Conditions []CertificateSigningRequestCondition `json:"conditions,omitempty"`
}
```
In this example, specific validation (`+k8s:zeroOrOneOfMember`) is applied only to list items where the `type` field is "Approved" or "Denied".

## Migrating from Handwritten Validation

The `+k8s:item` tag simplifies conditional validation of specific elements within a `listType=map` slice. This replaces handwritten logic that would iterate through the slice, check the value of a key field for each element, and then apply specific validation rules.

## Detailed Example: Conditionally Validating Certificate Conditions

This example demonstrates using `+k8s:item` to apply `+k8s:zeroOrOneOfMember` validation selectively to "Approved" and "Denied" conditions within the `Conditions` slice of `CertificateSigningRequestStatus`.

### 1. Define the Tag in `types.go`
Apply `+k8s:item` tags to the `Conditions` field, specifying which items to target based on their `type` field.

**File:** `staging/src/k8s.io/api/certificates/v1/types.go`
```go
type CertificateSigningRequestStatus {
    // ...
    // +listType=map
    // +listMapKey=type
    // +k8s:item(type="Approved")=+k8s:zeroOrOneOfMember
    // +k8s:item(type="Denied")=+k8s:zeroOrOneOfMember
    Conditions []CertificateSigningRequestCondition `json:"conditions,omitempty" protobuf:"bytes,1,rep,name=conditions"`
}
```

### 2. Update Handwritten Validation
Update the handwritten validation function to remove or mark as covered the explicit conditional checks that `+k8s:item` now handles.

**File:** `pkg/apis/certificates/validation/validation.go`
```go
func validateConditions(fldPath *field.Path, csr *certificates.CertificateSigningRequest, opts certificateValidationOptions) field.ErrorList {
    allErrs := field.ErrorList{}
    // ...

    for i, c := range csr.Status.Conditions {
        // Original handwritten logic (simplified example for zeroOrOneOfMember):
        // if c.Type == certificates.CertificateApproved || c.Type == certificates.CertificateDenied {
        //     // Apply zeroOrOneOfMember validation logic here
        //     // ...
        // }

        // With +k8s:item, the generated code handles this. If still present for backward compatibility:
        // if (c.Type == certificates.CertificateApproved || c.Type == certificates.CertificateDenied) && !opts.allowBothApprovedAndDenied {
        //     // The specific zeroOrOneOfMember error for Approved/Denied conditions would be marked covered.
        //     allErrs = append(allErrs, field.Invalid(fldPath, c.Type, "Approved and Denied conditions are mutually exclusive").WithOrigin("zeroOrOneOf").MarkCoveredByDeclarative())
        // }
    }
    return allErrs
}
```
The `+k8s:item` tag allows for more concise and explicit definitions of validation rules for specific elements within a list, reducing the complexity of handwritten validation code.
