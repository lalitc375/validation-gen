# Kind: Binding

## 1. Internal API Type Definition
This is the internal representation of the `Binding` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/core/types.go`
*   **Type**: `type Binding struct`

## 2. External API Type Definitions
These are the versioned representations of the `Binding` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/core/v1/types.go`
    *   **Type**: `type Binding struct`

## 3. Validation Logic
Contains the business logic to validate `Binding` objects during creation and updates.
*   **Location**: `pkg/apis/core/validation/validation.go`
*   **Key Functions**:
    *   `ValidatePodBinding`: Validates a PodBinding.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/core/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_Binding`: (Not explicitly found in core defaults, likely handled implicitly or via generic defaulting).

## 5. Registry Strategy
The strategy for Binding is often associated with the Pod resource as it's typically used to bind a Pod to a Node.
*   **Location**: Likely within `pkg/registry/core/pod/` or handled specially as a subresource.
