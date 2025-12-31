# Kind: ControllerRevision

## 1. Internal API Type Definition
This is the internal representation of the `ControllerRevision` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/apps/types.go`
*   **Type**: `type ControllerRevision struct`

## 2. External API Type Definitions
These are the versioned representations of the `ControllerRevision` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/apps/v1/types.go`
    *   **Type**: `type ControllerRevision struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/apps/v1beta1/types.go`
    *   **Type**: `type ControllerRevision struct`
*   **Version: v1beta2**
    *   **Location**: `staging/src/k8s.io/api/apps/v1beta2/types.go`
    *   **Type**: `type ControllerRevision struct`

## 3. Validation Logic
Contains the business logic to validate `ControllerRevision` objects during creation and updates.
*   **Location**: `pkg/apis/apps/validation/validation.go`
*   **Key Functions**:
    *   `ValidateControllerRevisionCreate`: Validates a ControllerRevision during creation.
    *   `ValidateControllerRevisionUpdate`: Validates a ControllerRevision during updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: Not explicitly found in `pkg/apis/apps/v1/defaults.go`.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/apps/controllerrevision/strategy.go`
*   **Key Struct**: `strategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
