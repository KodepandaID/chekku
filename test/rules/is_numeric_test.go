package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
	"github.com/stretchr/testify/assert"
)

func TestIsNumericInvalidWithText(t *testing.T) {
	type Numeric struct {
		numeric string `chekku:"isNumeric"`
	}

	e := chekku.Validate(Numeric{
		numeric: "hello",
	})

	if e == nil {
		t.Error("isNumeric shoud be invalid")
		return
	}

	if !assert.Equal(t, e.Error(), `"isNumeric", numeric must be a numeric`) {
		t.Error("isNumeric shoud be invalid")
		return
	}
}

func TestIsNumericValidWithTextnumericInteger(t *testing.T) {
	type Numeric struct {
		numeric string `chekku:"isNumeric"`
	}

	e := chekku.Validate(Numeric{
		numeric: "123",
	})

	if e == nil {
		return
	}

	if !assert.Equal(t, e.Error(), `"isNumeric", numeric must be a numeric`) {
		t.Error("isNumeric shoud be invalid")
		return
	}
}

func TestIsNumericValidWithTextnumericFloat(t *testing.T) {
	type Numeric struct {
		numeric string `chekku:"isNumeric"`
	}

	e := chekku.Validate(Numeric{
		numeric: "1.23",
	})

	if e == nil {
		return
	}

	if !assert.Equal(t, e.Error(), `"isNumeric", numeric must be a numeric`) {
		t.Error("isNumeric shoud be invalid")
		return
	}
}

func TestIsNumericValidWithInteger(t *testing.T) {
	type Numeric struct {
		numeric int `chekku:"isNumeric"`
	}

	e := chekku.Validate(Numeric{
		numeric: 123,
	})

	if e == nil {
		return
	}

	if !assert.Equal(t, e.Error(), `"isNumeric", numeric must be a numeric`) {
		t.Error("isNumeric shoud be invalid")
		return
	}
}

func TestIsNumericValidWithFloat(t *testing.T) {
	type Numeric struct {
		numeric float32 `chekku:"isNumeric"`
	}

	e := chekku.Validate(Numeric{
		numeric: 1.23,
	})

	if e == nil {
		return
	}

	if !assert.Equal(t, e.Error(), `"isNumeric", numeric must be a numeric`) {
		t.Error("isNumeric shoud be invalid")
		return
	}
}
