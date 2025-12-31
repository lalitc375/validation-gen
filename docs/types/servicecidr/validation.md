# Validation: ServiceCIDR

| Field Path | Go Type | Validation Tags | Reasoning / Notes |
| :--- | :--- | :--- | :--- |
| `metadata` | `metav1.ObjectMeta` | `+k8s:subfield(name=name)=+k8s:required`<br>`+k8s:subfield(name=name)=+k8s:format=k8s-long-name` | ServiceCIDR name is required and must be a DNS subdomain. |
| `spec.cidrs` | `[]string` | `+k8s:required`<br>`+k8s:minItems=1`<br>`+k8s:maxItems=2`<br>`+k8s:eachVal=+k8s:format=k8s-cidr` | Mandatory list of CIDRs. Max 2 (one for each IP family). |
