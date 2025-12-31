# Kind: PodDisruptionBudget

## 1. Internal API Type Definition
This is the internal representation of the `PodDisruptionBudget` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/policy/types.go`
*   **Type**: `type PodDisruptionBudget struct`

## 2. External API Type Definitions
These are the versioned representations of the `PodDisruptionBudget` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/policy/v1/types.go`
    *   **Type**: `type PodDisruptionBudget struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/policy/v1beta1/types.go`
    *   **Type**: `type PodDisruptionBudget struct`

## 3. Validation Logic
Contains the business logic to validate `PodDisruptionBudget` objects during creation and updates.
*   **Location**: `pkg/apis/policy/validation/validation.go`
*   **Key Functions**:
    *   `ValidatePodDisruptionBudget`: Validates a PodDisruptionBudget.
    *   `ValidatePodDisruptionBudgetSpec`: Validates the PodDisruptionBudgetSpec.
    *   `ValidatePodDisruptionBudgetStatusUpdate`: Validates status updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: Not explicitly found in `pkg/apis/policy`.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/policy/poddisruptionbudget/strategy.go`
*   **Key Struct**: `podDisruptionBudgetStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
