package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestMinLengthInvalid(t *testing.T) {
	type Text struct {
		text string `chekku:"minLength:5"`
	}

	e := chekku.Validate(Text{
		text: "hell",
	})

	if e == nil {
		t.Error("minLength shoud be invalid")
		return
	}
}

func TestMinLengthInvalidWitnInteger(t *testing.T) {
	type Text struct {
		text int `chekku:"minLength:5"`
	}

	e := chekku.Validate(Text{
		text: 123,
	})

	if e == nil {
		t.Error("minLength shoud be invalid")
		return
	}
}

func TestMinLengthValid(t *testing.T) {
	type Text struct {
		text string `chekku:"minLength:5"`
	}

	e := chekku.Validate(Text{
		text: "hello",
	})

	if e != nil {
		t.Error("minLength shoud be valid")
		return
	}
}

func TestMinLengthValidNoData(t *testing.T) {
	type Text struct {
		text string `chekku:"minLength:5"`
	}

	e := chekku.Validate(Text{
		text: "",
	})

	if e != nil {
		t.Error("minLength shoud be valid")
		return
	}
}
