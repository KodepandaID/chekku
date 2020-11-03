package test

import (
	"testing"

	"github.com/KodepandaID/chekku"
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

	e = chekku.Validate(Error{
		Username: "hello world",
	})

	if e == nil {
		t.Error("custom errors must be invalid")
		return
	}
}
