# Kind: PersistentVolumeClaim

## 1. Internal API Type Definition
This is the internal representation of the `PersistentVolumeClaim` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/core/types.go`
*   **Type**: `type PersistentVolumeClaim struct`

## 2. External API Type Definitions
These are the versioned representations of the `PersistentVolumeClaim` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/core/v1/types.go`
    *   **Type**: `type PersistentVolumeClaim struct`

## 3. Validation Logic
Contains the business logic to validate `PersistentVolumeClaim` objects during creation and updates.
*   **Location**: `pkg/apis/core/validation/validation.go`
*   **Key Functions**:
    *   `ValidatePersistentVolumeClaim`: Validates a PersistentVolumeClaim.
    *   `ValidatePersistentVolumeClaimUpdate`: Validates a PersistentVolumeClaim during updates.
    *   `ValidatePersistentVolumeClaimSpec`: Validates the PersistentVolumeClaimSpec.
    *   `ValidatePersistentVolumeClaimStatusUpdate`: Validates status updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/core/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_PersistentVolumeClaim`: Sets defaults for the PersistentVolumeClaim object.
    *   `SetDefaults_PersistentVolumeClaimSpec`: Sets defaults for the PersistentVolumeClaimSpec.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/core/persistentvolumeclaim/strategy.go`
*   **Key Struct**: `persistentvolumeclaimStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
