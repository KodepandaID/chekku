// Imported from go-playground/validator

package rules

import "regexp"

const (
	alphaRegexString               = "^[a-zA-Z]+$"
	alphaNumericRegexString        = "^[a-zA-Z0-9]+$"
	alphaUnicodeRegexString        = "^[\\p{L}]+$"
	alphaUnicodeNumericRegexString = "^[\\p{L}\\p{N}]+$"
	numericRegexString             = "^[-+]?[0-9]+(?:\\.[0-9]+)?$"
	numberRegexString              = "^[0-9]+$"
	emailRegexString               = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	base64RegexString              = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	base64URLRegexString           = "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})$"
	iSBNRegexString                = "ISBN[\x20\x2D]{0,1}(?:13[\x20\x2D]{0,1}){0,1}"
	uUID3RegexString               = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	uUID4RegexString               = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUID5RegexString               = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUIDRegexString                = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	latitudeRegexString            = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	longitudeRegexString           = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
)

var (
	alphaRegex               = regexp.MustCompile(alphaRegexString)
	alphaNumericRegex        = regexp.MustCompile(alphaNumericRegexString)
	alphaUnicodeRegex        = regexp.MustCompile(alphaUnicodeRegexString)
	alphaUnicodeNumericRegex = regexp.MustCompile(alphaUnicodeNumericRegexString)
	numericRegex             = regexp.MustCompile(numericRegexString)
	numberRegex              = regexp.MustCompile(numberRegexString)
	emailRegex               = regexp.MustCompile(emailRegexString)
	base64Regex              = regexp.MustCompile(base64RegexString)
	base64URLRegex           = regexp.MustCompile(base64URLRegexString)
	iSBNRegex                = regexp.MustCompile(iSBNRegexString)
	uUID3Regex               = regexp.MustCompile(uUID3RegexString)
	uUID4Regex               = regexp.MustCompile(uUID4RegexString)
	uUID5Regex               = regexp.MustCompile(uUID5RegexString)
	uUIDRegex                = regexp.MustCompile(uUIDRegexString)
	latitudeRegex            = regexp.MustCompile(latitudeRegexString)
	longitudeRegex           = regexp.MustCompile(longitudeRegexString)
)
