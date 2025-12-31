# Kind: CSIDriver

## 1. Internal API Type Definition
This is the internal representation of the `CSIDriver` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/storage/types.go`
*   **Type**: `type CSIDriver struct`

## 2. External API Type Definitions
These are the versioned representations of the `CSIDriver` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/storage/v1/types.go`
    *   **Type**: `type CSIDriver struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/storage/v1beta1/types.go`
    *   **Type**: `type CSIDriver struct`

## 3. Validation Logic
Contains the business logic to validate `CSIDriver` objects during creation and updates.
*   **Location**: `pkg/apis/storage/validation/validation.go`
*   **Key Functions**:
    *   `ValidateCSIDriver`: Validates a CSIDriver.
    *   `ValidateCSIDriverUpdate`: Validates a CSIDriver during updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/storage/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_CSIDriver`: Sets defaults for the CSIDriver object.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/storage/csidriver/strategy.go`
*   **Key Struct**: `csiDriverStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
