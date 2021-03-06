package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestInvalidIntMinNumber(t *testing.T) {
	type Number struct {
		number int `chekku:"minNumber:5"`
	}

	e := chekku.Validate(Number{
		number: 4,
	})

	if e == nil {
		t.Error("minNumber shoud be invalid")
		return
	}
}

func TestInvalidFloatMinNumber(t *testing.T) {
	type Number struct {
		number float64 `chekku:"minNumber:5"`
	}

	e := chekku.Validate(Number{
		number: 4.9,
	})

	if e == nil {
		t.Error("minNumber shoud be invalid")
		return
	}
}

func TestValidIntMinNumber(t *testing.T) {
	type Number struct {
		number int `chekku:"minNumber:5"`
	}

	e := chekku.Validate(Number{
		number: 6,
	})

	if e != nil {
		t.Error("minNumber shoud be valid")
		return
	}
}

func TestValidFloatMinNumber(t *testing.T) {
	type Number struct {
		number float64 `chekku:"minNumber:5.2"`
	}

	e := chekku.Validate(Number{
		number: 5.3,
	})

	if e != nil {
		t.Error("minNumber shoud be valid")
		return
	}
}

func TestValidMinNumberNoData(t *testing.T) {
	type Number struct {
		number float64 `chekku:"minNumber:5.2"`
	}

	e := chekku.Validate(Number{
		number: 0,
	})

	if e == nil {
		t.Error("minNumber shoud be valid")
		return
	}
}
