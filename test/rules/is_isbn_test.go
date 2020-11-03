package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestIsISBNInvalid(t *testing.T) {
	type Isbn struct {
		isbn string `chekku:"isISBN"`
	}

	invalidValues := []string{
		"123",
		"hello123",
		"1-123-456789-0",
		"979-1-23456-789-6",
		"ISBX 979-1-23456-789-6",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Isbn{
			isbn: v,
		})

		if e == nil {
			t.Error("isISBN shoud be invalid")
			return
		}
	}
}

func TestIsISBNValid(t *testing.T) {
	type Isbn struct {
		isbn string `chekku:"isISBN"`
	}

	validValues := []string{
		"ISBN 90-70002-34-5",
		"ISBN 1-23456-789-X",
		"ISBN 979-1-23456-789-6",
	}

	for _, v := range *&validValues {
		e := chekku.Validate(Isbn{
			isbn: v,
		})

		if e != nil {
			t.Error("isISBN shoud be valid")
			return
		}
	}
}
