package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/KodepandaID/chekku/pkg/parser"
	"github.com/KodepandaID/chekku/pkg/rules"
)

// Errors stack return
type Errors struct {
	Code   string
	Detail string
}

// Validate function
func Validate(inputs interface{}) []Errors {
	values := reflect.ValueOf(inputs)
	d := parser.Parse(values)

	r := rules.Rules{
		Inputs: values,
	}

	var eStack []Errors
	var required bool

	for _, v := range d {
		required = Find(v.FieldTag, "required")
		for _, tag := range v.FieldTag {
			if tag != "" {
				var result []reflect.Value
				tagVar := strings.Split(tag, ":")

				if len(tagVar) > 1 {
					params := make([]reflect.Value, 3)
					params[0] = reflect.ValueOf(v.FieldName)
					params[1] = reflect.ValueOf(v.FieldValue)
					params[2] = reflect.ValueOf(tagVar[1])

					result = reflect.ValueOf(r).MethodByName(strings.Title(tagVar[0])).Call(params)
				} else if len(tagVar) == 3 {
					params := make([]reflect.Value, 4)
					params[0] = reflect.ValueOf(v.FieldName)
					params[1] = reflect.ValueOf(v.FieldValue)
					params[2] = reflect.ValueOf(tagVar[1])
					params[3] = reflect.ValueOf(tagVar[2])

					result = reflect.ValueOf(r).MethodByName(strings.Title(tagVar[0])).Call(params)
				} else {
					params := make([]reflect.Value, 2)
					params[0] = reflect.ValueOf(v.FieldName)
					params[1] = reflect.ValueOf(v.FieldValue)

					result = reflect.ValueOf(r).MethodByName(strings.Title(tag)).Call(params)
				}

				if e := result[0].Interface(); e != nil {
					t := tag
					if len(v.ErrorTag) > 0 {
						if len(tagVar) > 1 {
							t = tagVar[0]
						}

						for _, et := range v.ErrorTag {
							if strings.Contains(et, t) {
								e = strings.Split(et, ":")[1]
							}
						}

					}

					if required || !required && !v.FieldValue.IsNil() {
						eStack = append(eStack, Errors{
							Code:   v.FieldName + ":" + t,
							Detail: fmt.Sprintf("%v", e),
						})
					}
				}
			}
		}
	}

	return eStack
}

// Find an element on slice
func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
