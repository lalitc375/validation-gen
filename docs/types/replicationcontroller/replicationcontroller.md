# Kind: ReplicationController

## 1. Internal API Type Definition
This is the internal representation of the `ReplicationController` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/core/types.go`
*   **Type**: `type ReplicationController struct`

## 2. External API Type Definitions
These are the versioned representations of the `ReplicationController` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/core/v1/types.go`
    *   **Type**: `type ReplicationController struct`

## 3. Validation Logic
Contains the business logic to validate `ReplicationController` objects during creation and updates.
*   **Location**: `pkg/apis/core/validation/validation.go`
*   **Key Functions**:
    *   `ValidateReplicationController`: Validates a ReplicationController.
    *   `ValidateReplicationControllerUpdate`: Validates a ReplicationController during updates.
    *   `ValidateReplicationControllerStatusUpdate`: Validates status updates.
    *   `ValidateReplicationControllerSpec`: Validates the ReplicationControllerSpec.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/core/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_ReplicationController`: Sets defaults for the ReplicationController object.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/core/replicationcontroller/strategy.go`
*   **Key Struct**: `rcStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
