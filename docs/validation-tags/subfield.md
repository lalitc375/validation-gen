# +k8s:subfield

## Description
Targets a subfield of a struct for validation, allowing validation tags to be applied to nested fields.

## Scope
`Type`, `Field`

## Supported Go Types
Any Go type. This tag targets a field within a struct, regardless of its own type.

## Arguments
`name=<field-json-name>` (Required): The JSON name of the subfield to target.

## Payload
`+<validation-tag>`: The validation tag(s) to apply to the targeted subfield.

## Stability
**Stable**

## Usage

### Field
```go
type MyResource struct {
    // Targets the "name" subfield of the embedded metav1.ObjectMeta,
    // applying specific validation and marking it as optional.
    // +k8s:subfield(name="name")=+k8s:optional
    // +k8s:subfield(name="name")=+k8s:format=k8s-long-name
    metav1.ObjectMeta `json:"metadata,omitempty"`
}
```

### Type
```go
// Apply validation to a subfield of a type alias.
// +k8s:subfield(name="value")=+k8s:minLength=5
type MyCustomField struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

type MyStruct struct {
    Field MyCustomField `json:"field"`
}
```
This enables granular validation on complex nested structures without requiring extensive boilerplate code.

## Migrating from Handwritten Validation

The `+k8s:subfield` tag significantly simplifies validation of nested fields by allowing direct application of validation tags to specific sub-elements within a struct. This replaces handwritten validation logic that would otherwise require manually constructing `field.Path` objects and traversing nested structures to apply checks.

## Detailed Example: Validating a DeviceClass Metadata Name

This example demonstrates how `+k8s:subfield` is used to apply format validation (`+k8s:format=k8s-long-name`) and optionality to the `name` field within the embedded `metav1.ObjectMeta` of a `DeviceClass`.

### 1. Define the Tag in `types.go`
Apply the `+k8s:subfield` tags to the embedded `metav1.ObjectMeta` field in the `DeviceClass` struct.

**File:** `staging/src/k8s.io/api/resource/v1/types.go`
```go
type DeviceClass struct {
    metav1.TypeMeta `json:",inline"`
    // Standard object metadata
    // +optional
    // +k8s:subfield(name)=+k8s:optional // The name field of ObjectMeta is optional for DeviceClass
    // +k8s:subfield(name)=+k8s:format=k8s-long-name // The name must conform to k8s-long-name format
    metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

    // Spec defines what can be allocated and how to configure it.
    Spec DeviceClassSpec `json:"spec" protobuf:"bytes,2,name=spec"`
}
```

### 2. Update Handwritten Validation
In the corresponding handwritten validation function (e.g., for `DeviceClass`), remove or mark as covered any explicit logic that constructs paths to `metadata.name` and applies format checks.

**File:** `pkg/apis/resource/validation/validation.go`
```go
func ValidateDeviceClass(obj *resource.DeviceClass, fldPath *field.Path) field.ErrorList {
    allErrs := field.ErrorList{}

    // ... other validations ...

    // Original handwritten validation for metadata.name (example):
    // if obj.ObjectMeta.Name != "" { // Check if name is provided
    //     for _, msg := range validation.IsDNS1123Subdomain(obj.ObjectMeta.Name) {
    //         allErrs = append(allErrs, field.Invalid(fldPath.Child("metadata", "name"), obj.ObjectMeta.Name, msg))
    //     }
    // }

    // After adding +k8s:subfield(name)=+k8s:optional and +k8s:subfield(name)=+k8s:format=k8s-long-name,
    // the generated code handles this. If still present for backward compatibility:
    namePath := fldPath.Child("metadata", "name")
    if obj.ObjectMeta.Name != "" {
        for _, msg := range validation.IsDNS1123Subdomain(obj.ObjectMeta.Name) { // Assuming IsDNS1123Subdomain aligns with k8s-long-name
            allErrs = append(allErrs, field.Invalid(namePath, obj.ObjectMeta.Name, msg).MarkCoveredByDeclarative())
        }
    }
    // Note: The optionality is handled automatically by the generated code due to +k8s:subfield(name)=+k8s:optional.

    // ... other validations ...
    return allErrs
}
```
The `+k8s:subfield` tag makes the validation of nested fields explicit and easily readable in the API definition, reducing the verbosity of handwritten validation code.
