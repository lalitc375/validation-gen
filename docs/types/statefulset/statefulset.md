# Kind: StatefulSet

## 1. Internal API Type Definition
This is the internal representation of the `StatefulSet` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/apps/types.go`
*   **Type**: `type StatefulSet struct`

## 2. External API Type Definitions
These are the versioned representations of the `StatefulSet` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/apps/v1/types.go`
    *   **Type**: `type StatefulSet struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/apps/v1beta1/types.go`
    *   **Type**: `type StatefulSet struct`
*   **Version: v1beta2**
    *   **Location**: `staging/src/k8s.io/api/apps/v1beta2/types.go`
    *   **Type**: `type StatefulSet struct`

## 3. Validation Logic
Contains the business logic to validate `StatefulSet` objects during creation and updates.
*   **Location**: `pkg/apis/apps/validation/validation.go`
*   **Key Functions**:
    *   `ValidateStatefulSet`: Validates a StatefulSet.
    *   `ValidateStatefulSetUpdate`: Validates a StatefulSet during updates.
    *   `ValidateStatefulSetStatusUpdate`: Validates status updates.
    *   `ValidateStatefulSetSpec`: Validates the StatefulSetSpec.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/apps/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_StatefulSet`: Sets defaults for the StatefulSet object.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/apps/statefulset/strategy.go`
*   **Key Struct**: `statefulSetStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
