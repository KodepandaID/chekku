package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
	"github.com/stretchr/testify/assert"
)

func TestIsIntInvalidWithText(t *testing.T) {
	type Number struct {
		number string `chekku:"isInt"`
	}

	e := chekku.Validate(Number{
		number: "hello",
	})

	if e == nil {
		t.Error("isInt shoud be invalid")
		return
	}

	if !assert.Equal(t, e.Error(), `"isInt", number must be an integer`) {
		t.Error("isInt shoud be invalid")
		return
	}
}

func TestIsIntInvalidWithFloat(t *testing.T) {
	type Number struct {
		number float64 `chekku:"isInt"`
	}

	e := chekku.Validate(Number{
		number: 0.123,
	})

	if e == nil {
		t.Error("isInt shoud be invalid")
		return
	}

	if !assert.Equal(t, e.Error(), `"isInt", number must be an integer`) {
		t.Error("isInt shoud be invalid")
		return
	}
}

func TestIsIntValid(t *testing.T) {
	type Number struct {
		number int `chekku:"isInt"`
	}

	e := chekku.Validate(Number{
		number: 123,
	})

	if e == nil {
		return
	}

	t.Error("isInt shoud be valid")
	return
}
