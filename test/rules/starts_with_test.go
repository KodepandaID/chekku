package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestStartsWithInvalid(t *testing.T) {
	type Str struct {
		str string `chekku:"startsWith:foo"`
	}

	invalidValues := []string{
		"afbar",
		"barfoo",
		"fobar",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Str{
			str: v,
		})

		if e == nil {
			t.Error("startsWith shoud be invalid")
			return
		}
	}
}

func TestStartsWithValid(t *testing.T) {
	type Str struct {
		str string `chekku:"startsWith:foo"`
	}

	invalidValues := []string{
		"foobar",
		"foo fighter",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Str{
			str: v,
		})

		if e != nil {
			t.Error("startsWith shoud be valid")
			return
		}
	}
}
