package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
	"github.com/stretchr/testify/assert"
)

func TestStartsWithInvalid(t *testing.T) {
	type Str struct {
		str string `chekku:"startsWith:foo"`
	}

	invalidValues := []string{
		"",
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

		if !assert.Equal(t, e.Error(), `"startsWith", str value must be start with foo`) {
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
