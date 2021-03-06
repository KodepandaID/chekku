package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestMaxLengthInvalid(t *testing.T) {
	type Text struct {
		text string `chekku:"maxLength:11"`
	}

	e := chekku.Validate(Text{
		text: "hello worlds",
	})

	if e == nil {
		t.Error("maxLength shoud be invalid")
		return
	}
}

func TestMaxLengthInvalidWitnInteger(t *testing.T) {
	type Text struct {
		text int `chekku:"maxLength:5"`
	}

	e := chekku.Validate(Text{
		text: 123,
	})

	if e == nil {
		t.Error("maxLength shoud be invalid")
		return
	}
}

func TestMaxLengthValid(t *testing.T) {
	type Text struct {
		text string `chekku:"maxLength:11"`
	}

	e := chekku.Validate(Text{
		text: "hello world",
	})

	if e != nil {
		t.Error("maxLength shoud be valid")
		return
	}
}
