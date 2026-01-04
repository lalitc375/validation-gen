/*
Copyright 2025 The Kubernetes Authors.

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

package serviceaccount

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/endpoints/request"
	apitesting "k8s.io/kubernetes/pkg/api/testing"
	"k8s.io/kubernetes/pkg/apis/core"
)

func TestDeclarativeValidation(t *testing.T) {
	ctx := request.WithRequestInfo(context.TODO(), &request.RequestInfo{
		APIGroup:   "",
		APIVersion: "v1",
		Resource:   "serviceaccounts",
		Verb:       "create",
		Name:       "foo",
		Namespace:  "bar",
	})

	tests := []struct {
		name         string
		obj          *core.ServiceAccount
		expectedErrs field.ErrorList
	}{
		{
			name: "valid",
			obj: &core.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
			},
			expectedErrs: field.ErrorList{},
		},
		{
			name: "missing name",
			obj: &core.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{Namespace: "bar"},
			},
			expectedErrs: field.ErrorList{
				field.Required(field.NewPath("metadata", "name"), ""),
			},
		},
		{
			name: "invalid name",
			obj: &core.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{Name: "Foo", Namespace: "bar"},
			},
			expectedErrs: field.ErrorList{
				field.Invalid(field.NewPath("metadata", "name"), "Foo", "a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character (e.g. 'example.com', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*')").WithOrigin("format=k8s-long-name"),
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
			APIGroup:   "",
			APIVersion: "v1",
			Resource:   "serviceaccounts",
			Verb:       "update",
			Name:       "foo",
			Namespace:  "bar",
		}),
		&core.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:            "foo",
				Namespace:       "bar",
				ResourceVersion: "1",
			},
		},
		&core.ServiceAccount{
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
