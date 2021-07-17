package rules

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// RequiredIf to check another field is a match value
func (r Rules) RequiredIf(fieldName string, v reflect.Value, m string) error {
	condition := strings.Split(m, ",")
	e := r.Required(fieldName, v)

	switch v.Kind() {
	case reflect.Bool:
		b, _ := strconv.ParseBool(condition[1])
		if v.Bool() == b && e != nil {
			return fmt.Errorf("%v", e)
		}

		return nil
	case reflect.Float32, reflect.Float64:
		f, _ := strconv.ParseFloat(condition[1], 64)
		if v.Float() == f && e != nil {
			return fmt.Errorf("%v", e)
		}

		return nil
	case reflect.String:
		if v.String() == string(condition[1]) && e != nil {
			return fmt.Errorf("%v", e)
		}

		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, _ := strconv.ParseInt(condition[1], 10, 64)
		if v.Int() == i && e != nil {
			return fmt.Errorf("%v", e)
		}

		return nil
	default:
		if v.IsNil() {
			return fmt.Errorf("\"required\", %v is required", fieldName)
		}

		return nil
	}
}

// RequiredUnless to check another field is a doesn't match value
func (r Rules) RequiredUnless(fieldName string, v reflect.Value, m string) error {
	condition := strings.Split(m, ",")
	e := r.Required(fieldName, v)

	switch v.Kind() {
	case reflect.Bool:
		b, _ := strconv.ParseBool(condition[1])
		if v.Bool() != b && e != nil {
			return fmt.Errorf("%v", e)
		}

		return nil
	case reflect.Float32, reflect.Float64:
		f, _ := strconv.ParseFloat(condition[1], 64)
		if v.Float() != f && e != nil {
			return fmt.Errorf("%v", e)
		}

		return nil
	case reflect.String:
		if v.String() != string(condition[1]) && e != nil {
			return fmt.Errorf("%v", e)
		}

		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, _ := strconv.ParseInt(condition[1], 10, 64)
		if v.Int() != i && e != nil {
			return fmt.Errorf("%v", e)
		}

		return nil
	default:
		if v.IsNil() {
			return fmt.Errorf("\"required\", %v is required", fieldName)
		}

		return nil
	}
}
