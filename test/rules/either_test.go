package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestInvalidEither(t *testing.T) {
	type IP struct {
		ip string `chekku:"either:isIPv4,isIPv6"`
	}

	e := chekku.Validate(IP{
		ip: "not an IP Address",
	})

	if e == nil {
		t.Error("either shoud be invalid")
		return
	}
}

func TestValidEither(t *testing.T) {
	type IP struct {
		ip string `chekku:"either:isIPv4,isIPv6"`
	}

	validValues := []string{"1.2.3.4", "::1"}
	for _, v := range validValues {
		e := chekku.Validate(IP{
			ip: v,
		})
		if e != nil {
			t.Error("either should be valid")
			return
		}
	}
}
