# Kind: Eviction

## 1. Internal API Type Definition
This is the internal representation of the `Eviction` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/policy/types.go`
*   **Type**: `type Eviction struct`

## 2. External API Type Definitions
These are the versioned representations of the `Eviction` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/policy/v1/types.go`
    *   **Type**: `type Eviction struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/policy/v1beta1/types.go`
    *   **Type**: `type Eviction struct`

## 3. Validation Logic
Contains the business logic to validate `Eviction` objects during creation.
*   **Location**: Not explicitly found in `pkg/apis/policy/validation/validation.go` in a `ValidateEviction` function. Validation is likely handled within the eviction subresource handler.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: Not explicitly found in `pkg/apis/policy`.

## 5. Registry Strategy
The strategy for Eviction is part of the Pod resource's storage implementation.
*   **Location**: `pkg/registry/core/pod/storage/eviction.go`
