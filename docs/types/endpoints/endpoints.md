# Kind: Endpoints

## 1. Internal API Type Definition
This is the internal representation of the `Endpoints` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/core/types.go`
*   **Type**: `type Endpoints struct`

## 2. External API Type Definitions
These are the versioned representations of the `Endpoints` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/core/v1/types.go`
    *   **Type**: `type Endpoints struct`

## 3. Validation Logic
Contains the business logic to validate `Endpoints` objects during creation and updates.
*   **Location**: `pkg/apis/core/validation/validation.go`
*   **Key Functions**:
    *   `ValidateEndpoints`: General validation for Endpoints.
    *   `ValidateEndpointsCreate`: Validates Endpoints during creation.
    *   `ValidateEndpointsUpdate`: Validates Endpoints during updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/core/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_Endpoints`: Sets defaults for the Endpoints object.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/core/endpoint/strategy.go`
*   **Key Struct**: `endpointsStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
