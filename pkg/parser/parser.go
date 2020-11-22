package parser

import (
	"reflect"
	"strings"
)

// Data is a user input validation structure
type Data struct {
	FieldName  string
	FieldType  reflect.Type
	FieldValue reflect.Value
	FieldTag   []string
	ErrorTag   []string
	IsNull     bool
}

// Parse to get values, type, tag and field name
func Parse(values reflect.Value) []Data {
	d := []Data{}

	switch values.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		for i := 0; i < values.Len(); i++ {
			t := values.Index(i).Type()
			for j := 0; j < values.Index(i).NumField(); j++ {
				d = append(d, Data{
					FieldName:  t.Field(j).Name,
					FieldType:  t.Field(j).Type,
					FieldValue: values.Index(i).Field(j),
					FieldTag:   strings.Split(t.Field(j).Tag.Get("chekku"), "|"),
					ErrorTag:   strings.Split(t.Field(j).Tag.Get("errors"), "|"),
				})
			}
		}
	default:
		t := values.Type()
		for i := 0; i < values.NumField(); i++ {
			d = append(d, Data{
				FieldName:  t.Field(i).Name,
				FieldType:  t.Field(i).Type,
				FieldValue: values.Field(i),
				FieldTag:   strings.Split(t.Field(i).Tag.Get("chekku"), "|"),
				ErrorTag:   strings.Split(t.Field(i).Tag.Get("errors"), "|"),
			})
		}
	}

	return d
}
