# Kind: Event

## 1. Internal API Type Definition
This is the internal representation of the `Event` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/core/types.go`
*   **Type**: `type Event struct`

## 2. External API Type Definitions
These are the versioned representations of the `Event` object used by clients and persisted in etcd.
*   **Version: v1** (Stable, Core)
    *   **Location**: `staging/src/k8s.io/api/core/v1/types.go`
    *   **Type**: `type Event struct`
*   **Version: v1** (events.k8s.io)
    *   **Location**: `staging/src/k8s.io/api/events/v1/types.go`
    *   **Type**: `type Event struct`
*   **Version: v1beta1** (events.k8s.io)
    *   **Location**: `staging/src/k8s.io/api/events/v1beta1/types.go`
    *   **Type**: `type Event struct`

## 3. Validation Logic
Contains the business logic to validate `Event` objects during creation and updates.
*   **Location**: `pkg/apis/core/validation/events.go`
*   **Key Functions**:
    *   `ValidateEventCreate`: Validates an Event during creation.
    *   `ValidateEventUpdate`: Validates an Event during updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/core/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_Event`: (Not explicitly found in core defaults, likely handled implicitly or via generic defaulting).

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/core/event/strategy.go`
*   **Key Struct**: `eventStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
