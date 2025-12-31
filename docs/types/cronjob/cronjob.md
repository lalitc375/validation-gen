# Kind: CronJob

## 1. Internal API Type Definition
This is the internal representation of the `CronJob` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/batch/types.go`
*   **Type**: `type CronJob struct`

## 2. External API Type Definitions
These are the versioned representations of the `CronJob` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/batch/v1/types.go`
    *   **Type**: `type CronJob struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/batch/v1beta1/types.go`
    *   **Type**: `type CronJob struct`

## 3. Validation Logic
Contains the business logic to validate `CronJob` objects during creation and updates.
*   **Location**: `pkg/apis/batch/validation/validation.go`
*   **Key Functions**:
    *   `ValidateCronJobCreate`: Validates a CronJob during creation.
    *   `ValidateCronJobUpdate`: Validates a CronJob during updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/batch/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_CronJob`: Sets defaults for the CronJob object.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/batch/cronjob/strategy.go`
*   **Key Struct**: `cronJobStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
