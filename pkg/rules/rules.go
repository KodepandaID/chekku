package rules

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"
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

// IsInt to check value is a integer like int, int6, int8, int16, int32 or int64
func (r Rules) IsInt(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return nil
	default:
		return fmt.Errorf("\"isInt\", %v must be an integer", fieldName)
	}
}

// IsFloat to check value is a float32 or float64
func (r Rules) IsFloat(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return nil
	default:
		return fmt.Errorf("\"isFloat\", %v must be a float", fieldName)
	}
}

// IsUUID to check values is uuid, uuid3, uuid4 or uuid5
func (r Rules) IsUUID(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		if !uUIDRegex.MatchString(v.String()) && !uUID3Regex.MatchString(v.String()) && !uUID4Regex.MatchString(v.String()) && !uUID5Regex.MatchString(v.String()) {
			return fmt.Errorf("\"isUUID\", %v must be an UUID", fieldName)
		}

		return nil
	default:
		return fmt.Errorf("\"isUUID\", %v value must be string", fieldName)
	}
}

// IsISBN to check values is ISBN or not
func (r Rules) IsISBN(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		if !iSBNRegex.MatchString(v.String()) {
			return fmt.Errorf("\"isISBN\", %v value must be ISBN", fieldName)
		}

		return nil
	default:
		return fmt.Errorf("\"isISBN\", %v value must be string", fieldName)
	}
}

// IsBase64 to check values is base64 or not
func (r Rules) IsBase64(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		if !base64Regex.MatchString(v.String()) || !base64URLRegex.MatchString(v.String()) {
			return fmt.Errorf("\"isBase64\", %v value must be base64", fieldName)
		}

		return nil
	default:
		return fmt.Errorf("\"isBase64\", %v value must be string", fieldName)
	}
}

// IsDate to check values is date valid format or not
func (r Rules) IsDate(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		for _, l := range *&layout {
			_, e := time.Parse(l, v.String())
			if e == nil {
				return nil
			}
		}

		return fmt.Errorf("\"isDate\", %v value must be Date format", fieldName)
	default:
		return fmt.Errorf("\"isDate\", %v value must be string", fieldName)
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

// IsEmail to check values is email address or not
func (r Rules) IsEmail(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		if !emailRegex.MatchString(v.String()) {
			return fmt.Errorf("\"isEmail\", %v value must be an email address", fieldName)
		}

		return nil
	default:
		return fmt.Errorf("\"isEmail\", %v value must be string", fieldName)
	}
}

// IsLatitude to check values is geo latitude or not
func (r Rules) IsLatitude(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		if !latitudeRegex.MatchString(fmt.Sprintf("%v", v)) {
			return fmt.Errorf("\"isLatitude\", %v value must be the geolocation Latitude", fieldName)
		}

		return nil
	case reflect.String:
		if !latitudeRegex.MatchString(v.String()) {
			return fmt.Errorf("\"isLatitude\", %v value must be the geolocation Latitude", fieldName)
		}

		return nil
	default:
		return fmt.Errorf("\"isLatitude\", %v value must be float or string", fieldName)
	}
}

// IsLongitude to check values is geo longitude or not
func (r Rules) IsLongitude(fieldName string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		if !longitudeRegex.MatchString(fmt.Sprintf("%v", v)) {
			return fmt.Errorf("\"isLongitude\", %v value must be the geolocation Longitude", fieldName)
		}

		return nil
	case reflect.String:
		if !longitudeRegex.MatchString(v.String()) {
			return fmt.Errorf("\"isLongitude\", %v value must be the geolocation Longitude", fieldName)
		}

		return nil
	default:
		return fmt.Errorf("\"isLongitude\", %v value must be float or string", fieldName)
	}
}

// IsIn to check value in array strings
func (r Rules) IsIn(fieldName string, v reflect.Value, in string) error {
	switch v.Kind() {
	case reflect.String:
		if !strings.Contains(in, v.String()) {
			return fmt.Errorf("\"isIn\", %v value is not allowed", fieldName)
		}

		return nil
	default:
		return fmt.Errorf("\"isIn\", %v value must be a string", fieldName)
	}
}

// IsNotIn to check value not in array strings
func (r Rules) IsNotIn(fieldName string, v reflect.Value, in string) error {
	switch v.Kind() {
	case reflect.String:
		if strings.Contains(in, v.String()) {
			return fmt.Errorf("\"isNotIn\", %v value is not allowed", fieldName)
		}

		return nil
	default:
		return fmt.Errorf("\"isNotIn\", %v value must be a string", fieldName)
	}
}

// LengthBetween to check string values that have char length minimal and maximal
func (r Rules) LengthBetween(fieldName string, v reflect.Value, m string) error {
	switch v.Kind() {
	case reflect.String:
		strLen := len(v.String())
		in := strings.Split(m, ",")
		min, _ := strconv.Atoi(in[0])
		max, _ := strconv.Atoi(in[1])

		if strLen > max {
			return fmt.Errorf("\"lengthBetween\", %v value length must be less than %v", fieldName, max)
		}

		if strLen < min {
			return fmt.Errorf("\"lengthBetween\", %v value length must be more than %v", fieldName, min)
		}

		return nil
	default:
		return fmt.Errorf("\"lengthBetween\", %v value must be string", fieldName)
	}
}

// MinLength to check string values that have char length minimal
func (r Rules) MinLength(fieldName string, v reflect.Value, m string) error {
	switch v.Kind() {
	case reflect.String:
		min, _ := strconv.Atoi(m)
		if len(v.String()) < min {
			return fmt.Errorf("\"minLength\", %v value length must be more than %v", fieldName, min)
		}

		return nil
	default:
		return fmt.Errorf("\"minLength\", %v value must be string", fieldName)
	}
}

// MaxLength to check string values that have char length maximal
func (r Rules) MaxLength(fieldName string, v reflect.Value, m string) error {
	switch v.Kind() {
	case reflect.String:
		max, _ := strconv.Atoi(m)
		if len(v.String()) > max {
			return fmt.Errorf("\"maxLength\", %v value length should not be more than %v", fieldName, max)
		}

		return nil
	default:
		return fmt.Errorf("\"maxLength\", %v value must be string", fieldName)
	}
}

// StartsWith to check if string has a specific prefix
func (r Rules) StartsWith(fieldName string, v reflect.Value, m string) error {
	switch v.Kind() {
	case reflect.String:
		if !strings.HasPrefix(v.String(), m) {
			return fmt.Errorf("\"startsWith\", %v value must be start with %v", fieldName, m)
		}

		return nil
	default:
		return fmt.Errorf("\"startsWith\", %v value must be string", fieldName)
	}
}

// EndWith to check if the end string has a specific prefix
func (r Rules) EndWith(fieldName string, v reflect.Value, m string) error {
	switch v.Kind() {
	case reflect.String:
		if !strings.HasSuffix(v.String(), m) {
			return fmt.Errorf("\"endWith\", %v value must be end with %v", fieldName, m)
		}

		return nil
	default:
		return fmt.Errorf("\"endWith\", %v value must be string", fieldName)
	}
}
