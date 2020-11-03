package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestIsLatitudeInvalidWithString(t *testing.T) {
	type Latitude struct {
		lat string `chekku:"isLatitude"`
	}

	invalidValues := []string{
		"+90.0123",
		"123",
		"-123",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Latitude{
			lat: v,
		})

		if e == nil {
			t.Error("isLatitude shoud be invalid")
			return
		}
	}
}

func TestIsLatitudeInvalidWithFloat(t *testing.T) {
	type Latitude struct {
		lat float64 `chekku:"isLatitude"`
	}

	invalidValues := []float64{
		+90.0123,
		123,
		-123,
		106.900447,
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Latitude{
			lat: v,
		})

		if e == nil {
			t.Error("isLatitude shoud be invalid")
			return
		}
	}
}

func TestIsLatitudeValid(t *testing.T) {
	type Latitude struct {
		lat float64 `chekku:"isLatitude"`
	}

	invalidValues := []float64{
		77.11112223331,
		-6.225013,
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(Latitude{
			lat: v,
		})

		if e != nil {
			t.Error("isLatitude shoud be valid")
			return
		}
	}
}
