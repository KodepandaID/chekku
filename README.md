# Chekku
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/KodepandaID/chekku)
![GitHub](https://img.shields.io/github/license/KodepandaID/chekku)
![](https://github.com/KodepandaID/chekku/workflows/Go/badge.svg)

Chekku is Golang validation library

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Examples](#examples)
    - [Installation](#installation)
    - [Basic Usage](#basic-usage)
    - [Return Error](#return-error)
    - [Custom Error](#custom-error)
- [Available Rules](#available-rules)
    - [`required`](#required)
    - [`either`](#eitherrule1rule2)
    - [`isAlpha`](#isalpha)
    - [`isAlphanumeric`](#isalphanumeric)
    - [`isArray`](#isarray)
    - [`isBase64`](#isbase64)
    - [`isBool`](#isbool)
    - [`isDate`](#isdate)
    - [`isEmail`](#isemail)
    - [`isFloat`](#isfloat)
    - [`isIn`](#isIn)
    - [`isInt`](#isint)
    - [`isIPv4`](#isipv4)
    - [`isIPv6`](#isipv6)
    - [`isISBN`](#isisbn)
    - [`isLatitude`](#islatitude)
    - [`isLongitude`](#islongitude)
    - [`isNumber`](#isnumber)
    - [`isNumeric`](#isnumeric)
    - [`isUUID`](#isuuid)
    - [`lengthBetween:min,max`](#lengthbetweenminmax)
    - [`maxLength:max`](#maxlengthmax)
    - [`minLength:min`](#minlengthmin)
    - [`numberBetween:min,max`](#numberbetweenminmax)
    - [`maxNumber:max`](#maxnumbermax)
    - [`minNumber:min`](#minnumbermin)
    - [`notIn`](#notin)
    - [`requiredIf`](#requiredifanotherfieldvalue)
    - [`requiredUnless`](#requiredunlessanotherfieldvalue)
    - [`startsWith:str`](#startsWithstr)
    - [`endWith:str`](#endWithstr)
    - [`mimetype:value`](#mimetypevalue)
    - [`filesize:value`](#filesizevalue)
    - [`dimensions:value`](#dimensionsvalue)

## Examples

#### Installation
```bash
go get github.com/KodepandaID/chekku
```

#### Basic Usage
```go
package main

import "github.com/KodepandaID/chekku"

type Payload struct {
    Name string `chekku:"required"`
    Age string `chekku:"required|isNumber"`
}

inputs := Payload{
    Name: "",
    Age: "20",
}

e := chekku.Validate(inputs)
if e != nil {
    panic(e)
}
```

Use ```|``` as separators to use multiple validations at once.

#### Return Error

in version 1 the return error is converted to an array struct as below:
```go
type Error struct {
    Code string
    Detail string
}
```

#### Custom Error

If you want to set error message by yourself, you can use tag ```errors``` on your struct.

Example:

```go
type Error struct {
    Username string `chekku:"required|maxLength:10" errors:"required:username tidak boleh kosong|maxLength:Panjang username tidak boleh lebih dari 10 karakter"`
}

e := chekku.Validate(Error{
    Username: "",
})

if e == nil {
    panic(e)
}
```


## Available Rules

#### `required`

Value under this field should not be `nil` or an empty string (`""`).

* Invalid values: `nil`, `""`, `0`, `false`, `[]string{}`, `[]int{}`, `[]float{}`
* Valid values: `"not empty string"`, `0.1`, `[]string{""}`

#### `either:rule1,rule2`

Use this as an ```OR``` operator.

For example you may want to accept ipv4 or ipv6, you can use either rule like below:

Example:

```go
type IP struct {
    ip string `chekku:"either:isIPv4,isIPv6"`
}

validValues := []string{
    "1.2.3.4",  // Valid
    "::1", // Valid
    "not an IP Address" // Invalid
}
for _, v := range validValues {
    e := chekku.Validate(IP{
        ip: v,
    })
    if e != nil {
        panic(e)
    }
}
```

#### `isAlpha`

Value under this field must be the words Aa - Zz.

* Invalid values: `"123"`, `"hello123"`, etc.
* Valid values: `"hello"`, `"Hello"`, etc.

#### `isAlphanumeric`

Value under this field must be the words or numbers.

* Invalid values: `"hello@mail.com"`, `"Here!!!"`, `"~@#$%^&*()"`, `123`, etc.
* Valid values: `"123"`, `"hello"`, `"hello123"`, etc.

#### `isArray`

Value under this field must be an array.

* Invalid values: `""`, `10`, `0.5`, etc.
* Valid values: `[]string{}`, `map[string]string{}`, etc.

#### `isBase64`

Value under this field must be a base64.

* Invalid values: `"123"`, `0.5`, `"aGVsbG%8gd29ybGQ="`, etc.
* Valid values: `"aGVsbG8gd29ybGQ="`, etc.

#### `isBool`

Value under this field must be a boolean.

* Invalid values: `"true"`, `1`, etc.
* Valid values: `true`, `false`.

#### `isDate`

Value under this field must be a Date format.

* Invalid values: `"2020"`, `"01-2020-10"`, `"2020-01"`, etc.
* valid values: `"2020-01-02"`, `"2020-01-02 23:59:59"`, `"02/12/2020"`, `"2020-January-20"`, etc.

#### `isEmail`

Value under this field mus be an email address.

* Invalid values: `123`, `"hello"`, etc.
* Valid values: `"hello@kodepanda.com"`, etc.

#### `isFloat`

Value under this field must be a float.

* Invalid values: `1`, `"hello"`, etc.
* Valid values: `0.123`, etc.

#### `isIn`

Value under this field must be one of allowed values

Example:
```go

type Payload struct {
    text1 string `chekku:"isIn:yes,no"`
    text2 string `chekku:"isIn:yes,no"`
    text3 string `chekku:"isIn:yes,no"`
}

p := Payload{
    text1: "yes", // passes
    text2: "no", // passes
    text3: "maybe", // fail
}

e := chekku.Validate(p)
if e != nil {
    panic(e)
}

```

#### `isInt`

Value under this field must be an integer.

* Invalid values: `0.5`, `"123"`, etc.
* Valid values: `0`, `123`, etc.

#### `isIPv4`

Value under this field must be IP Address version 4.

* Invalid values: `123`, `"1.2.3.a"`, `"1.2.3.256"`, `"FEDC:BA98:7654:3210:FEDC:BA98:7654:3210"` etc.
* Valid values: `"0.1.2.3"`, `"255.255.255.255"`, etc.

#### `isIPv6`

Value under this field must be IP Address version 6.

* Invalid values: `"2001:db8::7"`, `"::1:2:3:4:5:6:7"`, etc.
* Valid values: `"2001:db8::7"`, `"::1:2:3:4:5:6:7"`, etc.

#### `isISBN`

Value under this field must be an ISBN 10 or ISBN 13.

* Invalid values: `"123"`, `"ISBX 979-1-23456-789-6"`, etc.
* Valid values: `"ISBN 979-1-23456-789-6"`, `"ISBN 1-23456-789-X"`, etc.

#### `isLatitude`

Value under this field must be a latitude geolocation.

* Invalid values: `"+90.0123"`, `"123"`, `"-123"`, etc.
* Valid values: `+90.0123`, `123`, `-123`, etc.

#### `isLongitude`

Value under this field must be a longitude geolocation.

* Invalid values: `"182.1234"`, `"-181.1234"`, `182.1234`, `-181.1234` etc.
* Valid values: `106.900447`, `-120.1234`, `+90.123`, etc.

#### `isNumber`

Value under this value must be a number.

* Invalid values: `"hello"`, `"123"`, etc.
* Valid values: `123`, `1.23`, etc.

#### `isNumeric`

Value under this field must be a numeric.

* Invalid values: `"hello"`, etc.
* Valid values: `"123"`, `"1.23"`, `123`, `1.23`, etc.

#### `isUUID`

Value under this field must be an UUID.

* Invalid values: `"hello"`, `"123-456-789"`, etc.
* Valid values: `"17b32300-f33c-11ea-adc1-0242ac120002"`, etc.

#### `lengthBetween:min,max`

Value under this field must be a string that has char length between `min` and `max`.

Example:
```go
type Text struct {
    text1 string `chekku:"lengthBetween:5,10"`
    text2 string `chekku:"lengthBetween:5,10"`
    text3 string `chekku:"lengthBetween:5,10"`
}

p := Text{
    text1: "hell", // fail
    text2: "hello world, how are you?", // fail
    text3: "hello" // passes
}

e := chekku.Validate(p)
if e != nil {
    panic(e)
}
```

#### `maxLength:max`

Value under this field must be a string that has char length lower or equals `max`.

Example:
```go
type Text struct {
    text1 string `chekku:"maxLength:5"`
    text2 string `chekku:"maxLength:5"`
}

p := Text{
    text1: "hello", // passes
    text2:"hello world" // fail
}
e := chekku.Validate(p)
if e != nil {
    panic(e)
}
```

#### `minLength:min`

Value under this field must be a string that has char length higher or equals `min`.

Example:
```go
type Text struct {
    text1 string `chekku:"minLength:5"`
    text2 string `chekku:"minLength:5"`
}

p := Text{
    text1: "hell", // fail
    text2:"hello world" // passes
}

e := chekku.Validate(p)
if e != nil {
    panic(e)
}
```

#### `numberBetween:min,max`

Value under this field must be a int or float that has value between `min` and `max`.

Example:
```go
type Number struct {
    number1 int `chekku:"numberBetween:5,10"`
    number2 int `chekku:"numberBetween:5,10"`
    number3 float64 `chekku:"numberBetween:5.1,5.3"`
}

p := Number{
    number1: 4, // fail
    number2: 11, // fail
    number3: 5.2 // passes
}

e := chekku.Validate(p)
if e != nil {
    panic(e)
}
```

#### `maxNumber:max`

Value under this field must be a int or float that has value lower or equals `max`.

Example:
```go
type Number struct {
    number1 int `chekku:"maxNumber:5"`
    number2 float64 `chekku:"maxNumber:5"`
}

p := Number{
    number1: 4, // passes
    number2: 5.1 // fail
}
e := chekku.Validate(p)
if e != nil {
    panic(e)
}
```

#### `minNumber:min`

Value under this field must be a int or float that has value higher or equals `min`.

Example:
```go
type Number struct {
    number1 int `chekku:"minNumber:5"`
    number2 int `chekku:"minNumber:5"`
}

p := Number{
    number1: 4, // fail
    number2: 6 // passes
}

e := chekku.Validate(p)
if e != nil {
    panic(e)
}
```

#### `requiredIf:anotherField,value`

Field within this rule will be required if given ```anotherField``` match ```value```.

Example Passes (text1 is optional because text0 value not match):
```go
type Text struct {
    text0 int
    text1 string `chekku:"requiredIf:text0,200"`
}

e := chekku.Validate(Text{
    text0: 201,
})
if e != nil {
    panic(e)
}
```

Example Fail (text1 is required because text0 value is a match):
```go
type Text struct {
    text0 int
    text1 string `chekku:"requiredIf:text0,200"`
}

e := chekku.Validate(Text{
    text0: 200,
})
if e != nil {
    panic(e)
}
```

#### `requiredUnless:anotherField,value`

Field within this rule will be required if given ```anotherField``` doesn't match with ```value```.

Example Passes (text1 is optional because text0 value is a match):
```go
type Text struct {
    text0 int
    text1 string `chekku:"requiredUnless:text0,200"`
}

e := chekku.Validate(Text{
    text0: 200,
})
if e != nil {
    panic(e)
}
```

Example Fail (text1 is required because text0 value doesn't match):
```go
type Text struct {
    text0 int
    text1 string `chekku:"requiredUnless:text0,200"`
}

e := chekku.Validate(Text{
    text0: 201,
})
if e != nil {
    panic(e)
}
```

#### `startsWith:str`

Value under this field must be a string that starts with prefix `str`.

Example:
```go
type Text struct {
    text1 string `chekku:"startsWith:foo"`
    text2 string `chekku:"startsWith:foo"`
    text3 string `chekku:"startsWith:foo"`
}

p := Text{
    text1: "", // fail
    text2: "barfoo", // fail
    text3: "foobar", // passes
}

e := chekku.Validate(p)
if e != nil {
    panic(e)
}
```

#### `endWith:str`

Value under this field must be a string that end with prefix `str`.

Example:
```go
type Text struct {
    text1 string `chekku:"endWith:foo"`
    text2 string `chekku:"endWith:foo"`
    text3 string `chekku:"endWith:foo"`
}

p := Text{
    text1: "", // fail
    text2: "barfoo", // passes
    text3: "foobar", // fail
}

e := chekku.Validate(p)
if e != nil {
    panic(e)
}
```

#### `mimetype:value`

The file under validation must match one of the given MIME types:

Example:
```go
type Request struct {
    File *multipart.FileHeader `chekku:"mimetype:image/png"`
}

// fh is a variable to pass multipart.FileHeader you can use another name whatever you want.
e = chekku.Validate(Request{
    File: fh,
})
if e != nil {
    panic(e)
}
```

A full listing of MIME types and their corresponding extensions may be found at the following location: [https://svn.apache.org/repos/asf/httpd/httpd/trunk/docs/conf/mime.types](https://svn.apache.org/repos/asf/httpd/httpd/trunk/docs/conf/mime.types)

#### `filesize:value`

The file under validation cannot be exceeded max size in ```Kilobytes```.

Example:
```go
type Request struct {
    File *multipart.FileHeader `chekku:"filesize:512"`
}

// fh is a variable to pass multipart.FileHeader you can use another name whatever you want.
e = chekku.Validate(Request{
    File: fh,
})
if e != nil {
    panic(e)
}
```

#### `dimensions:value`

The file under validation must be an image meeting the dimension constraints as specified by the rule's parameters.

Available constraints are: min_width, max_width, min_height, max_height, width, height, ratio.

A ratio constraint should be represented as width divided by height. This can be specified either by a statement like 1/1, 4/3, 16/9 or etc.

Example:
```go
type Request struct {
    File *multipart.FileHeader `chekku:"dimensions:max_width=300,min_width=150,max_height=300,min_height=150,ratio=1/1"`
}

// fh is a variable to pass multipart.FileHeader you can use another name whatever you want.
e = chekku.Validate(Request{
    File: fh,
})
if e != nil {
    panic(e)
}
```