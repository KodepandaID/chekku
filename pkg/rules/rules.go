package rules

import (
	"fmt"
	"net"
	"reflect"
)

// Rules struct to run dynamic method name
type Rules struct{}

// Required to check value is empty or not becasue is required
func (r Rules) Required(fieldName string, field reflect.Value) error {
	switch field.Kind() {
	case reflect.Bool:
		if field.IsZero() {
			return fmt.Errorf("\"required\", %v is required. Bool must be true", fieldName)
		}
	case reflect.Array, reflect.Slice:
		if field.Len() == 0 {
			return fmt.Errorf("\"required\", %v is required. Array or Slice length must be more than 0", fieldName)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.IsZero() {
			return fmt.Errorf("\"required\", %v is required. Integer must be more than 0", fieldName)
		}
	case reflect.Float32, reflect.Float64:
		if field.IsZero() {
			return fmt.Errorf("\"required\", %v is required. Float must be more than 0.0", fieldName)
		}
	case reflect.String:
		if field.Len() == 0 {
			return fmt.Errorf("\"required\", %v is required. String can not be empty", fieldName)
		}
	case reflect.Interface:
		if field.IsNil() {
			return fmt.Errorf("\"required\", %v is required. Interface{} can not be nil", fieldName)
		}
	case reflect.Struct:
		if field.IsNil() {
			return fmt.Errorf("\"required\", %v is required. Struct can not be nil", fieldName)
		}
	}

	return nil
}

// IsArray to check value is an array or not
func (r Rules) IsArray(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		return nil
	default:
		return fmt.Errorf("\"isArray\", %v must be an array or slice", fieldName)
	}
}

// IsBool to check value is a boolean or not
func (r Rules) IsBool(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Bool:
		return nil
	default:
		return fmt.Errorf("\"isBool\", %v must be a boolean", fieldName)
	}
}

// IsIPv4 to check values is ipv4 or not
func (r Rules) IsIPv4(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		ip := net.ParseIP(v.String())
		if ip == nil || ip.To4() == nil {
			return fmt.Errorf("\"isIPv4\", %v value must be ip address version 4", fieldName)
		}

		return nil
	default:
		return fmt.Errorf("\"isIPv4\", %v value must be string", fieldName)
	}
}

// IsIPv6 to check values is ipv6 or not
func (r Rules) IsIPv6(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		ip := net.ParseIP(v.String())
		if ip == nil || ip.To16() == nil {
			return fmt.Errorf("\"isIPv6\", %v value must be ip address version 6", fieldName)
		}

		return nil
	default:
		return fmt.Errorf("\"isIPv6\", %v value must be string", fieldName)
	}
}

// IsNumber to check value is a number or not include integer and float
func (r Rules) IsNumber(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return nil
	case reflect.Float32, reflect.Float64:
		return nil
	default:
		return fmt.Errorf("\"isNumber\", %v must be a number", fieldName)
	}
}

// IsNumeric to check value is a numeric or not
func (r Rules) IsNumeric(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return nil
	case reflect.Float32, reflect.Float64:
		return nil
	default:
		if !numericRegex.MatchString(v.String()) {
			return fmt.Errorf("\"isNumeric\", %v must be a numeric", fieldName)
		}

		return nil
	}
}

// IsAlpha to check values is alpha or not
func (r Rules) IsAlpha(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		if !alphaRegex.MatchString(v.String()) {
			return fmt.Errorf("\"isAlpha\", %v value must be the Aa - Zz", fieldName)
		}

		return nil
	default:
		return fmt.Errorf("\"isAlpha\", %v value must be string", fieldName)
	}
}

// IsAlphanumeric to check values is alphanumeric or not
func (r Rules) IsAlphanumeric(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		if !alphaNumericRegex.MatchString(v.String()) {
			return fmt.Errorf("\"isAlphanumeric\", %v value must be the combination words and numbers", fieldName)
		}

		return nil
	default:
		return fmt.Errorf("\"isAlphanumeric\", %v value must be string", fieldName)
	}
}
