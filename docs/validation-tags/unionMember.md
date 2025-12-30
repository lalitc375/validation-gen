# +k8s:unionMember

## Description
Marks a field as a member of a union. This tag is used in conjunction with `+k8s:unionDiscriminator` to define a discriminated union, where only one member field can be set at a time based on the value of a discriminator field.

## Scope
`Field`

## Supported Go Types
Any Go type. The tag indicates that this field is one of the possible choices in a union.

## Arguments
*   `union=<name>` (optional): Specifies the name of the union if a struct contains multiple unions.
*   `memberName=<value>` (optional, defaults to field name): Specifies the value of the discriminator field that activates this union member. If omitted, the JSON field name is used.

## Stability
**Stable**

## Usage

### Field
```go
type MyUnionStruct struct {
    // This field acts as the discriminator.
    // +k8s:unionDiscriminator
    D MyDiscriminatorType `json:"d"`

    // This field is a member of the union, activated when D is "M1".
    // +k8s:unionMember
    // +k8s:optional
    M1 *MemberType1 `json:"m1"`

    // This field is another member, activated when D is "M2".
    // +k8s:unionMember
    // +k8s:optional
    M2 *MemberType2 `json:"m2"`
}

type MyDiscriminatorType string

const (
    DiscriminatorM1 MyDiscriminatorType = "M1"
    DiscriminatorM2 MyDiscriminatorType = "M2"
)

type MemberType1 struct { /* ... */ }
type MemberType2 struct { /* ... */ }
```
When `D` is `M1`, only `M1` is allowed to be set (non-nil). If `D` is `M2`, only `M2` is allowed.

## Migrating from Handwritten Validation

The `+k8s:unionMember` tag, used with `+k8s:unionDiscriminator`, streamlines the validation of discriminated unions. This declarative approach replaces handwritten logic that would traditionally involve:
1.  Checking the value of a discriminator field.
2.  Based on the discriminator, verifying that only the corresponding union member field(s) are set and others are unset (mutual exclusivity).
3.  Generating appropriate `field.Required` or `field.Forbidden` errors.

By defining these relationships with tags, the Kubernetes API machinery handles the validation automatically, leading to more concise and understandable API definitions.

## Detailed Example: Discriminated Union for Configuration Type

This example demonstrates how `+k8s:unionMember` fields are controlled by a `+k8s:unionDiscriminator` to ensure that only one configuration type is active at a time.

### 1. Define the Tags in `types.go`
Define a struct with a discriminator field (`ConfigType`) and multiple union member fields (`HTTPConfig`, `FileConfig`), each marked appropriately.

**File:** `pkg/apis/example/v1/types.go`
```go
type ApplicationConfig struct {
    // This field acts as the discriminator.
    // +k8s:unionDiscriminator
    ConfigType string `json:"configType"`

    // Only one of these can be set based on ConfigType.
    // +k8s:unionMember
    // +k8s:optional
    HTTPConfig *HTTPConfiguration `json:"httpConfig,omitempty"`

    // +k8s:unionMember
    // +k8s:optional
    FileConfig *FileConfiguration `json:"fileConfig,omitempty"`
}

type HTTPConfiguration struct {
    Port int `json:"port"`
}

type FileConfiguration struct {
    Path string `json:"path"`
}
```

### 2. Define Discriminator Values
Constants are typically used to define the possible values for the discriminator.

**File:** `pkg/apis/example/v1/types.go`
```go
const (
    ConfigTypeHTTP string = "HTTP"
    ConfigTypeFile string = "File"
)
```

### 3. Update Handwritten Validation
With the `+k8s:unionDiscriminator` and `+k8s:unionMember` tags, the structural validation for mutual exclusivity is handled automatically by the generated code. Any previous handwritten checks for this can be removed or marked as covered.

**File:** `pkg/apis/example/validation/validation.go`
```go
func ValidateApplicationConfig(appConfig *ApplicationConfig, fldPath *field.Path) field.ErrorList {
    allErrs := field.ErrorList{}

    // Original handwritten validation logic (example):
    // if appConfig.ConfigType == ConfigTypeHTTP {
    //     if appConfig.HTTPConfig == nil {
    //         allErrs = append(allErrs, field.Required(fldPath.Child("httpConfig"), "must be specified for HTTP config type"))
    //     }
    //     if appConfig.FileConfig != nil {
    //         allErrs = append(allErrs, field.Forbidden(fldPath.Child("fileConfig"), "must not be specified for HTTP config type"))
    //     }
    // } else if appConfig.ConfigType == ConfigTypeFile {
    //     if appConfig.FileConfig == nil {
    //         allErrs = append(allErrs, field.Required(fldPath.Child("fileConfig"), "must be specified for File config type"))
    //     }
    //     if appConfig.HTTPConfig != nil {
    //         allErrs = append(allErrs, field.Forbidden(fldPath.Child("httpConfig"), "must not be specified for File config type"))
    //     }
    // } else {
    //     allErrs = append(allErrs, field.Invalid(fldPath.Child("configType"), appConfig.ConfigType, "unknown config type"))
    // }

    // After adding +k8s:unionDiscriminator and +k8s:unionMember, this logic is handled automatically.
    // Any remaining handwritten validation would typically be for the content of HTTPConfiguration or FileConfiguration.
    if appConfig.HTTPConfig != nil {
        // Validate specific fields within HTTPConfig
        if appConfig.HTTPConfig.Port < 1 || appConfig.HTTPConfig.Port > 65535 {
            allErrs = append(allErrs, field.Invalid(fldPath.Child("httpConfig", "port"), appConfig.HTTPConfig.Port, "must be a valid port number").MarkCoveredByDeclarative())
        }
    }
    if appConfig.FileConfig != nil {
        // Validate specific fields within FileConfig
        if appConfig.FileConfig.Path == "" {
            allErrs = append(allErrs, field.Required(fldPath.Child("fileConfig", "path"), "path must be specified").MarkCoveredByDeclarative())
        }
    }

    return allErrs
}
```
The `+k8s:unionMember` tag, together with `+k8s:unionDiscriminator`, provides a robust and declarative way to define and validate one-of relationships in API objects.
