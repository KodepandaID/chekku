package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
	"github.com/stretchr/testify/assert"
)

func TestIsNumberInvalidWithText(t *testing.T) {
	type Number struct {
		number string `chekku:"isNumber"`
	}

	e := chekku.Validate(Number{
		number: "hello",
	})

	if e == nil {
		t.Error("isNumber shoud be invalid")
		return
	}

	if !assert.Equal(t, e.Error(), `"isNumber", number must be a number`) {
		t.Error("isNumber shoud be invalid")
		return
	}
}

func TestIsNumberInvalidWithTextNumber(t *testing.T) {
	type Number struct {
		number string `chekku:"isNumber"`
	}

	e := chekku.Validate(Number{
		number: "1",
	})

	if e == nil {
		t.Error("isNumber shoud be invalid")
		return
	}

	if !assert.Equal(t, e.Error(), `"isNumber", number must be a number`) {
		t.Error("isNumber shoud be invalid")
		return
	}
}

func TestIsNumberValidWithFloat(t *testing.T) {
	type Number struct {
		number float32 `chekku:"isNumber"`
	}

	e := chekku.Validate(Number{
		number: 1.23,
	})

	if e == nil {
		return
	}

	t.Error("isNumber shoud be valid")
	return
}

func TestIsNumberValidWithInteger(t *testing.T) {
	type Number struct {
		number int `chekku:"isNumber"`
	}

	e := chekku.Validate(Number{
		number: 123,
	})

	if e == nil {
		return
	}

	t.Error("isNumber shoud be valid")
	return
}
