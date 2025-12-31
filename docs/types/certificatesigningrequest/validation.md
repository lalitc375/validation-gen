# Validation: CertificateSigningRequest

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `spec.request` | `[]byte` | `+k8s:required`<br>`+k8s:immutable` | CSR data is required and immutable after creation. |
| `spec.signerName` | `string` | `+k8s:required`<br>`+k8s:format=k8s-label-key`<br>`+k8s:immutable` | Signer name is required, must be a qualified name, and is immutable. |
| `spec.expirationSeconds` | `*int32` | `+k8s:optional`<br>`+k8s:minimum=600`<br>`+k8s:immutable` | Optional, minimum 10 minutes, immutable. |
| `spec.usages` | `[]KeyUsage` | `+k8s:required`<br>`+k8s:listType=atomic`<br>`+k8s:unique=set`<br>`+k8s:immutable` | At least one usage is required, items must be unique, and field is immutable. Underlying type is an enum. |
| `spec.username` | `string` | `+k8s:optional`<br>`+k8s:immutable` | Populated by the server, immutable. |
| `spec.uid` | `string` | `+k8s:optional`<br>`+k8s:immutable` | Populated by the server, immutable. |
| `spec.groups` | `[]string` | `+k8s:optional`<br>`+k8s:listType=atomic`<br>`+k8s:immutable` | Populated by the server, immutable. |
| `spec.extra` | `map[string]ExtraValue` | `+k8s:optional`<br>`+k8s:immutable` | Populated by the server, immutable. |
| `status.conditions` | `[]CertificateSigningRequestCondition` | `+k8s:listType=map`<br>`+k8s:listMapKey=type`<br>`+k8s:customUnique`<br>`+k8s:item(type="Approved")=+k8s:validation:zeroOrOneOfMember="approval"`<br>`+k8s:item(type="Denied")=+k8s:validation:zeroOrOneOfMember="approval"` | Managed via specialized update paths; Approved and Denied are mutually exclusive. |
| `status.certificate` | `[]byte` | `+k8s:optional`<br>`+k8s:update=NoModify` | Once set, the certificate cannot be modified (NoModify payload used here to reflect "updates may not modify existing certificate content"). |
