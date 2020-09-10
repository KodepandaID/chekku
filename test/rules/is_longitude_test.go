package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
	"github.com/stretchr/testify/assert"
)

func TestIsLongitudeInvalidWithString(t *testing.T) {
	type Longitude struct {
		long string `chekku:"isLongitude"`
	}

	invalidValues := []string{
		"182.1234",
		"-181.1234",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Longitude{
			long: v,
		})

		if e == nil {
			t.Error("isLongitude shoud be invalid")
			return
		}

		if !assert.Equal(t, e.Error(), `"isLongitude", long value must be the geolocation Longitude`) {
			t.Error("isLongitude shoud be invalid")
			return
		}
	}
}

func TestIsLongitudeInvalidWithFloat(t *testing.T) {
	type Longitude struct {
		long float64 `chekku:"isLongitude"`
	}

	invalidValues := []float64{
		182.1234,
		-181.1234,
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Longitude{
			long: v,
		})

		if e == nil {
			t.Error("isLongitude shoud be invalid")
			return
		}

		if !assert.Equal(t, e.Error(), `"isLongitude", long value must be the geolocation Longitude`) {
			t.Error("isLongitude shoud be invalid")
			return
		}
	}
}

func TestIsLongitudeValid(t *testing.T) {
	type Longitude struct {
		long float64 `chekku:"isLongitude"`
	}

	validValues := []float64{
		106.900447,
		+90.123,
		-120.1234,
	}

	for _, v := range *&validValues {
		e := chekku.Validate(Longitude{
			long: v,
		})

		if e != nil {
			t.Error("isLongitude shoud be valid")
			return
		}
	}
}
