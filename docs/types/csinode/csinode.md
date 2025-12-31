# Kind: CSINode

## 1. Internal API Type Definition
This is the internal representation of the `CSINode` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/storage/types.go`
*   **Type**: `type CSINode struct`

## 2. External API Type Definitions
These are the versioned representations of the `CSINode` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/storage/v1/types.go`
    *   **Type**: `type CSINode struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/storage/v1beta1/types.go`
    *   **Type**: `type CSINode struct`

## 3. Validation Logic
Contains the business logic to validate `CSINode` objects during creation and updates.
*   **Location**: `pkg/apis/storage/validation/validation.go`
*   **Key Functions**:
    *   `ValidateCSINode`: Validates a CSINode.
    *   `ValidateCSINodeUpdate`: Validates a CSINode during updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: Not explicitly found in `pkg/apis/storage/v1/defaults.go`.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/storage/csinode/strategy.go`
*   **Key Struct**: `csiNodeStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
