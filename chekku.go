package chekku

import "github.com/KodepandaID/chekku/pkg/validation"

// Validate function
func Validate(inputs interface{}) []validation.Errors {
	e := validation.Validate(inputs)
	return e
}
