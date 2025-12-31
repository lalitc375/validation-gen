# Kind: Pod

## 1. Internal API Type Definition
This is the internal representation of the `Pod` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/core/types.go`
*   **Type**: `type Pod struct`

## 2. External API Type Definitions
These are the versioned representations of the `Pod` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/core/v1/types.go`
    *   **Type**: `type Pod struct`

## 3. Validation Logic
Contains the business logic to validate `Pod` objects during creation and updates.
*   **Location**: `pkg/apis/core/validation/validation.go`
*   **Key Functions**:
    *   `ValidatePodCreate`: Validates a Pod during creation.
    *   `ValidatePodUpdate`: Validates a Pod during updates.
    *   `ValidatePodSpec`: Validates the `PodSpec`.
    *   `ValidatePodStatusUpdate`: Validates status updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/core/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_Pod`: Sets defaults for the top-level Pod object.
    *   `SetDefaults_PodSpec`: Sets defaults for the `PodSpec`.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/core/pod/strategy.go`
*   **Key Struct**: `podStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
