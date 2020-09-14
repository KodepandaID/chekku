package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
	"github.com/stretchr/testify/assert"
)

func TestIsInInvalid(t *testing.T) {
	type In struct {
		in string `chekku:"isIn:yes,no"`
	}

	e := chekku.Validate(In{
		in: "maybe",
	})

	if e == nil {
		t.Error("isIn shoud be invalid")
		return
	}

	if !assert.Equal(t, e.Error(), `"isIn", in value is not allowed`) {
		t.Error("isIn shoud be invalid")
		return
	}
}

func TestIsInValid(t *testing.T) {
	type In struct {
		in string `chekku:"isIn:yes,no"`
	}

	e := chekku.Validate(In{
		in: "no",
	})

	if e != nil {
		t.Error("isIn shoud be valid")
		return
	}
}
