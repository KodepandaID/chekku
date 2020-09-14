package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestLengthBetweenInvalid(t *testing.T) {
	type Between struct {
		between string `chekku:"lengthBetween:5,10"`
	}

	invalidValues := []string{
		"hell",
		"hello world, how are you?",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Between{
			between: v,
		})

		if e == nil {
			t.Error("lengthBetween should be invalid")
			return
		}
	}
}

func TestLengthBetweenValid(t *testing.T) {
	type Between struct {
		between string `chekku:"lengthBetween:5,10"`
	}

	invalidValues := []string{
		"hello",
		"helloworld",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Between{
			between: v,
		})

		if e != nil {
			t.Error("lengthBetween should be valid")
			return
		}
	}
}
