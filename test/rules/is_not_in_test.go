package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestIsNotInInvalid(t *testing.T) {
	type In struct {
		in string `chekku:"isNotIn:yes,no"`
	}

	e := chekku.Validate(In{
		in: "yes",
	})

	if e == nil {
		t.Error("isNotIn shoud be invalid")
		return
	}
}

func TestIsNotInValid(t *testing.T) {
	type In struct {
		in string `chekku:"isNotIn:yes,no"`
	}

	e := chekku.Validate(In{
		in: "maybe",
	})

	if e != nil {
		t.Error("isNotIn shoud be valid")
		return
	}
}
