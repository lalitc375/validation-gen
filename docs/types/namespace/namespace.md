# Kind: Namespace

## 1. Internal API Type Definition
This is the internal representation of the `Namespace` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/core/types.go`
*   **Type**: `type Namespace struct`

## 2. External API Type Definitions
These are the versioned representations of the `Namespace` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/core/v1/types.go`
    *   **Type**: `type Namespace struct`

## 3. Validation Logic
Contains the business logic to validate `Namespace` objects during creation and updates.
*   **Location**: `pkg/apis/core/validation/validation.go`
*   **Key Functions**:
    *   `ValidateNamespace`: Validates a Namespace.
    *   `ValidateNamespaceUpdate`: Validates a Namespace during updates.
    *   `ValidateNamespaceStatusUpdate`: Validates status updates.
    *   `ValidateNamespaceFinalizeUpdate`: Validates updates during finalization.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/core/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_Namespace`: Sets defaults for the Namespace object.
    *   `SetDefaults_NamespaceStatus`: Sets defaults for the Namespace status.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/core/namespace/strategy.go`
*   **Key Struct**: `namespaceStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
