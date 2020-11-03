package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestIsAlphaInvalid(t *testing.T) {
	type Alpha struct {
		alpha string `chekku:"isAlpha"`
	}

	invalidValues := []string{
		"123",
		"hello123",
		"hello@kodepanda.com",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Alpha{
			alpha: v,
		})

		if e == nil {
			t.Error("isAlpha shoud be invalid")
			return
		}
	}
}

func TestIsAlphaValid(t *testing.T) {
	type Alpha struct {
		alpha string `chekku:"isAlpha"`
	}

	validValues := []string{
		"hello",
		"Hello",
	}

	for _, v := range *&validValues {
		e := chekku.Validate(Alpha{
			alpha: v,
		})

		if e != nil {
			t.Error("isAlpha shoud be valid")
			return
		}
	}
}
