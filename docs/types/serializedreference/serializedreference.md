# Kind: SerializedReference

## 1. Internal API Type Definition
This is the internal representation of the `SerializedReference` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/core/types.go`
*   **Type**: `type SerializedReference struct`

## 2. External API Type Definitions
These are the versioned representations of the `SerializedReference` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/core/v1/types.go`
    *   **Type**: `type SerializedReference struct`

## 3. Validation Logic
Contains the business logic to validate `SerializedReference` objects.
*   **Location**: Not explicitly found in `pkg/apis/core/validation/validation.go`. This type is largely deprecated in favor of `ObjectReference`.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: Not explicitly found in `pkg/apis/core/v1/defaults.go`.

## 5. Registry Strategy
SerializedReference is typically not exposed as a top-level REST resource.
