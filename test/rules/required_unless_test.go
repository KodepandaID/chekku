package test

import (
	"fmt"
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestInvalidBoolRequiredUnless(t *testing.T) {
	type Text struct {
		text0 bool
		text1 string `chekku:"requiredUnless:text0,true"`
	}

	e := chekku.Validate(Text{
		text0: false,
	})

	if e == nil {
		t.Error("required should be invalid")
		return
	}
}

func TestInvalidIntRequiredUnless(t *testing.T) {
	type Text struct {
		text0 int
		text1 string `chekku:"requiredUnless:text0,200"`
	}

	e := chekku.Validate(Text{
		text0: 201,
	})

	if e == nil {
		t.Error("required should be invalid")
		return
	}
}

func TestInvalidFloatRequiredUnless(t *testing.T) {
	type Text struct {
		text0 float64
		text1 string `chekku:"requiredUnless:text0,1.5"`
	}

	e := chekku.Validate(Text{
		text0: 1.6,
	})

	if e == nil {
		t.Error("required should be invalid")
		return
	}
}

func TestInvalidStringRequiredUnless(t *testing.T) {
	type Text struct {
		text0 string
		text1 string `chekku:"requiredUnless:text0,hello"`
	}

	e := chekku.Validate(Text{
		text0: "hello!",
	})

	if e == nil {
		t.Error("required should be invalid")
		return
	}
}

func TestValidBoolRequiredUnless(t *testing.T) {
	type Text struct {
		text0 bool
		text1 string `chekku:"requiredUnless:text0,true"`
	}

	e := chekku.Validate(Text{
		text0: true,
	})

	if e != nil {
		t.Error("required should be valid")
		return
	}
}

func TestValidIntRequiredUnless(t *testing.T) {
	type Text struct {
		text0 int
		text1 string `chekku:"requiredUnless:text0,200"`
	}

	e := chekku.Validate(Text{
		text0: 200,
	})

	if e != nil {
		t.Error("required should be valid")
		return
	}
}

func TestValidFloatRequiredUnless(t *testing.T) {
	type Text struct {
		text0 float64
		text1 string `chekku:"requiredUnless:text0,1.5"`
	}

	e := chekku.Validate(Text{
		text0: 1.5,
	})

	if e != nil {
		t.Error("required should be valid")
		return
	}
}

func TestValidStringRequiredUnless(t *testing.T) {
	type Text struct {
		text0 string
		text1 string `chekku:"requiredUnless:text0,hello"`
	}

	e := chekku.Validate(Text{
		text0: "hello",
	})
	fmt.Print(e)

	if e != nil {
		t.Error("required should be valid")
		return
	}
}
