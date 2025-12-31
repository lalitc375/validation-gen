# Kind: IPAddress

## 1. Internal API Type Definition
This is the internal representation of the `IPAddress` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/networking/types.go`
*   **Type**: `type IPAddress struct`

## 2. External API Type Definitions
These are the versioned representations of the `IPAddress` object used by clients and persisted in etcd.
*   **Version: v1**
    *   **Location**: `staging/src/k8s.io/api/networking/v1/types.go`
    *   **Type**: `type IPAddress struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/networking/v1beta1/types.go`
    *   **Type**: `type IPAddress struct`

## 3. Validation Logic
Contains the business logic to validate `IPAddress` objects during creation and updates.
*   **Location**: `pkg/apis/networking/validation/validation.go`
*   **Key Functions**:
    *   `ValidateIPAddress`: Validates an IPAddress.
    *   `ValidateIPAddressUpdate`: Validates an IPAddress during updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: Not explicitly found in `pkg/apis/networking/v1/defaults.go` or `v1beta1/defaults.go`.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/networking/ipaddress/strategy.go`
*   **Key Struct**: `ipAddressStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
