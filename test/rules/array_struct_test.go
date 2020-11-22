package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestArrayRequiredInvalidText(t *testing.T) {
	type TextStruct struct {
		text3 string `chekku:"required"`
	}

	type Text struct {
		text  string `chekku:"required"`
		text2 string
		str   TextStruct `chekku:"required"`
	}

	test := []Text{
		{
			text: "test1",
			str: TextStruct{
				text3: "test",
			},
		},
		{
			text2: "test2",
		},
	}

	e := chekku.Validate(test)

	if e == nil {
		t.Error("array struct test should be invalid")
		return
	}
}

func TestArrayRequiredValidText(t *testing.T) {
	type TextStruct struct {
		text3 string `chekku:"required"`
	}

	type Text struct {
		text  string `chekku:"required"`
		text2 string
		str   TextStruct `chekku:"required"`
	}

	test := []Text{
		{
			text: "test1",
			str: TextStruct{
				text3: "test",
			},
		},
		{
			text: "test2",
			str: TextStruct{
				text3: "test",
			},
		},
	}

	e := chekku.Validate(test)

	if e != nil {
		t.Error("array struct test should be valid")
		return
	}
}
