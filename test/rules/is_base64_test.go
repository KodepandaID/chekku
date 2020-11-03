package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestIsBase64Invalid(t *testing.T) {
	type Base64 struct {
		base string `chekku:"isBase64"`
	}

	invalidValues := []string{
		"123",
		"hello",
		"aGVsbG%8gd29ybGQ=",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Base64{
			base: v,
		})

		if e == nil {
			t.Error("isBase64 shoud be invalid")
			return
		}
	}
}

func TestIsBase64Valid(t *testing.T) {
	type Base64 struct {
		base string `chekku:"isBase64"`
	}

	validValues := []string{
		"aGVsbG8gd29ybGQ=",
		"a29kZXBhbmRhLmNvbQ==",
	}

	for _, v := range *&validValues {
		e := chekku.Validate(Base64{
			base: v,
		})

		if e != nil {
			t.Error("isBase64 shoud be valid")
			return
		}
	}
}
