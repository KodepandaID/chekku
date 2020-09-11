package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
	"github.com/stretchr/testify/assert"
)

func TestIsDateInvalid(t *testing.T) {
	type Date struct {
		date string `chekku:"isDate"`
	}

	invalidValues := []string{
		"2020",
		"01-2020-10",
		"2020-01",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Date{
			date: v,
		})

		if e == nil {
			t.Error("isDate shoud be invalid")
			return
		}

		if !assert.Equal(t, e.Error(), `"isDate", date value must be Date format`) {
			t.Error("isDate shoud be invalid")
			return
		}
	}
}

func TestIsDateValid(t *testing.T) {
	type Date struct {
		date string `chekku:"isDate"`
	}

	validValues := []string{
		"2020-01-02",
		"2020-01-02 23:59:59",
		"02/12/2020",
		"02/12/2020 23:59:59",
		"2020-January-20",
	}

	for _, v := range *&validValues {
		e := chekku.Validate(Date{
			date: v,
		})

		if e != nil {
			t.Error("isDate shoud be valid")
			return
		}
	}
}
