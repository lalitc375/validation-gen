# Kubernetes Union Validations

This document lists the Kubernetes Kinds that implement union validation logic and demonstrates how to apply the `+k8s:unionDiscriminator` and `+k8s:unionMember` tags to them.

## 1. Authorization

**File:** `pkg/apis/authorization/types.go`

### `SubjectAccessReviewSpec` (and `SelfSubjectAccessReviewSpec`)

This is an undiscriminated union.

```go
type SubjectAccessReviewSpec struct {
    // +k8s:unionMember
    // +k8s:optional
    ResourceAttributes *ResourceAttributes
    // +k8s:unionMember
    // +k8s:optional
    NonResourceAttributes *NonResourceAttributes

    // ... other fields
}
```

### `LabelSelectorAttributes` (and `FieldSelectorAttributes`)

This is an undiscriminated union.

```go
type LabelSelectorAttributes struct {
    // +k8s:unionMember
    // +k8s:optional
    RawSelector string

    // +k8s:unionMember
    // +k8s:optional
    Requirements []metav1.LabelSelectorRequirement
}
```

## 2. Resource

**File:** `pkg/apis/resource/types.go`

### `ResourceSliceSpec`

This spec contains two separate undiscriminated unions.

**Union 1: Node Selection**

```go
type ResourceSliceSpec struct {
    // +k8s:unionMember(union="NodeSelection")
    // +k8s:optional
    NodeName *string

    // +k8s:unionMember(union="NodeSelection")
    // +k8s:optional
    NodeSelector *core.NodeSelector

    // +k8s:unionMember(union="NodeSelection")
    // +k8s:optional
    AllNodes *bool

    // +k8s:unionMember(union="NodeSelection")
    // +k8s:optional
    PerDeviceNodeSelection *bool

    // ... other fields
}
```

**Union 2: Content Type**

```go
type ResourceSliceSpec struct {
    // ... other fields

    // +k8s:unionMember(union="ResourceSliceType")
    // +k8s:optional
    Devices []Device

    // +k8s:unionMember(union="ResourceSliceType")
    // +k8s:optional
    SharedCounters []CounterSet
}
```

### `Device`

A partial undiscriminated union for node selection (only if `PerDeviceNodeSelection` is true).

```go
type Device struct {
    // ... other fields

    // +k8s:unionMember(union="DeviceNodeSelection")
    // +k8s:optional
    NodeName *string

    // +k8s:unionMember(union="DeviceNodeSelection")
    // +k8s:optional
    NodeSelector *core.NodeSelector

    // +k8s:unionMember(union="DeviceNodeSelection")
    // +k8s:optional
    AllNodes *bool

    // ... other fields
}
```

### `DeviceRequest`

This is an undiscriminated union.

```go
type DeviceRequest struct {
    // ... other fields

    // +k8s:unionMember
    // +k8s:optional
    Exactly *ExactDeviceRequest

    // +k8s:unionMember
    // +k8s:optional
    FirstAvailable []DeviceSubRequest
}
```

### `DeviceConstraint`

This is an undiscriminated union.

```go
type DeviceConstraint struct {
    // ... other fields

    // +k8s:unionMember
    // +k8s:optional
    MatchAttribute *FullyQualifiedName

    // +k8s:unionMember
    // +k8s:optional
    DistinctAttribute *FullyQualifiedName
}
```

### `DeviceAttribute`

This is an undiscriminated union.

```go
type DeviceAttribute struct {
    // +k8s:unionMember
    // +k8s:optional
    IntValue *int64

    // +k8s:unionMember
    // +k8s:optional
    BoolValue *bool

    // +k8s:unionMember
    // +k8s:optional
    StringValue *string

    // +k8s:unionMember
    // +k8s:optional
    VersionValue *string
}
```

### `CapacityRequestPolicy`

This is an undiscriminated union.

```go
type CapacityRequestPolicy struct {
    // ... other fields

    // +k8s:unionMember
    // +k8s:optional
    ValidValues []resource.Quantity

    // +k8s:unionMember
    // +k8s:optional
    ValidRange *CapacityRequestPolicyRange
}
```

## 3. AdmissionRegistration

**File:** `pkg/apis/admissionregistration/types.go`

### `WebhookClientConfig`

This is an undiscriminated union.

