package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestIsAlphanumericInvalid(t *testing.T) {
	type AlphaNumeric struct {
		alphanumeric string `chekku:"isAlphanumeric"`
	}

	invalidValues := []string{
		"hello@kodepanda.com",
		"Halo@@",
		"Here!!!",
		"~@#$%^&*()",
		".,/?{}",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(AlphaNumeric{
			alphanumeric: v,
		})

		if e == nil {
			t.Error("isAlphanumeric shoud be invalid")
			return
		}
	}
}

func TestIsAlphanumericInvalidWithInteger(t *testing.T) {
	type AlphaNumeric struct {
		alphanumeric int `chekku:"isAlphanumeric"`
	}

	e := chekku.Validate(AlphaNumeric{
		alphanumeric: 123,
	})

	if e == nil {
		t.Error("isAlphanumeric shoud be invalid")
		return
	}
}

func TestIsAlphanumericValid(t *testing.T) {
	type AlphaNumeric struct {
		alphanumeric string `chekku:"isAlphanumeric"`
	}

	validValues := []string{
		"123",
		"hello",
		"hello123",
	}

	for _, v := range *&validValues {
		e := chekku.Validate(AlphaNumeric{
			alphanumeric: v,
		})

		if e != nil {
			t.Error("isAlphanumeric shoud be valid")
			return
		}
	}
}
