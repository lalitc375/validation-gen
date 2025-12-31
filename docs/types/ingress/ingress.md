# Kind: Ingress

## 1. Internal API Type Definition
This is the internal representation of the `Ingress` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/networking/types.go`
*   **Type**: `type Ingress struct`

## 2. External API Type Definitions
These are the versioned representations of the `Ingress` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/networking/v1/types.go`
    *   **Type**: `type Ingress struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/networking/v1beta1/types.go`
    *   **Type**: `type Ingress struct`
*   **Version: extensions/v1beta1**
    *   **Location**: `staging/src/k8s.io/api/extensions/v1beta1/types.go`
    *   **Type**: `type Ingress struct`

## 3. Validation Logic
Contains the business logic to validate `Ingress` objects during creation and updates.
*   **Location**: `pkg/apis/networking/validation/validation.go`
*   **Key Functions**:
    *   `ValidateIngressCreate`: Validates an Ingress during creation.
    *   `ValidateIngressUpdate`: Validates an Ingress during updates.
    *   `ValidateIngressSpec`: Validates the IngressSpec.
    *   `ValidateIngressStatusUpdate`: Validates status updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/networking/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_IngressClass`: Sets defaults for the IngressClass. (Note: `SetDefaults_Ingress` not explicitly found).

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/networking/ingress/strategy.go`
*   **Key Struct**: `ingressStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
