# Kind: RoleBinding

## 1. Internal API Type Definition
This is the internal representation of the `RoleBinding` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/rbac/types.go`
*   **Type**: `type RoleBinding struct`

## 2. External API Type Definitions
These are the versioned representations of the `RoleBinding` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/rbac/v1/types.go`
    *   **Type**: `type RoleBinding struct`
*   **Version: v1alpha1**
    *   **Location**: `staging/src/k8s.io/api/rbac/v1alpha1/types.go`
    *   **Type**: `type RoleBinding struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/rbac/v1beta1/types.go`
    *   **Type**: `type RoleBinding struct`

## 3. Validation Logic
Contains the business logic to validate `RoleBinding` objects during creation and updates.
*   **Location**: `pkg/apis/rbac/validation/validation.go`
*   **Key Functions**:
    *   `ValidateRoleBinding`: Validates a RoleBinding.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/rbac/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_RoleBinding`: Sets defaults for the RoleBinding object.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/rbac/rolebinding/strategy.go`
*   **Key Struct**: `strategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
