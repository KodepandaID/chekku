package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
	"github.com/stretchr/testify/assert"
)

func TestIsBoolInvalidWithText(t *testing.T) {
	type Boolean struct {
		boolean string `chekku:"isBool"`
	}

	e := chekku.Validate(Boolean{
		boolean: "hello",
	})

	if e == nil {
		t.Error("isBool shoud be invalid")
		return
	}

	if !assert.Equal(t, e.Error(), `"isBool", boolean must be a boolean`) {
		t.Error("isBool shoud be invalid")
		return
	}
}

func TestIsBoolInvalidWithInteger(t *testing.T) {
	type Boolean struct {
		boolean int `chekku:"isBool"`
	}

	e := chekku.Validate(Boolean{
		boolean: 1,
	})

	if e == nil {
		t.Error("isBool shoud be invalid")
		return
	}

	if !assert.Equal(t, e.Error(), `"isBool", boolean must be a boolean`) {
		t.Error("isBool shoud be invalid")
		return
	}
}

func TestIsBoolValidTrue(t *testing.T) {
	type Boolean struct {
		boolean bool `chekku:"isBool"`
	}

	e := chekku.Validate(Boolean{
		boolean: true,
	})

	if e == nil {
		return
	}

	t.Error("isBool shoud be invalid")
	return
}

func TestIsBoolValidFalse(t *testing.T) {
	type Boolean struct {
		boolean bool `chekku:"isBool"`
	}

	e := chekku.Validate(Boolean{
		boolean: false,
	})

	if e == nil {
		return
	}

	t.Error("isBool shoud be valid")
	return
}
