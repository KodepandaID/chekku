package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestInvalidBoolRequiredIf(t *testing.T) {
	type Text struct {
		text0 bool
		text1 string `chekku:"requiredIf:text0,true"`
	}

	e := chekku.Validate(Text{
		text0: true,
	})

	if e == nil {
		t.Error("required should be invalid")
		return
	}
}

func TestInvalidIntRequiredIf(t *testing.T) {
	type Text struct {
		text0 int
		text1 string `chekku:"requiredIf:text0,200"`
	}

	e := chekku.Validate(Text{
		text0: 200,
	})

	if e == nil {
		t.Error("required should be invalid")
		return
	}
}

func TestInvalidFloatRequiredIf(t *testing.T) {
	type Text struct {
		text0 float64
		text1 string `chekku:"requiredIf:text0,1.5"`
	}

	e := chekku.Validate(Text{
		text0: 1.5,
	})

	if e == nil {
		t.Error("required should be invalid")
		return
	}
}

func TestInvalidStringRequiredIf(t *testing.T) {
	type Text struct {
		text0 string
		text1 string `chekku:"requiredIf:text0,hello"`
	}

	e := chekku.Validate(Text{
		text0: "hello",
	})

	if e == nil {
		t.Error("required should be invalid")
		return
	}
}

func TestValidBoolRequiredIf(t *testing.T) {
	type Text struct {
		text0 bool
		text1 string `chekku:"requiredIf:text0,true"`
	}

	e := chekku.Validate(Text{
		text0: true,
		text1: "hello",
	})

	if e != nil {
		t.Error("required should be valid")
		return
	}
}

func TestValidIntRequiredIf(t *testing.T) {
	type Text struct {
		text0 int
		text1 string `chekku:"requiredIf:text0,200"`
	}

	e := chekku.Validate(Text{
		text0: 200,
		text1: "hello",
	})

	if e != nil {
		t.Error("required should be valid")
		return
	}
}

func TestValidFloatRequiredIf(t *testing.T) {
	type Text struct {
		text0 float64
		text1 string `chekku:"requiredIf:text0,1.5"`
	}

	e := chekku.Validate(Text{
		text0: 1.5,
		text1: "hello",
	})

	if e != nil {
		t.Error("required should be valid")
		return
	}
}

func TestValidStringRequiredIf(t *testing.T) {
	type Text struct {
		text0 string
		text1 string `chekku:"requiredIf:text0,hello"`
	}

	e := chekku.Validate(Text{
		text0: "hello",
		text1: "hello world",
	})

	if e != nil {
		t.Error("required should be valid")
		return
	}
}
