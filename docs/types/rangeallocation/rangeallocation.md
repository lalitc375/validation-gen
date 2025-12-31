# Kind: RangeAllocation

## 1. Internal API Type Definition
This is the internal representation of the `RangeAllocation` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/core/types.go`
*   **Type**: `type RangeAllocation struct`

## 2. External API Type Definitions
These are the versioned representations of the `RangeAllocation` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/core/v1/types.go`
    *   **Type**: `type RangeAllocation struct`

## 3. Validation Logic
Contains the business logic to validate `RangeAllocation` objects.
*   **Location**: Not explicitly found in `pkg/apis/core/validation/validation.go`. This is primarily an internal helper object for allocators.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: Not explicitly found in `pkg/apis/core/v1/defaults.go`.

## 5. Registry Strategy
RangeAllocation is typically used internally by controllers (like the Service controller) to persist allocation state.
*   **Location**: `pkg/registry/core/rangeallocation/`
