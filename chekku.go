package chekku

import "github.com/KodepandaID/chekku/pkg/validation"

// Validate function
func Validate(inputs interface{}) error {
	e := validation.Validate(inputs)
	return e
}