```go
type WebhookClientConfig struct {
    // +k8s:unionMember
    // +k8s:optional
    URL *string

    // +k8s:unionMember
    // +k8s:optional
    Service *ServiceReference

    // ... other fields
}
```

### `ParamRef`

This is an undiscriminated union.

```go
type ParamRef struct {
    // +k8s:unionMember
    // +k8s:optional
    Name string

    // ... other fields

    // +k8s:unionMember
    // +k8s:optional
    Selector *metav1.LabelSelector

    // ... other fields
}
```

### `Mutation`

This is a discriminated union.

```go
type Mutation struct {
    // +k8s:unionDiscriminator
    PatchType PatchType

    // +k8s:unionMember(memberName="ApplyConfiguration")
    // +k8s:optional
    ApplyConfiguration *ApplyConfiguration

    // +k8s:unionMember(memberName="JSONPatch")
    // +k8s:optional
    JSONPatch *JSONPatch
}
```

### `ClusterTrustBundleProjection`

This is an undiscriminated union.

```go
type ClusterTrustBundleProjection struct {
    // +k8s:unionMember
    // +k8s:optional
    Name *string

    // +k8s:unionMember
    // +k8s:optional
    SignerName *string

    // +k8s:unionMember
    // +k8s:optional
    LabelSelector *metav1.LabelSelector

    // ... other fields
}
```

## 4. Scheduling

**File:** `pkg/apis/scheduling/types.go`

### `PodGroupPolicy`

This is an undiscriminated union.

```go
type PodGroupPolicy struct {
    // +k8s:unionMember
    // +k8s:optional
    Basic *BasicSchedulingPolicy

    // +k8s:unionMember
    // +k8s:optional
    Gang *GangSchedulingPolicy
}
```

## 5. FlowControl

**File:** `pkg/apis/flowcontrol/types.go`

### `Subject`

This is a discriminated union.

```go
type Subject struct {
    // +k8s:unionDiscriminator
    Kind SubjectKind

    // +k8s:unionMember(memberName="User")
    // +k8s:optional
    User *UserSubject

    // +k8s:unionMember(memberName="Group")
    // +k8s:optional
    Group *GroupSubject

    // +k8s:unionMember(memberName="ServiceAccount")
    // +k8s:optional
    ServiceAccount *ServiceAccountSubject
}
```

### `PriorityLevelConfigurationSpec`

This is a discriminated union.

```go
type PriorityLevelConfigurationSpec struct {
    // +k8s:unionDiscriminator
    Type PriorityLevelEnablement

    // +k8s:unionMember(memberName="Limited")
    // +k8s:optional
    Limited *LimitedPriorityLevelConfiguration

    // +k8s:unionMember(memberName="Exempt")
    // +k8s:optional
    Exempt *ExemptPriorityLevelConfiguration
}
```

### `LimitResponse`

This is a discriminated union.

```go
type LimitResponse struct {
    // +k8s:unionDiscriminator
    Type LimitResponseType

    // +k8s:unionMember(memberName="Queue")
    // +k8s:optional
    Queuing *QueuingConfiguration
}
```

## 6. Storage

**File:** `pkg/apis/storage/types.go`

### `VolumeAttachmentSource`

This is an undiscriminated union.

```go
type VolumeAttachmentSource struct {
    // +k8s:unionMember
    // +k8s:optional
    PersistentVolumeName *string

    // +k8s:unionMember
    // +k8s:optional
    InlineVolumeSpec *api.PersistentVolumeSpec
}
```

## 7. Networking

**File:** `pkg/apis/networking/types.go`

### `IngressBackend`

This is an undiscriminated union.

```go
type IngressBackend struct {
    // +k8s:unionMember
    // +k8s:optional
    Service *IngressServiceBackend

    // +k8s:unionMember
    // +k8s:optional
    Resource *api.TypedLocalObjectReference
}
```

### `ServiceBackendPort`

This is an undiscriminated union.

```go
type ServiceBackendPort struct {
    // +k8s:unionMember
    // +k8s:optional
    Name string

    // +k8s:unionMember
    // +k8s:optional
    Number int32
}
```

### `NetworkPolicyPeer`

This is an undiscriminated union between `IPBlock` and the combination of `PodSelector`/`NamespaceSelector`.

