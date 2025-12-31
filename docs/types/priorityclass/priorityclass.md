# Kind: PriorityClass

## 1. Internal API Type Definition
This is the internal representation of the `PriorityClass` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/scheduling/types.go`
*   **Type**: `type PriorityClass struct`

## 2. External API Type Definitions
These are the versioned representations of the `PriorityClass` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/scheduling/v1/types.go`
    *   **Type**: `type PriorityClass struct`
*   **Version: v1alpha1**
    *   **Location**: `staging/src/k8s.io/api/scheduling/v1alpha1/types.go`
    *   **Type**: `type PriorityClass struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/scheduling/v1beta1/types.go`
    *   **Type**: `type PriorityClass struct`

## 3. Validation Logic
Contains the business logic to validate `PriorityClass` objects during creation and updates.
*   **Location**: `pkg/apis/scheduling/validation/validation.go`
*   **Key Functions**:
    *   `ValidatePriorityClass`: Validates a PriorityClass.
    *   `ValidatePriorityClassUpdate`: Validates a PriorityClass during updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/scheduling/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_PriorityClass`: Sets defaults for the PriorityClass object.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/scheduling/priorityclass/strategy.go`
*   **Key Struct**: `priorityClassStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
