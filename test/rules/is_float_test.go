package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestIsFloatInvalidWithText(t *testing.T) {
	type Number struct {
		number string `chekku:"isFloat"`
	}

	e := chekku.Validate(Number{
		number: "hello",
	})

	if e == nil {
		t.Error("isFloat shoud be invalid")
		return
	}
}

func TestIsFloatInvalidWithInteger(t *testing.T) {
	type Number struct {
		number int `chekku:"isFloat"`
	}

	e := chekku.Validate(Number{
		number: 123,
	})

	if e == nil {
		t.Error("isFloat shoud be invalid")
		return
	}
}

func TestIsFloatValid(t *testing.T) {
	type Number struct {
		number float64 `chekku:"isFloat"`
	}

	e := chekku.Validate(Number{
		number: 0.123,
	})

	if e == nil {
		return
	}

	t.Error("isFloat shoud be valid")
	return
}
