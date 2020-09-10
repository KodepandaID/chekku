package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
	"github.com/stretchr/testify/assert"
)

func TestIsUUIDInvalid(t *testing.T) {
	type UUID struct {
		uuid string `chekku:"isUUID"`
	}

	invalidValues := []string{
		"hello",
		"123",
		"123-456-789",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(UUID{
			uuid: v,
		})

		if e == nil {
			t.Error("isUUID shoud be invalid")
			return
		}

		if !assert.Equal(t, e.Error(), `"isUUID", uuid must be an UUID`) {
			t.Error("isUUID shoud be invalid")
			return
		}
	}
}

func TestIsUUIDValid(t *testing.T) {
	type UUID struct {
		uuid string `chekku:"isUUID"`
	}

	// UUIDv1, UUIDv3, UUIDv4, UUIDv5
	validValues := []string{
		"17b32300-f33c-11ea-adc1-0242ac120002",
		"a3bb189e-8bf9-3888-9912-ace4e6543002",
		"e2cdc879-9e4d-42f8-9fa9-3c52e23ddd38",
		"a6eabfc1-a111-5125-a670-b746e70fdd2e",
	}

	for _, v := range *&validValues {
		e := chekku.Validate(UUID{
			uuid: v,
		})

		if e != nil {
			t.Error("isUUID shoud be valid")
			return
		}
	}
}
