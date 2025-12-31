# Kind: Node

## 1. Internal API Type Definition
This is the internal representation of the `Node` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/core/types.go`
*   **Type**: `type Node struct`

## 2. External API Type Definitions
These are the versioned representations of the `Node` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/core/v1/types.go`
    *   **Type**: `type Node struct`

## 3. Validation Logic
Contains the business logic to validate `Node` objects during creation and updates.
*   **Location**: `pkg/apis/core/validation/validation.go`
*   **Key Functions**:
    *   `ValidateNode`: Validates a Node.
    *   `ValidateNodeUpdate`: Validates a Node during updates.
    *   `ValidateNodeResources`: Validates Node resources.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/core/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_NodeStatus`: Sets defaults for the Node status. (Note: `SetDefaults_Node` was not explicitly found, but `SetDefaults_NodeStatus` handles part of it).

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/core/node/strategy.go`
*   **Key Struct**: `nodeStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
