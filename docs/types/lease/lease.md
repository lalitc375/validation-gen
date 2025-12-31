# Kind: Lease

## 1. Internal API Type Definition
This is the internal representation of the `Lease` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/coordination/types.go`
*   **Type**: `type Lease struct`

## 2. External API Type Definitions
These are the versioned representations of the `Lease` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/coordination/v1/types.go`
    *   **Type**: `type Lease struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/coordination/v1beta1/types.go`
    *   **Type**: `type Lease struct`

## 3. Validation Logic
Contains the business logic to validate `Lease` objects during creation and updates.
*   **Location**: `pkg/apis/coordination/validation/validation.go`
*   **Key Functions**:
    *   `ValidateLease`: Validates a Lease.
    *   `ValidateLeaseUpdate`: Validates a Lease during updates.
    *   `ValidateLeaseSpec`: Validates the LeaseSpec.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: Not explicitly found in `pkg/apis/coordination`.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/coordination/lease/strategy.go`
*   **Key Struct**: `leaseStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
