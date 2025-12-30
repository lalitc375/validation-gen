# +k8s:ifEnabled

## Description
Applies the chained validation only if a specific feature gate/option is enabled.

## Scope
`Field`, `Type`

## Supported Go Types
Any Go type. This tag acts as a conditional wrapper for other validation tags applied to the field or type.

## Arguments
`<OptionName>` (Required): The name of the feature gate or option.

## Payload
`+<validation-tag>`: The validation tag(s) to apply if the option is enabled.

## Stability
**Alpha**

## Usage

### Field
```go
type MyStruct struct {
    // If "MyFeature" is enabled, this field is required.
    // +k8s:ifEnabled(MyFeature)=+k8s:required
    // +k8s:optional
    Config string `json:"config,omitempty"`
}
```

### Type
```go
// If "AlphaFeature" is enabled, "AlphaValue" is a valid enum value.
// +k8s:enum
// +k8s:ifEnabled(AlphaFeature)=+k8s:enumExclude
type FeatureEnum string

const (
    DefaultValue FeatureEnum = "Default"
    // +k8s:ifEnabled(AlphaFeature)=+k8s:enumExclude
    AlphaValue   FeatureEnum = "Alpha"
)
```
## Migrating from Handwritten Validation

Similar to `+k8s:ifDisabled`, the `+k8s:ifEnabled` tag simplifies conditional validation based on feature gates. It allows you to express that certain validation rules only apply when a feature is active, moving this logic from imperative Go code to declarative API definitions.

## Detailed Example: Conditionally Required Field

This example demonstrates how `+k8s:ifEnabled` can make a field conditionally required when a specific feature gate is enabled.

### 1. Define the Tag in `types.go`
Apply the `+k8s:ifEnabled` tag to a field, chaining it with `+k8s:required`.

**File:** `pkg/apis/example/v1/types.go`
```go
type MyResourceSpec struct {
    // This field is required only if "ExperimentalFeature" is enabled.
    // +k8s:ifEnabled(ExperimentalFeature)=+k8s:required
    // +k8s:optional
    ExperimentalConfig string `json:"experimentalConfig,omitempty"`
    // ...
}
```

### 2. Update Handwritten Validation (if applicable)
If there was prior handwritten validation that manually checked `ExperimentalFeature`'s status to make `ExperimentalConfig` required, that logic can now be removed or marked as covered.

**File:** `pkg/apis/example/validation/validation.go`
```go
func ValidateMyResourceSpec(spec *MyResourceSpec, fldPath *field.Path) field.ErrorList {
    allErrs := field.ErrorList{}

    // Original handwritten logic (example):
    // if utilfeature.DefaultFeatureGate.Enabled(features.ExperimentalFeature) && len(spec.ExperimentalConfig) == 0 {
    //     allErrs = append(allErrs, field.Required(fldPath.Child("experimentalConfig"), "must be specified when ExperimentalFeature is enabled"))
    // }

    // With +k8s:ifEnabled, the generated code handles this. If still present for backward compatibility:
    // if utilfeature.DefaultFeatureGate.Enabled(features.ExperimentalFeature) && len(spec.ExperimentalConfig) == 0 {
    //     allErrs = append(allErrs, field.Required(fldPath.Child("experimentalConfig"), "must be specified when ExperimentalFeature is enabled").MarkCoveredByDeclarative())
    // }

    return allErrs
}
```
The `+k8s:ifEnabled` tag makes the conditional requirement explicit in the API definition, reducing the need for imperative checks in handwritten code.

