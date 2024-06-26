// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pprofile

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpprofiles "go.opentelemetry.io/collector/pdata/internal/data/protogen/profiles/v1experimental"
)

func TestValueType_MoveTo(t *testing.T) {
	ms := generateTestValueType()
	dest := NewValueType()
	ms.MoveTo(dest)
	assert.Equal(t, NewValueType(), ms)
	assert.Equal(t, generateTestValueType(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newValueType(&otlpprofiles.ValueType{}, &sharedState)) })
	assert.Panics(t, func() { newValueType(&otlpprofiles.ValueType{}, &sharedState).MoveTo(dest) })
}

func TestValueType_CopyTo(t *testing.T) {
	ms := NewValueType()
	orig := NewValueType()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestValueType()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newValueType(&otlpprofiles.ValueType{}, &sharedState)) })
}

func TestValueType_Type(t *testing.T) {
	ms := NewValueType()
	assert.Equal(t, int64(0), ms.Type())
	ms.SetType(int64(1))
	assert.Equal(t, int64(1), ms.Type())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newValueType(&otlpprofiles.ValueType{}, &sharedState).SetType(int64(1)) })
}

func TestValueType_Unit(t *testing.T) {
	ms := NewValueType()
	assert.Equal(t, int64(0), ms.Unit())
	ms.SetUnit(int64(1))
	assert.Equal(t, int64(1), ms.Unit())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newValueType(&otlpprofiles.ValueType{}, &sharedState).SetUnit(int64(1)) })
}

func TestValueType_AggregationTemporality(t *testing.T) {
	ms := NewValueType()
	assert.Equal(t, otlpprofiles.AggregationTemporality(0), ms.AggregationTemporality())
	ms.SetAggregationTemporality(otlpprofiles.AggregationTemporality(1))
	assert.Equal(t, otlpprofiles.AggregationTemporality(1), ms.AggregationTemporality())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() {
		newValueType(&otlpprofiles.ValueType{}, &sharedState).SetAggregationTemporality(otlpprofiles.AggregationTemporality(1))
	})
}

func generateTestValueType() ValueType {
	tv := NewValueType()
	fillTestValueType(tv)
	return tv
}

func fillTestValueType(tv ValueType) {
	tv.orig.Type = int64(1)
	tv.orig.Unit = int64(1)
	tv.orig.AggregationTemporality = otlpprofiles.AggregationTemporality(1)
}
