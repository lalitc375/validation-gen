# Kind: CertificateSigningRequest

## 1. Internal API Type Definition
This is the internal representation of the `CertificateSigningRequest` object used within the Kubernetes control plane.
*   **Location**: `pkg/apis/certificates/types.go`
*   **Type**: `type CertificateSigningRequest struct`

## 2. External API Type Definitions
These are the versioned representations of the `CertificateSigningRequest` object used by clients and persisted in etcd.
*   **Version: v1** (Stable)
    *   **Location**: `staging/src/k8s.io/api/certificates/v1/types.go`
    *   **Type**: `type CertificateSigningRequest struct`
*   **Version: v1beta1**
    *   **Location**: `staging/src/k8s.io/api/certificates/v1beta1/types.go`
    *   **Type**: `type CertificateSigningRequest struct`

## 3. Validation Logic
Contains the business logic to validate `CertificateSigningRequest` objects during creation and updates.
*   **Location**: `pkg/apis/certificates/validation/validation.go`
*   **Key Functions**:
    *   `ValidateCertificateSigningRequestCreate`: Validates a CertificateSigningRequest during creation.
    *   `ValidateCertificateSigningRequestUpdate`: Validates a CertificateSigningRequest during updates.
    *   `ValidateCertificateSigningRequestStatusUpdate`: Validates status updates.
    *   `ValidateCertificateSigningRequestApprovalUpdate`: Validates approval updates.

## 4. Defaulting Logic
Applies default values to fields that are not specified by the user.
*   **Location**: Not explicitly found in `pkg/apis/certificates/v1/defaults.go`.

## 5. Registry Strategy
Defines hooks for the lifecycle of the object (PrepareForCreate, PrepareForUpdate, Canonicalize, etc.).
*   **Location**: `pkg/registry/certificates/certificates/strategy.go`
*   **Key Struct**: `csrStrategy` (implements `RESTCreateStrategy`, `RESTUpdateStrategy`, etc.)
