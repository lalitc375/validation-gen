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

package event

import (
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	apitesting "k8s.io/kubernetes/pkg/api/testing"
	api "k8s.io/kubernetes/pkg/apis/core"
)

func TestDeclarativeValidate(t *testing.T) {
	gvs := []schema.GroupVersion{
		{Group: "", Version: "v1"},
		{Group: "events.k8s.io", Version: "v1beta1"},
		{Group: "events.k8s.io", Version: "v1"},
	}

	for _, gv := range gvs {
		t.Run(gv.String(), func(t *testing.T) {
			ctx := genericapirequest.WithRequestInfo(genericapirequest.NewDefaultContext(), &genericapirequest.RequestInfo{
				APIGroup:   gv.Group,
				APIVersion: gv.Version,
				Resource:   "events",
			})

			reportingControllerField := "reportingController"
			if gv.Group == "" && gv.Version == "v1" {
				reportingControllerField = "reportingComponent"
			}

			// Strict validation is only enforced for events.k8s.io/v1.
			isStrict := gv.Group == "events.k8s.io" && gv.Version == "v1"

			testCases := map[string]struct {
				input              api.Event
				expectedErrs       field.ErrorList
				skipVersionedCheck bool
			}{
				"valid": {
					input: mkValidEvent(),
				},
				"invalid reportingController": {
					input: func() api.Event {
						e := mkValidEvent()
						e.ReportingController = "invalid_controller!"
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.Invalid(field.NewPath(reportingControllerField), "invalid_controller!", "").WithOrigin("format=k8s-label-key"),
					},
				},
			}

			if isStrict {
				testCases["too long reportingInstance"] = struct {
					input              api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					input: func() api.Event {
						e := mkValidEvent()
						e.ReportingInstance = "a" + string(make([]byte, 128))
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.TooLong(field.NewPath("reportingInstance"), "", 128).WithOrigin("maxLength"),
					},
					skipVersionedCheck: true,
				}
				testCases["missing reportingController"] = struct {
					input              api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					input: func() api.Event {
						e := mkValidEvent()
						e.ReportingController = ""
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.Required(field.NewPath(reportingControllerField), ""),
					},
					skipVersionedCheck: true,
				}
				testCases["missing reportingInstance"] = struct {
					input              api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					input: func() api.Event {
						e := mkValidEvent()
						e.ReportingInstance = ""
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.Required(field.NewPath("reportingInstance"), ""),
					},
					skipVersionedCheck: true,
				}
				testCases["missing action"] = struct {
					input              api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					input: func() api.Event {
						e := mkValidEvent()
						e.Action = ""
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.Required(field.NewPath("action"), ""),
					},
					skipVersionedCheck: true,
				}
				testCases["too long action"] = struct {
					input              api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					input: func() api.Event {
						e := mkValidEvent()
						e.Action = "a" + string(make([]byte, 128))
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.TooLong(field.NewPath("action"), "", 128).WithOrigin("maxLength"),
					},
					skipVersionedCheck: true,
				}
				testCases["missing reason"] = struct {
					input              api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					input: func() api.Event {
						e := mkValidEvent()
						e.Reason = ""
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.Required(field.NewPath("reason"), ""),
					},
					skipVersionedCheck: true,
				}
				testCases["too long reason"] = struct {
					input              api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					input: func() api.Event {
						e := mkValidEvent()
						e.Reason = "a" + string(make([]byte, 128))
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.TooLong(field.NewPath("reason"), "", 128).WithOrigin("maxLength"),
					},
					skipVersionedCheck: true,
				}
				testCases["missing type"] = struct {
					input              api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					input: func() api.Event {
						e := mkValidEvent()
						e.Type = ""
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.Required(field.NewPath("type"), ""),
					},
					skipVersionedCheck: true,
				}
				testCases["too long note"] = struct {
					input              api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					input: func() api.Event {
						e := mkValidEvent()
						e.Message = "a" + string(make([]byte, 1024))
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.TooLong(field.NewPath("message"), "", 1024).WithOrigin("maxLength"),
					},
					skipVersionedCheck: true,
				}
			}

			for k, tc := range testCases {
				t.Run(k, func(t *testing.T) {
					if tc.skipVersionedCheck {
						errs := Strategy.Validate(ctx, &tc.input)
						if len(tc.expectedErrs) > 0 {
							field.ErrorMatcher{}.ByType().ByOrigin().ByFieldNormalized(eventNormalizationRules).ByDeclarativeNative().Test(t, tc.expectedErrs, errs)
						} else if len(errs) != 0 {
							t.Errorf("expected no errors, but got: %v", errs)
						}
					} else {
						apitesting.VerifyValidationEquivalence(t, ctx, &tc.input, Strategy.Validate, tc.expectedErrs, apitesting.WithNormalizationRules(eventNormalizationRules...))
					}
				})
			}
		})
	}
}

