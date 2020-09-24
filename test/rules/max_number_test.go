package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
	"github.com/stretchr/testify/assert"
)

func TestInvalidIntMaxNumber(t *testing.T) {
	type Number struct {
		number int `chekku:"maxNumber:5"`
	}

	e := chekku.Validate(Number{
		number: 6,
	})

	if e == nil {
		t.Error("maxNumber shoud be invalid")
		return
	}

	if !assert.Equal(t, e.Error(), `"maxNumber", number value must be less than 5`) {
		t.Error("maxNumber shoud be invalid")
		return
	}
}

func TestInvalidFloatMaxNumber(t *testing.T) {
	type Number struct {
		number float64 `chekku:"maxNumber:5"`
	}

	e := chekku.Validate(Number{
		number: 5.1,
	})

	if e == nil {
		t.Error("maxNumber shoud be invalid")
		return
	}

	if !assert.Equal(t, e.Error(), `"maxNumber", number value must be less than 5`) {
		t.Error("maxNumber shoud be invalid")
		return
	}
}

func TestValidIntMaxNumber(t *testing.T) {
	type Number struct {
		number int `chekku:"maxNumber:5"`
	}

	e := chekku.Validate(Number{
		number: 4,
	})

	if e != nil {
		t.Error("maxNumber shoud be valid")
		return
	}
}

func TestValidFloatMaxNumber(t *testing.T) {
	type Number struct {
		number float64 `chekku:"maxNumber:5.2"`
	}

	e := chekku.Validate(Number{
		number: 5.1,
	})

	if e != nil {
		t.Error("maxNumber shoud be valid")
		return
	}
}
