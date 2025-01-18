// Code created by gotmpl. DO NOT MODIFY.
// source: internal/shared/matchers/expecter.go.tmpl

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package matchers // import "go.opentelemetry.io/otel/exporters/zipkin/internal/matchers"

import (
	"testing"
)

type Expecter struct {
	t *testing.T
}

func NewExpecter(t *testing.T) *Expecter {
	return &Expecter{
		t: t,
	}
}

func (a *Expecter) Expect(actual interface{}) *Expectation {
	return &Expectation{
		t:      a.t,
		actual: actual,
	}
}
