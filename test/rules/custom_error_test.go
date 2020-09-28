package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
	"github.com/stretchr/testify/assert"
)

func TestCustomErrorValid(t *testing.T) {
	type Error struct {
		Username string `chekku:"required|maxLength:10" errors:"required:username tidak boleh kosong|maxLength:Panjang username tidak boleh lebih dari 10 karakter"`
	}

	e := chekku.Validate(Error{
		Username: "",
	})

	if e == nil {
		t.Error("custom errors must be invalid")
		return
	}

	if e != nil && !assert.Equal(t, e.Error(), `username tidak boleh kosong`) {
		t.Error("Wrong test")
		return
	}

	e = chekku.Validate(Error{
		Username: "hello world",
	})

	if e == nil {
		t.Error("custom errors must be invalid")
		return
	}

	if e != nil && !assert.Equal(t, e.Error(), `Panjang username tidak boleh lebih dari 10 karakter`) {
		t.Error("Wrong test")
		return
	}
}
