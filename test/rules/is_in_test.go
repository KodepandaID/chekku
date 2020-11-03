package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
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
