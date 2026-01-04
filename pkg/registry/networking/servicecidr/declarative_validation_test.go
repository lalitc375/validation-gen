package servicecidr

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
		Resource:   "servicecidrs",
	})
	apitesting.VerifyValidationEquivalence(t,
		ctx,
		&networking.ServiceCIDR{
			ObjectMeta: metav1.ObjectMeta{Name: "valid-name"},
			Spec: networking.ServiceCIDRSpec{
				CIDRs: []string{"10.0.0.0/24"},
			},
		},
		func(ctx context.Context, obj runtime.Object) field.ErrorList {
			return Strategy.Validate(ctx, obj)
		},
		nil,
	)
	apitesting.VerifyValidationEquivalence(t,
		ctx,
		&networking.ServiceCIDR{
			ObjectMeta: metav1.ObjectMeta{Name: "-invalid-name"},
			Spec: networking.ServiceCIDRSpec{
				CIDRs: []string{"10.0.0.0/24"},
			},
		},
		func(ctx context.Context, obj runtime.Object) field.ErrorList {
			return Strategy.Validate(ctx, obj)
		},
		field.ErrorList{
			field.Invalid(field.NewPath("metadata", "name"), "-invalid-name", "a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character (e.g. 'example.com', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*')").WithOrigin("format=k8s-long-name"),
		},
	)
	apitesting.VerifyValidationEquivalence(t,
		ctx,
		&networking.ServiceCIDR{
			ObjectMeta: metav1.ObjectMeta{},
			Spec: networking.ServiceCIDRSpec{
				CIDRs: []string{"10.0.0.0/24"},
			},
		},
		func(ctx context.Context, obj runtime.Object) field.ErrorList {
			return Strategy.Validate(ctx, obj)
		},
		field.ErrorList{
			field.Required(field.NewPath("metadata", "name"), "name or generateName is required"),
		},
	)
	apitesting.VerifyValidationEquivalence(t,
		ctx,
		&networking.ServiceCIDR{
			ObjectMeta: metav1.ObjectMeta{Name: "valid-name"},
			Spec: networking.ServiceCIDRSpec{
				CIDRs: []string{"invalid-cidr"},
			},
		},
		func(ctx context.Context, obj runtime.Object) field.ErrorList {
			return Strategy.Validate(ctx, obj)
		},
		field.ErrorList{
			field.Invalid(field.NewPath("spec", "cidrs").Index(0), "invalid-cidr", "invalid CIDR address: invalid-cidr"),
		},
	)
	apitesting.VerifyValidationEquivalence(t,
		ctx,
		&networking.ServiceCIDR{
			ObjectMeta: metav1.ObjectMeta{Name: "valid-name"},
			Spec: networking.ServiceCIDRSpec{
				CIDRs: []string{},
			},
		},
		func(ctx context.Context, obj runtime.Object) field.ErrorList {
			return Strategy.Validate(ctx, obj)
		},
		field.ErrorList{
			field.Required(field.NewPath("spec", "cidrs"), "at least one CIDR required"),
		},
	)
	apitesting.VerifyValidationEquivalence(t,
		ctx,
		&networking.ServiceCIDR{
			ObjectMeta: metav1.ObjectMeta{Name: "valid-name"},
			Spec: networking.ServiceCIDRSpec{
				CIDRs: nil,
			},
		},
		func(ctx context.Context, obj runtime.Object) field.ErrorList {
			return Strategy.Validate(ctx, obj)
		},
		field.ErrorList{
			field.Required(field.NewPath("spec", "cidrs"), "at least one CIDR required"),
		},
	)
	apitesting.VerifyValidationEquivalence(t,
		ctx,
		&networking.ServiceCIDR{
			ObjectMeta: metav1.ObjectMeta{Name: "valid-name"},
			Spec: networking.ServiceCIDRSpec{
				CIDRs: []string{"10.0.0.0/24", "2001:db8::/64", "192.168.1.0/24"},
			},
		},
		func(ctx context.Context, obj runtime.Object) field.ErrorList {
			return Strategy.Validate(ctx, obj)
		},
		field.ErrorList{
			field.TooMany(field.NewPath("spec", "cidrs"), 3, 2).WithOrigin("maxItems"),
		},
	)
}

func TestDeclarativeUpdateValidation(t *testing.T) {
	ctx := request.WithRequestInfo(context.TODO(), &request.RequestInfo{
		APIGroup:   "networking.k8s.io",
		APIVersion: "v1",
		Resource:   "servicecidrs",
	})
	apitesting.VerifyUpdateValidationEquivalence(t,
		ctx,
		&networking.ServiceCIDR{
			ObjectMeta: metav1.ObjectMeta{Name: "valid-name", ResourceVersion: "1"},
			Spec: networking.ServiceCIDRSpec{
				CIDRs: []string{"10.0.0.0/24"},
			},
		},
		&networking.ServiceCIDR{
			ObjectMeta: metav1.ObjectMeta{Name: "valid-name", ResourceVersion: "1"},
			Spec: networking.ServiceCIDRSpec{
				CIDRs: []string{"10.0.0.0/24"},
			},
		},
		func(ctx context.Context, obj, old runtime.Object) field.ErrorList {
			return Strategy.ValidateUpdate(ctx, obj, old)
		},
		nil,
	)
}
