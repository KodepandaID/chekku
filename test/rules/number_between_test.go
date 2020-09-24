package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestInvalidIntNumberBetween(t *testing.T) {
	type Number struct {
		number int `chekku:"numberBetween:5, 10"`
	}

	invalidValues := []int{3, 4, 11, 12}
	for _, v := range *&invalidValues {
		e := chekku.Validate(Number{
			number: v,
		})

		if e == nil {
			t.Error("numberBetween should be invalid")
			return
		}
	}
}

func TestInvalidFloatNumberBetween(t *testing.T) {
	type Number struct {
		number float64 `chekku:"numberBetween:5.1, 5.3"`
	}

	invalidValues := []float64{5, 5.4}
	for _, v := range *&invalidValues {
		e := chekku.Validate(Number{
			number: v,
		})

		if e == nil {
			t.Error("numberBetween should be invalid")
			return
		}
	}
}

func TestValidIntNumberBetween(t *testing.T) {
	type Number struct {
		number int `chekku:"numberBetween:5,10"`
	}

	invalidValues := []int{5, 7, 10}
	for _, v := range *&invalidValues {
		e := chekku.Validate(Number{
			number: v,
		})

		if e != nil {
			t.Error("numberBetween should be valid")
			return
		}
	}
}

func TestValidFloatNumberBetween(t *testing.T) {
	type Number struct {
		number float64 `chekku:"numberBetween:5.1,5.5"`
	}

	invalidValues := []float64{5.1, 5.3, 5.5}
	for _, v := range *&invalidValues {
		e := chekku.Validate(Number{
			number: v,
		})

		if e != nil {
			t.Error("numberBetween should be valid")
			return
		}
	}
}
