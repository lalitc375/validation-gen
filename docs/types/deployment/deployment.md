# Kind: Deployment

## 1. Internal API Type Definition
This is the internal representation of the `Deployment` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/apps/types.go`
*   **Type**: `type Deployment struct`

## 2. External API Type Definitions
These are the versioned representations of the `Deployment` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/apps/v1/types.go`
    *   **Type**: `type Deployment struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/apps/v1beta1/types.go`
    *   **Type**: `type Deployment struct`
*   **Version: v1beta2**
    *   **Location**: `staging/src/k8s.io/api/apps/v1beta2/types.go`
    *   **Type**: `type Deployment struct`
*   **Version: extensions/v1beta1**
    *   **Location**: `staging/src/k8s.io/api/extensions/v1beta1/types.go`
    *   **Type**: `type Deployment struct`

## 3. Validation Logic
Contains the business logic to validate `Deployment` objects during creation and updates.
*   **Location**: `pkg/apis/apps/validation/validation.go`
*   **Key Functions**:
    *   `ValidateDeployment`: Validates a Deployment.
    *   `ValidateDeploymentUpdate`: Validates a Deployment during updates.
    *   `ValidateDeploymentStatusUpdate`: Validates status updates.
    *   `ValidateDeploymentSpec`: Validates the DeploymentSpec.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: `pkg/apis/apps/v1/defaults.go`
*   **Key Functions**:
    *   `SetDefaults_Deployment`: Sets defaults for the Deployment object.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/apps/deployment/strategy.go`
*   **Key Struct**: `deploymentStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
