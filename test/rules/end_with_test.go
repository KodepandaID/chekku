package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestEndWithInvalid(t *testing.T) {
	type Str struct {
		str string `chekku:"endWith:foo"`
	}

	invalidValues := []string{
		"akhbar",
		"foobar",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Str{
			str: v,
		})

		if e == nil {
			t.Error("endWith shoud be invalid")
			return
		}
	}
}

func TestEndWithValid(t *testing.T) {
	type Str struct {
		str string `chekku:"endWith:foo"`
	}

	validValues := []string{
		"bar foo",
		"fighter foo",
	}

	for _, v := range *&validValues {
		e := chekku.Validate(Str{
			str: v,
		})

		if e != nil {
			t.Error("endWith shoud be valid")
			return
		}
	}
}
