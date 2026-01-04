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
	apitesting.VerifyValidationEquivalence(t,
		request.WithRequestInfo(context.TODO(), &request.RequestInfo{
			APIGroup:   "networking.k8s.io",
			APIVersion: "v1",
			Resource:   "networkpolicies",
		}),
		&networking.NetworkPolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: "bar",
			},
		},
		func(ctx context.Context, obj runtime.Object) field.ErrorList {
			return Strategy.Validate(ctx, obj)
		},
		nil,
	)
}

func TestDeclarativeUpdateValidation(t *testing.T) {
	apitesting.VerifyUpdateValidationEquivalence(t,
		request.WithRequestInfo(context.TODO(), &request.RequestInfo{
			APIGroup:   "networking.k8s.io",
			APIVersion: "v1",
			Resource:   "networkpolicies",
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
