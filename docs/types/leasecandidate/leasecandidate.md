# Kind: LeaseCandidate

## 1. Internal API Type Definition
This is the internal representation of the `LeaseCandidate` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/coordination/types.go`
*   **Type**: `type LeaseCandidate struct`

## 2. External API Type Definitions
These are the versioned representations of the `LeaseCandidate` object used by clients and persisted in etcd.
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/coordination/v1beta1/types.go`
    *   **Type**: `type LeaseCandidate struct`
*   **Version: v1alpha2**
    *   **Location**: `staging/src/k8s.io/api/coordination/v1alpha2/types.go`
    *   **Type**: `type LeaseCandidate struct`

## 3. Validation Logic
Contains the business logic to validate `LeaseCandidate` objects during creation and updates.
*   **Location**: `pkg/apis/coordination/validation/validation.go`
*   **Key Functions**:
    *   `ValidateLeaseCandidate`: Validates a LeaseCandidate.
    *   `ValidateLeaseCandidateUpdate`: Validates a LeaseCandidate during updates.
    *   `ValidateLeaseCandidateSpec`: Validates the LeaseCandidateSpec.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: Not explicitly found in `pkg/apis/coordination`.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/coordination/leasecandidate/strategy.go`
*   **Key Struct**: `LeaseCandidateStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
