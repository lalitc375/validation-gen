# Kind: HorizontalPodAutoscaler

## 1. Internal API Type Definition
This is the internal representation of the `HorizontalPodAutoscaler` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/autoscaling/types.go`
*   **Type**: `type HorizontalPodAutoscaler struct`

## 2. External API Type Definitions
These are the versioned representations of the `HorizontalPodAutoscaler` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/autoscaling/v1/types.go`
    *   **Type**: `type HorizontalPodAutoscaler struct`
*   **Version: v2**
    *   **Location**: `staging/src/k8s.io/api/autoscaling/v2/types.go`
    *   **Type**: `type HorizontalPodAutoscaler struct`
*   **Version: v2beta1**
    *   **Location**: `staging/src/k8s.io/api/autoscaling/v2beta1/types.go`
    *   **Type**: `type HorizontalPodAutoscaler struct`
*   **Version: v2beta2**
    *   **Location**: `staging/src/k8s.io/api/autoscaling/v2beta2/types.go`
    *   **Type**: `type HorizontalPodAutoscaler struct`

## 3. Validation Logic
Contains the business logic to validate `HorizontalPodAutoscaler` objects during creation and updates.
*   **Location**: `pkg/apis/autoscaling/validation/validation.go`
*   **Key Functions**:
    *   `ValidateHorizontalPodAutoscaler`: Validates a HorizontalPodAutoscaler.
    *   `ValidateHorizontalPodAutoscalerUpdate`: Validates a HorizontalPodAutoscaler during updates.
    *   `ValidateHorizontalPodAutoscalerStatusUpdate`: Validates status updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/autoscaling/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_HorizontalPodAutoscaler`: Sets defaults for the HorizontalPodAutoscaler object.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/autoscaling/horizontalpodautoscaler/strategy.go`
*   **Key Struct**: `autoscalerStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
