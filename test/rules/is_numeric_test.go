package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
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

	t.Error("isNumeric shoud be invalid")
	return
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

	t.Error("isNumeric shoud be invalid")
	return
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

	t.Error("isNumeric shoud be valid")
	return
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

	t.Error("isNumeric shoud be valid")
	return
}
