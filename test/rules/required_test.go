package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestRequiredInvalidText(t *testing.T) {
	type Text struct {
		text string `chekku:"required"`
	}

	e := chekku.Validate(Text{
		text: "",
	})

	if e == nil {
		t.Error("required should be invalid")
		return
	}
}

func TestRequiredInvalidInt(t *testing.T) {
	type Text struct {
		text int `chekku:"required"`
	}

	e := chekku.Validate(Text{
		text: 0,
	})

	if e == nil {
		t.Error("required should be invalid")
		return
	}
}

func TestRequiredInvalidBool(t *testing.T) {
	type Text struct {
		text bool `chekku:"required"`
	}

	e := chekku.Validate(Text{
		text: false,
	})

	if e == nil {
		t.Error("required should be invalid")
		return
	}
}

func TestRequiredInvalidArray(t *testing.T) {
	type Text struct {
		text []string `chekku:"required"`
	}

	e := chekku.Validate(Text{
		text: []string{},
	})

	if e == nil {
		t.Error("required should be invalid")
		return
	}
}

func TestRequiredValid(t *testing.T) {
	type Text struct {
		text string `chekku:"required"`
	}

	e := chekku.Validate(Text{
		text: "0",
	})

	if e != nil {
		t.Error("required should be valid")
		return
	}
}
