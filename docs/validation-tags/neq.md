# +k8s:neq

## Description
Verifies that the field's value is **not equal** to a specific value.

## Scope
`Field`, `Type`, `ListVal`, `MapKey`, `MapVal`

## Supported Go Types
`string`, `int`, `bool`, `float64` (and their pointer types and any aliases).

## Payload
`<value>`: The value that the field must not be equal to (string, int, bool, or float64 literal).

## Stability
**Alpha**

## Usage

### Field
```go
type MyResourceSpec struct {
    // The "ForbiddenValue" is not allowed for this field.
    // +k8s:neq="ForbiddenValue"
    ConfigOption string `json:"configOption"`
}
```

### Type
```go
// This type cannot hold the value "NotAllowed".
// +k8s:neq="NotAllowed"
type RestrictedString string

type MyResourceSpec struct {
    RestrictedField RestrictedString `json:"restrictedField"`
}
```

### Chaining
`+k8s:neq` can be chained with `+k8s:eachVal` or `+k8s:eachKey` to apply the check to elements of lists or keys of maps.

```go
type MyResourceSpec struct {
    // None of the tags in this slice can be "ForbiddenTag".
    // +k8s:eachVal=+k8s:neq="ForbiddenTag"
    Tags []string `json:"tags"`
}
```

## Migrating from Handwritten Validation

The `+k8s:neq` tag directly replaces handwritten validation logic that explicitly checks if a field's value is not equal to a specific forbidden value. This simplifies the validation code and makes the constraint clear in the API definition.

## Detailed Example: Forbidding a Specific Scheduler Name

This example demonstrates how `+k8s:neq` can be used to prevent a `Pod` from requesting a specific, disallowed scheduler.

### 1. Define the Tag in `types.go`
Apply the `+k8s:neq` tag to the `SchedulerName` field in a custom `PodSpec` struct, forbidding "forbidden-scheduler".

**File:** `pkg/apis/example/v1/types.go`
```go
type MyPodSpec struct {
    // ...
    // The "forbidden-scheduler" is not allowed for this Pod.
    // +k8s:neq="forbidden-scheduler"
    SchedulerName string `json:"schedulerName,omitempty"`
    // ...
}
```

### 2. Update Handwritten Validation
In the corresponding handwritten validation function for `MyPodSpec`, remove or mark as covered the explicit `!=` check for the forbidden scheduler name.

**File:** `pkg/apis/example/validation/validation.go`
```go
func ValidateMyPodSpec(spec *MyPodSpec, fldPath *field.Path) field.ErrorList {
    allErrs := field.ErrorList{}

    // ... other validations ...

    // Original handwritten validation for forbidding a scheduler:
    // if spec.SchedulerName == "forbidden-scheduler" {
    //     allErrs = append(allErrs, field.Invalid(fldPath.Child("schedulerName"), spec.SchedulerName, "forbidden-scheduler is not allowed"))
    // }

    // After adding +k8s:neq="forbidden-scheduler", the generated code handles this.
    // If still present for backward compatibility during migration, mark the error as covered:
    if spec.SchedulerName == "forbidden-scheduler" {
        allErrs = append(allErrs, field.Invalid(fldPath.Child("schedulerName"), spec.SchedulerName, "forbidden-scheduler is not allowed").MarkCoveredByDeclarative())
    }

    // ... other validations ...
    return allErrs
}
```
Using `+k8s:neq` makes the explicit exclusion of values a declarative part of the API definition.
