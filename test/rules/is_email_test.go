package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestIsEmailInvalid(t *testing.T) {
	type Email struct {
		email string `chekku:"isEmail"`
	}

	invalidValues := []string{
		"123",
		"hello",
		"kodepanda.com",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Email{
			email: v,
		})

		if e == nil {
			t.Error("isEmail shoud be invalid")
			return
		}
	}
}

func TestIsEmailValid(t *testing.T) {
	type Email struct {
		email string `chekku:"isEmail"`
	}

	e := chekku.Validate(Email{
		email: "hello@kodepanda.com",
	})

	if e != nil {
		t.Error("isEmail shoud be valid")
		return
	}
}
