package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
	"github.com/stretchr/testify/assert"
)

func TestIsIPv6InvalidText(t *testing.T) {
	type IP struct {
		ip string `chekku:"isIPv6"`
	}

	e := chekku.Validate(IP{
		ip: "hello",
	})

	if e == nil {
		t.Error("isIPv6 shoud be invalid")
		return
	}

	if !assert.Equal(t, e.Error(), `"isIPv6", ip value must be ip address version 6`) {
		t.Error("isIPv6 shoud be invalid")
		return
	}
}

func TestIsIPv6InvalidWithInteger(t *testing.T) {
	type IP struct {
		ip int `chekku:"isIPv6"`
	}

	e := chekku.Validate(IP{
		ip: 123,
	})

	if e == nil {
		t.Error("isIPv6 shoud be invalid")
		return
	}

	if !assert.Equal(t, e.Error(), `"isIPv6", ip value must be string`) {
		t.Error("isIPv6 shoud be invalid")
		return
	}
}

func TestIsIPv6InvalidIP(t *testing.T) {
	type IP struct {
		ip string `chekku:"isIPv6"`
	}

	invalidValues := []string{
		"2001:af40:::",
		"2001:af40:::1234",
		"2001::af40::1234",
		"1080:0:0:0:8:800:200C:417G",
	}

	for _, v := range *&invalidValues {
		e := chekku.Validate(IP{
			ip: v,
		})

		if e == nil {
			t.Error("isIPv6 shoud be invalid")
			return
		}

		if !assert.Equal(t, e.Error(), `"isIPv6", ip value must be ip address version 6`) {
			t.Error("isIPv6 shoud be invalid")
			return
		}
	}
}

func TestIsIPv6ValidIP(t *testing.T) {
	type IP struct {
		ip string `chekku:"isIPv6"`
	}

	validValues := []string{
		"2001:db8::7",
		"FEDC:BA98:7654:3210:FEDC:BA98:7654:3210",
		"FEDC:BA98:7654:3210:FEDC:BA98:7654:3210",
		"1080:0:0:0:8:800:200C:417A",
		"::1:2:3:4:5:6:7",
		"::1:2:3:4:5:6",
		"1::1:2:3:4:5:6",
		"::1:2:3:4:5",
		"1::1:2:3:4:5",
		"2:1::1:2:3:4:5",
		"::1:2:3:4",
		"1::1:2:3:4",
		"2:1::1:2:3:4",
		"3:2:1::1:2:3:4",
		"::1:2:3",
		"1::1:2:3",
		"2:1::1:2:3",
		"3:2:1::1:2:3",
		"4:3:2:1::1:2:3",
		"::1:2",
		"1::1:2",
		"2:1::1:2",
		"3:2:1::1:2",
		"4:3:2:1::1:2",
		"5:4:3:2:1::1:2",
		"::1",
		"1::1",
		"2:1::1",
		"3:2:1::1",
		"4:3:2:1::1",
		"5:4:3:2:1::1",
		"6:5:4:3:2:1::1",
		"::",
		"1::",
		"2:1::",
		"3:2:1::",
		"4:3:2:1::",
		"5:4:3:2:1::",
		"6:5:4:3:2:1::",
		"7:6:5:4:3:2:1::",
	}

	for _, v := range *&validValues {
		e := chekku.Validate(IP{
			ip: v,
		})

		if e == nil {
			return
		}

		if assert.Equal(t, e.Error(), `"isIPv6", ip value must be ip address version 6`) {
			t.Error("isIPv6 shoud be valid")
			return
		}
	}
}
