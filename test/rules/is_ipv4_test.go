package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
)

func TestIsIPv4InvalidText(t *testing.T) {
	type IP struct {
		ip string `chekku:"isIPv4"`
	}

	e := chekku.Validate(IP{
		ip: "hello",
	})

	if e == nil {
		t.Error("isIPv4 shoud be invalid")
		return
	}
}

func TestIsIPv4InvalidWithInteger(t *testing.T) {
	type IP struct {
		ip int `chekku:"isIPv4"`
	}

	e := chekku.Validate(IP{
		ip: 123,
	})

	if e == nil {
		t.Error("isIPv4 shoud be invalid")
		return
	}
}

func TestIsIPv4InvalidIP(t *testing.T) {
	type IP struct {
		ip string `chekku:"isIPv4"`
	}

	e := chekku.Validate(IP{
		ip: "1.2.3.a",
	})

	if e == nil {
		t.Error("isIPv4 shoud be invalid")
		return
	}
}

func TestIsIPv4InvalidIPRules(t *testing.T) {
	type IP struct {
		ip string `chekku:"isIPv4"`
	}

	e := chekku.Validate(IP{
		ip: "1.2.3.256",
	})

	if e == nil {
		t.Error("isIPv4 shoud be invalid")
		return
	}
}

func TestIsIPv4InvalidWithIPv6(t *testing.T) {
	type IP struct {
		ip string `chekku:"isIPv4"`
	}

	e := chekku.Validate(IP{
		ip: "FEDC:BA98:7654:3210:FEDC:BA98:7654:3210",
	})

	if e == nil {
		t.Error("isIPv4 shoud be invalid")
		return
	}
}

func TestIsIPv4Valid(t *testing.T) {
	type IP struct {
		ip string `chekku:"isIPv4"`
	}

	e := chekku.Validate(IP{
		ip: "0.1.2.3",
	})

	if e == nil {
		return
	}

	t.Error("isIPv4 shoud be valid")
	return
}

func TestIsIPv4ValidSecond(t *testing.T) {
	type IP struct {
		ip string `chekku:"isIPv4"`
	}

	e := chekku.Validate(IP{
		ip: "255.255.255.255",
	})

	if e == nil {
		return
	}

	t.Error("isIPv4 shoud be valid")
	return
}
