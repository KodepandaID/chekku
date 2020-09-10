package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
	"github.com/stretchr/testify/assert"
)

func TestIsArrayInvalidWithText(t *testing.T) {
	type Array struct {
		array string `chekku:"isArray"`
	}

	e := chekku.Validate(Array{
		array: "hello",
	})

	if e == nil {
		t.Error("isArray shoud be invalid")
		return
	}

	if !assert.Equal(t, e.Error(), `"isArray", array must be an array or slice`) {
		t.Error("isArray shoud be invalid")
		return
	}
}

func TestIsArrayInvalidWithInteger(t *testing.T) {
	type Array struct {
		array int `chekku:"isArray"`
	}

	e := chekku.Validate(Array{
		array: 123,
	})

	if e == nil {
		t.Error("isArray shoud be invalid")
		return
	}

	if !assert.Equal(t, e.Error(), `"isArray", array must be an array or slice`) {
		t.Error("isArray shoud be invalid")
		return
	}
}

func TestIsArrayValid(t *testing.T) {
	type Array struct {
		array []string `chekku:"isArray"`
	}

	e := chekku.Validate(Array{
		array: []string{},
	})

	if e == nil {
		return
	}

	t.Error("isArray shoud be valid")
	return
}

func TestIsArrayValidWithSlice(t *testing.T) {
	type Array struct {
		array map[string]string `chekku:"isArray"`
	}

	e := chekku.Validate(Array{
		array: map[string]string{},
	})

	if e == nil {
		return
	}

	t.Error("isArray shoud be valid")
	return
}
