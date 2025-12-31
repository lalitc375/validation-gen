# Kind: NetworkPolicy

## 1. Internal API Type Definition
This is the internal representation of the `NetworkPolicy` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/networking/types.go`
*   **Type**: `type NetworkPolicy struct`

## 2. External API Type Definitions
These are the versioned representations of the `NetworkPolicy` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/networking/v1/types.go`
    *   **Type**: `type NetworkPolicy struct`
*   **Version: extensions/v1beta1**
    *   **Location**: `staging/src/k8s.io/api/extensions/v1beta1/types.go`
    *   **Type**: `type NetworkPolicy struct`

## 3. Validation Logic
Contains the business logic to validate `NetworkPolicy` objects during creation and updates.
*   **Location**: `pkg/apis/networking/validation/validation.go`
*   **Key Functions**:
    *   `ValidateNetworkPolicy`: Validates a NetworkPolicy.
    *   `ValidateNetworkPolicyUpdate`: Validates a NetworkPolicy during updates.
    *   `ValidateNetworkPolicySpec`: Validates the NetworkPolicySpec.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/networking/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_NetworkPolicy`: Sets defaults for the NetworkPolicy object.
    *   `SetDefaults_NetworkPolicyPort`: Sets defaults for the NetworkPolicyPort.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/networking/networkpolicy/strategy.go`
*   **Key Struct**: `networkPolicyStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