func TestDeclarativeValidateUpdate(t *testing.T) {
	gvs := []schema.GroupVersion{
		{Group: "", Version: "v1"},
		{Group: "events.k8s.io", Version: "v1beta1"},
		{Group: "events.k8s.io", Version: "v1"},
	}

	for _, gv := range gvs {
		t.Run(gv.String(), func(t *testing.T) {
			ctx := genericapirequest.WithRequestInfo(genericapirequest.NewDefaultContext(), &genericapirequest.RequestInfo{
				APIGroup:   gv.Group,
				APIVersion: gv.Version,
				Resource:   "events",
			})

			validObj := mkValidEvent()
			testCases := map[string]struct {
				old                api.Event
				update             api.Event
				expectedErrs       field.ErrorList
				skipVersionedCheck bool
			}{
				"valid": {
					old:    validObj,
					update: validObj,
				},
			}

			reportingControllerField := "reportingController"
			if gv.Group == "" && gv.Version == "v1" {
				reportingControllerField = "reportingComponent"
			}

			// Immutability is only enforced for events.k8s.io/v1 and newer.
			if gv.Group == "events.k8s.io" && gv.Version == "v1" {
				testCases["immutable reportingController"] = struct {
					old                api.Event
					update             api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					old: validObj,
					update: func() api.Event {
						e := mkValidEvent()
						e.ReportingController = "other-controller"
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.Invalid(field.NewPath(reportingControllerField), "other-controller", "field is immutable").WithOrigin("immutable"),
					},
					skipVersionedCheck: true,
				}
				testCases["immutable reportingInstance"] = struct {
					old                api.Event
					update             api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					old: validObj,
					update: func() api.Event {
						e := mkValidEvent()
						e.ReportingInstance = "other-instance"
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.Invalid(field.NewPath("reportingInstance"), "other-instance", "field is immutable").WithOrigin("immutable"),
					},
					skipVersionedCheck: true,
				}
				testCases["immutable action"] = struct {
					old                api.Event
					update             api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					old: validObj,
					update: func() api.Event {
						e := mkValidEvent()
						e.Action = "OtherAction"
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.Invalid(field.NewPath("action"), "OtherAction", "field is immutable").WithOrigin("immutable"),
					},
					skipVersionedCheck: true,
				}
				testCases["immutable reason"] = struct {
					old                api.Event
					update             api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					old: validObj,
					update: func() api.Event {
						e := mkValidEvent()
						e.Reason = "OtherReason"
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.Invalid(field.NewPath("reason"), "OtherReason", "field is immutable").WithOrigin("immutable"),
					},
					skipVersionedCheck: true,
				}
				testCases["immutable eventTime"] = struct {
					old                api.Event
					update             api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					old: validObj,
					update: func() api.Event {
						e := mkValidEvent()
						e.EventTime = metav1.MicroTime{Time: fixedTime.Add(1 * time.Hour)}
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.Invalid(field.NewPath("eventTime"), metav1.MicroTime{Time: fixedTime.Add(1 * time.Hour)}, "field is immutable").WithOrigin("immutable"),
					},
					skipVersionedCheck: true,
				}
				testCases["immutable type"] = struct {
					old                api.Event
					update             api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					old: validObj,
					update: func() api.Event {
						e := mkValidEvent()
						e.Type = "Warning"
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.Invalid(field.NewPath("type"), "Warning", "field is immutable").WithOrigin("immutable"),
					},
					skipVersionedCheck: true,
				}
				testCases["immutable regarding"] = struct {
					old                api.Event
					update             api.Event
					expectedErrs       field.ErrorList
					skipVersionedCheck bool
				}{
					old: validObj,
					update: func() api.Event {
						e := mkValidEvent()
						e.InvolvedObject.Name = "other-pod"
						return e
					}(),
					expectedErrs: field.ErrorList{
						field.Invalid(field.NewPath("involvedObject"), api.ObjectReference{Kind: "Pod", Name: "other-pod", Namespace: "default"}, "field is immutable").WithOrigin("immutable"),
					},
					skipVersionedCheck: true,
				}
			}

			for k, tc := range testCases {
				t.Run(k, func(t *testing.T) {
					// Ensure ResourceVersion is set for update validation
					tc.old.ResourceVersion = "1"
					tc.update.ResourceVersion = "2"

					if tc.skipVersionedCheck {
						errs := Strategy.ValidateUpdate(ctx, &tc.update, &tc.old)
						if len(tc.expectedErrs) > 0 {
							field.ErrorMatcher{}.ByType().ByOrigin().ByFieldNormalized(eventNormalizationRules).ByDeclarativeNative().Test(t, tc.expectedErrs, errs)
						} else if len(errs) != 0 {
							t.Errorf("expected no errors, but got: %v", errs)
						}
					} else {
						apitesting.VerifyUpdateValidationEquivalence(t, ctx, &tc.update, &tc.old, Strategy.ValidateUpdate, tc.expectedErrs, apitesting.WithNormalizationRules(eventNormalizationRules...))
					}
				})
			}
		})
	}
}

var fixedTime = time.Date(2025, 12, 30, 0, 0, 0, 0, time.UTC)

func mkValidEvent() api.Event {
	return api.Event{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-event",
			Namespace: "default",
		},
		InvolvedObject: api.ObjectReference{
			Kind:      "Pod",
			Name:      "test-pod",
			Namespace: "default",
		},
		EventTime:           metav1.MicroTime{Time: fixedTime},
		ReportingController: "test-controller",
		ReportingInstance:   "test-instance",
		Action:              "Test",
		Reason:              "Testing",
		Type:                "Normal",
	}
}

var eventNormalizationRules = []field.NormalizationRule{}