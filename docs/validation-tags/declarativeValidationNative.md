# +k8s:declarativeValidationNative

## Description
Marker indicating that validation is purely declarative and has no handwritten equivalent. Affects error reporting.

## Scope
`Field`

## Supported Go Types
Any Go type for which other declarative validation tags are applied.

## Stability
**Stable**

## Usage

The `+k8s:declarativeValidationNative` tag is applied to a field. It must be used in conjunction with other stable validation tags.

### Field
```go
type MyStruct struct {
    // This field will be validated declaratively, and any errors will be reported directly.
    // +k8s:declarativeValidationNative
    // +k8s:required
    // +k8s:format=k8s-uuid
    ID string `json:"id"`
}
## Migrating from Handwritten Validation

The `+k8s:declarativeValidationNative` tag explicitly states that there is no handwritten validation for the field. If a field previously had handwritten validation, but is now solely validated declaratively, this tag should be added. This helps the validation system correctly attribute errors and ensures that the generated code does not attempt to mark non-existent handwritten errors as "covered."

## Detailed Example

Consider a scenario where an `ID` field has always been validated purely through declarative tags, without any corresponding handwritten Go validation.

### 1. Define the Tag in `types.go`
Apply the `+k8s:declarativeValidationNative` tag along with other declarative validation tags directly to the field.

**File:** `pkg/apis/example/v1/types.go`
```go
type MyResourceSpec struct {
    // This ID field is always validated declaratively, and has no handwritten equivalent.
    // +k8s:declarativeValidationNative
    // +k8s:required
    // +k8s:format=k8s-uuid
    ID string `json:"id"`
}
```

### 2. No Handwritten Validation Code
Because `+k8s:declarativeValidationNative` is present, there should be no corresponding handwritten validation function for this specific field (or its properties). Any errors will be directly reported by the declarative validation system.
