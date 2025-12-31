# Kind: RuntimeClass

## 1. Internal API Type Definition
This is the internal representation of the `RuntimeClass` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/node/types.go`
*   **Type**: `type RuntimeClass struct`

## 2. External API Type Definitions
These are the versioned representations of the `RuntimeClass` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/node/v1/types.go`
    *   **Type**: `type RuntimeClass struct`
*   **Version: v1alpha1**
    *   **Location**: `staging/src/k8s.io/api/node/v1alpha1/types.go`
    *   **Type**: `type RuntimeClass struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/node/v1beta1/types.go`
    *   **Type**: `type RuntimeClass struct`

## 3. Validation Logic
Contains the business logic to validate `RuntimeClass` objects during creation and updates.
*   **Location**: `pkg/apis/node/validation/validation.go`
*   **Key Functions**:
    *   `ValidateRuntimeClass`: Validates a RuntimeClass.
    *   `ValidateRuntimeClassUpdate`: Validates a RuntimeClass during updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: Not explicitly found in `pkg/apis/node`.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/node/runtimeclass/strategy.go`
*   **Key Struct**: Not explicitly named `Strategy`, but logic is present.
