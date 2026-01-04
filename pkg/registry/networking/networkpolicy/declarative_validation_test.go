/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package networkpolicy

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/endpoints/request"
	apitesting "k8s.io/kubernetes/pkg/api/testing"
	"k8s.io/kubernetes/pkg/apis/networking"
)

func TestDeclarativeValidation(t *testing.T) {
	ctx := request.WithRequestInfo(context.TODO(), &request.RequestInfo{
		APIGroup:   "networking.k8s.io",
		APIVersion: "v1",
		Resource:   "networkpolicies",
		Verb:       "create",
		Name:       "foo",
		Namespace:  "bar",
	})

	tests := []struct {
		name         string
		obj          *networking.NetworkPolicy
		expectedErrs field.ErrorList
	}{
		{
			name: "valid",
			obj: &networking.NetworkPolicy{
				ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
			},
			expectedErrs: field.ErrorList{},
		},
		{
			name: "missing CIDR",
			obj: &networking.NetworkPolicy{
				ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
				Spec: networking.NetworkPolicySpec{
					Ingress: []networking.NetworkPolicyIngressRule{
						{
							From: []networking.NetworkPolicyPeer{
								{
									IPBlock: &networking.IPBlock{
										CIDR: "",
									},
								},
							},
						},
					},
				},
			},
			expectedErrs: field.ErrorList{
				field.Required(field.NewPath("spec", "ingress").Index(0).Child("from").Index(0).Child("ipBlock", "cidr"), ""),
			},
		},
		{
			name: "invalid CIDR",
			obj: &networking.NetworkPolicy{
				ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
				Spec: networking.NetworkPolicySpec{
					Ingress: []networking.NetworkPolicyIngressRule{
						{
							From: []networking.NetworkPolicyPeer{
								{
									IPBlock: &networking.IPBlock{
										CIDR: "invalid-cidr",
									},
								},
							},
						},
					},
				},
			},
			expectedErrs: field.ErrorList{
				field.Invalid(field.NewPath("spec", "ingress").Index(0).Child("from").Index(0).Child("ipBlock", "cidr"), "invalid-cidr", "must be a valid CIDR value, (e.g. 10.9.8.0/24 or 2001:db8::/64)"),
			},
		},
		{
			name: "invalid CIDR in except",
			obj: &networking.NetworkPolicy{
				ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
				Spec: networking.NetworkPolicySpec{
					Ingress: []networking.NetworkPolicyIngressRule{
						{
							From: []networking.NetworkPolicyPeer{
								{
									IPBlock: &networking.IPBlock{
										CIDR:   "10.0.0.0/8",
										Except: []string{"invalid-cidr"},
									},
								},
							},
						},
					},
				},
			},
			expectedErrs: field.ErrorList{
				field.Invalid(field.NewPath("spec", "ingress").Index(0).Child("from").Index(0).Child("ipBlock", "except").Index(0), "invalid-cidr", "must be a valid CIDR value, (e.g. 10.9.8.0/24 or 2001:db8::/64)"),
			},
		},
		{
			name: "too many policyTypes",
			obj: &networking.NetworkPolicy{
				ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
				Spec: networking.NetworkPolicySpec{
					PolicyTypes: []networking.PolicyType{"Ingress", "Egress", "Ingress"},
				},
			},
			expectedErrs: field.ErrorList{
				field.TooMany(field.NewPath("spec", "policyTypes"), 3, 2).WithOrigin("maxItems"),
			},
		},
		{
			name: "invalid policyType",
			obj: &networking.NetworkPolicy{
				ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
				Spec: networking.NetworkPolicySpec{
					PolicyTypes: []networking.PolicyType{"Invalid"},
				},
			},
			expectedErrs: field.ErrorList{
				field.NotSupported(field.NewPath("spec", "policyTypes").Index(0), networking.PolicyType("Invalid"), []string{"Ingress", "Egress"}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apitesting.VerifyValidationEquivalence(t,
				ctx,
				tt.obj,
				func(ctx context.Context, obj runtime.Object) field.ErrorList {
					return Strategy.Validate(ctx, obj)
				},
				tt.expectedErrs,
			)
		})
	}
}

func TestDeclarativeUpdateValidation(t *testing.T) {
	apitesting.VerifyUpdateValidationEquivalence(t,
		request.WithRequestInfo(context.TODO(), &request.RequestInfo{
			APIGroup:   "networking.k8s.io",
			APIVersion: "v1",
			Resource:   "networkpolicies",
			Verb:       "update",
			Name:       "foo",
			Namespace:  "bar",
		}),
		&networking.NetworkPolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name:            "foo",
				Namespace:       "bar",
				ResourceVersion: "1",
			},
		},
		&networking.NetworkPolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name:            "foo",
				Namespace:       "bar",
				ResourceVersion: "1",
			},
		},
		func(ctx context.Context, obj, old runtime.Object) field.ErrorList {
			return Strategy.ValidateUpdate(ctx, obj, old)
		},
		nil,
	)
}