```go
type NetworkPolicyPeer struct {
    // ... PodSelector and NamespaceSelector are NOT part of a simple union
    // together, but they ARE mutually exclusive with IPBlock.

    // +k8s:unionMember(union="PeerType")
    // +k8s:optional
    PodSelector *metav1.LabelSelector

    // +k8s:unionMember(union="PeerType")
    // +k8s:optional
    NamespaceSelector *metav1.LabelSelector

    // +k8s:unionMember(union="PeerType")
    // +k8s:optional
    IPBlock *IPBlock
}
```

## 8. Core

**File:** `pkg/apis/core/types.go`

### `VolumeSource` (and `PersistentVolumeSource`)

Large undiscriminated unions.

```go
type VolumeSource struct {
    // +k8s:unionMember
    // +k8s:optional
    HostPath *HostPathVolumeSource
    // +k8s:unionMember
    // +k8s:optional
    EmptyDir *EmptyDirVolumeSource
    // ... all other VolumeSource fields should have +k8s:unionMember
}
```

### `VolumeProjection`

This is an undiscriminated union.

```go
type VolumeProjection struct {
    // +k8s:unionMember
    // +k8s:optional
    Secret *SecretProjection
    // +k8s:unionMember
    // +k8s:optional
    DownwardAPI *DownwardAPIProjection
    // +k8s:unionMember
    // +k8s:optional
    ConfigMap *ConfigMapProjection
    // +k8s:unionMember
    // +k8s:optional
    ServiceAccountToken *ServiceAccountTokenProjection
    // +k8s:unionMember
    // +k8s:optional
    ClusterTrustBundle *ClusterTrustBundleProjection
    // +k8s:unionMember
    // +k8s:optional
    PodCertificate *PodCertificateProjection
}
```

### `DownwardAPIVolumeFile`

This is an undiscriminated union.

```go
type DownwardAPIVolumeFile struct {
    // ... fields

    // +k8s:unionMember
    // +k8s:optional
    FieldRef *ObjectFieldSelector

    // +k8s:unionMember
    // +k8s:optional
    ResourceFieldRef *ResourceFieldSelector

    // ... fields
}
```

### `ProbeHandler` (and `LifecycleHandler`)

This is an undiscriminated union.

```go
type ProbeHandler struct {
    // +k8s:unionMember
    // +k8s:optional
    Exec *ExecAction
    // +k8s:unionMember
    // +k8s:optional
    HTTPGet *HTTPGetAction
    // +k8s:unionMember
    // +k8s:optional
    TCPSocket *TCPSocketAction
    // +k8s:unionMember
    // +k8s:optional
    GRPC *GRPCAction
}
```

### `EnvVar`

This is an undiscriminated union.

```go
type EnvVar struct {
    // ... Name

    // +k8s:unionMember
    // +k8s:optional
    Value string

    // +k8s:unionMember
    // +k8s:optional
    ValueFrom *EnvVarSource
}
```

### `EnvVarSource`

This is an undiscriminated union.

```go
type EnvVarSource struct {
    // +k8s:unionMember
    // +k8s:optional
    FieldRef *ObjectFieldSelector
    // +k8s:unionMember
    // +k8s:optional
    ResourceFieldRef *ResourceFieldSelector
    // +k8s:unionMember
    // +k8s:optional
    ConfigMapKeyRef *ConfigMapKeySelector
    // +k8s:unionMember
    // +k8s:optional
    SecretKeyRef *SecretKeySelector
    // +k8s:unionMember
    // +k8s:optional
    FileKeyRef *FileKeySelector
}
```

### `PodResourceClaim`

This is an undiscriminated union.

```go
type PodResourceClaim struct {
    // ... Name

    // +k8s:unionMember
    // +k8s:optional
    ResourceClaimName *string

    // +k8s:unionMember
    // +k8s:optional
    ResourceClaimTemplateName *string
}
```

### `ContainerState`

This is an undiscriminated union.

```go
type ContainerState struct {
    // +k8s:unionMember
    // +k8s:optional
    Waiting *ContainerStateWaiting
    // +k8s:unionMember
    // +k8s:optional
    Running *ContainerStateRunning
    // +k8s:unionMember
    // +k8s:optional
    Terminated *ContainerStateTerminated
}
```

### `SeccompProfile`

This is a discriminated union.

```go
type SeccompProfile struct {
    // +k8s:unionDiscriminator
    Type SeccompProfileType

    // +k8s:unionMember(memberName="Localhost")
    // +k8s:optional
    LocalhostProfile *string
}
```