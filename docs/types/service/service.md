# Kind: Service

## 1. Internal API Type Definition
This is the internal representation of the `Service` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/core/types.go`
*   **Type**: `type Service struct`

## 2. External API Type Definitions
These are the versioned representations of the `Service` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/core/v1/types.go`
    *   **Type**: `type Service struct`

## 3. Validation Logic
Contains the business logic to validate `Service` objects during creation and updates.
*   **Location**: `pkg/apis/core/validation/validation.go`
*   **Key Functions**:
    *   `ValidateServiceCreate`: Validates a Service during creation.
    *   `ValidateServiceUpdate`: Validates a Service during updates.
    *   `ValidateServiceStatusUpdate`: Validates status updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/core/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_Service`: Sets defaults for the Service object.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/core/service/strategy.go`
*   **Key Struct**: `svcStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
