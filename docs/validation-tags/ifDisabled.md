# +k8s:ifDisabled

## Description
Applies the chained validation only if a specific feature gate/option is disabled.

## Scope
`Field`, `Type`

## Supported Go Types
Any Go type. This tag acts as a conditional wrapper for other validation tags applied to the field or type.

## Arguments
`<OptionName>` (Required): The name of the feature gate or option.

## Payload
`+<validation-tag>`: The validation tag(s) to apply if the option is disabled.

## Stability
**Alpha**

## Usage

### Field
```go
type MyStruct struct {
    // If "MyFeature" is disabled, this field is required.
    // +k8s:ifDisabled(MyFeature)=+k8s:required
    // +k8s:optional
    Config string `json:"config,omitempty"`
}
```

### Type
```go
// If "NewProtocolRollout" feature is disabled, "LegacyProtocol" is a valid enum value.
// +k8s:enum
// +k8s:ifDisabled(NewProtocolRollout)=+k8s:enumExclude
type Protocol string

const (
    NewProtocol    Protocol = "New"
    LegacyProtocol Protocol = "Legacy"
)
```
In the above example, if `NewProtocolRollout` is disabled, `LegacyProtocol` will be an invalid value for the `Protocol` enum.

## Migrating from Handwritten Validation

The `+k8s:ifDisabled` tag directly translates conditional logic based on feature gates into declarative validation. This eliminates the need for boilerplate `if featuregate.IsDisabled(...)` checks in handwritten validation functions, making the validation rules more explicit and maintainable.

## Detailed Example: Conditional Enum Exclusion

This example demonstrates how `+k8s:ifDisabled` can be used to conditionally exclude an enum value based on the state of a feature gate, replacing handwritten checks for feature status.

### 1. Define the Enum Type with Conditional Exclusion in `types.go`
Apply the `+k8s:ifDisabled` tag to a `const` value, chaining it with `+k8s:enumExclude`.

**File:** `pkg/apis/example/v1/types.go`
```go
// +k8s:enum
type FeatureDependentState string

const (
	StateEnabled  FeatureDependentState = "Enabled"
	StateDisabled FeatureDependentState = "Disabled"
	// This state is only valid if "MyFeatureGate" is enabled.
	// If "MyFeatureGate" is disabled, "ConditionalState" will not be a valid option.
	// +k8s:ifDisabled(MyFeatureGate)=+k8s:enumExclude
	ConditionalState FeatureDependentState = "Conditional"
)
```

### 2. Update Handwritten Validation (if applicable)
If there was prior handwritten validation that manually checked `MyFeatureGate`'s status to validate `ConditionalState`, that logic can now be removed or marked as covered.

**File:** `pkg/apis/example/validation/validation.go`
```go
func ValidateFeatureDependentState(state FeatureDependentState, fldPath *field.Path) field.ErrorList {
    allErrs := field.ErrorList{}

    // Original handwritten logic (example):
    // if state == ConditionalState && utilfeature.DefaultFeatureGate.Disabled(features.MyFeatureGate) {
    //     allErrs = append(allErrs, field.Invalid(fldPath, state, "ConditionalState is not allowed when MyFeatureGate is disabled"))
    // }

    // With +k8s:ifDisabled, the generated code handles this. If still present for backward compatibility:
    // This check would ideally be removed, or its error marked as covered by the declarative tag.
    // For simplicity, we'll show how it would be marked if kept during transition.
    // if state == ConditionalState && utilfeature.DefaultFeatureGate.Disabled(features.MyFeatureGate) {
    //     allErrs = append(allErrs, field.Invalid(fldPath, state, "ConditionalState is not allowed when MyFeatureGate is disabled").MarkCoveredByDeclarative())
    // }

    return allErrs
}
```
The `+k8s:ifDisabled` tag makes the conditional validation rule explicit in the API definition, reducing the need for imperative code.
