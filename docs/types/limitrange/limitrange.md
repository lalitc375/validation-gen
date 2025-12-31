# Kind: LimitRange

## 1. Internal API Type Definition
This is the internal representation of the `LimitRange` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/core/types.go`
*   **Type**: `type LimitRange struct`

## 2. External API Type Definitions
These are the versioned representations of the `LimitRange` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/core/v1/types.go`
    *   **Type**: `type LimitRange struct`

## 3. Validation Logic
Contains the business logic to validate `LimitRange` objects during creation and updates.
*   **Location**: `pkg/apis/core/validation/validation.go`
*   **Key Functions**:
    *   `ValidateLimitRange`: Validates a LimitRange.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/core/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_LimitRangeItem`: Sets defaults for LimitRangeItem. (Note: Top-level `SetDefaults_LimitRange` not explicitly found in core defaults).

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/core/limitrange/strategy.go`
*   **Key Struct**: `limitrangeStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
