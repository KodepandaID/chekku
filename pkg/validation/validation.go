package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/KodepandaID/chekku/pkg/rules"
)

// Data is a user input validation structure
type Data struct {
	FieldName  string
	FieldType  reflect.Type
	FieldValue reflect.Value
	FieldTag   []string
	IsNull     bool
}

// Parse to get values, type, tag and field name
func Parse(values reflect.Value) []Data {
	d := []Data{}

	t := values.Type()
	for i := 0; i < values.NumField(); i++ {
		d = append(d, Data{
			FieldName:  t.Field(i).Name,
			FieldType:  t.Field(i).Type,
			FieldValue: values.Field(i),
			FieldTag:   strings.Split(t.Field(i).Tag.Get("chekku"), ","),
		})
	}

	return d
}

// Validate function
func Validate(inputs interface{}) error {
	values := reflect.ValueOf(inputs)
	d := Parse(values)

	r := rules.Rules{}
	for _, v := range d {
		for _, m := range v.FieldTag {
			params := make([]reflect.Value, 2)
			params[0] = reflect.ValueOf(v.FieldName)
			params[1] = reflect.ValueOf(v.FieldValue)

			result := reflect.ValueOf(r).MethodByName(strings.Title(m)).Call(params)
			if e := result[0].Interface(); e != nil {
				return fmt.Errorf("%v", e)
			}
		}
	}

	return nil
}