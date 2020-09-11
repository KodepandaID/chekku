package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
	"github.com/stretchr/testify/assert"
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

	if !assert.Equal(t, e.Error(), `"minLength", text value length must be more than 5`) {
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

	if !assert.Equal(t, e.Error(), `"minLength", text value must be string`) {
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
