# Kind: IngressClass

## 1. Internal API Type Definition
This is the internal representation of the `IngressClass` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/networking/types.go`
*   **Type**: `type IngressClass struct`

## 2. External API Type Definitions
These are the versioned representations of the `IngressClass` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/networking/v1/types.go`
    *   **Type**: `type IngressClass struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/networking/v1beta1/types.go`
    *   **Type**: `type IngressClass struct`

## 3. Validation Logic
Contains the business logic to validate `IngressClass` objects during creation and updates.
*   **Location**: `pkg/apis/networking/validation/validation.go`
*   **Key Functions**:
    *   `ValidateIngressClass`: Validates an IngressClass.
    *   `ValidateIngressClassUpdate`: Validates an IngressClass during updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/networking/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_IngressClass`: Sets defaults for the IngressClass object.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/networking/ingressclass/strategy.go`
*   **Key Struct**: `ingressClassStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
