package string

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"unsafe"
)

// Float64ToString declares a function type for floats to strings.
type Float64ToString func(float64) string

// Int64ToString declares a function type for integers to strings.
type Int64ToString func(int64) string

var (
	ErrorMessageInvalidLayout    = errors.New("invalid layout, try again")
	ErrorMessageInvalidParameter = errors.New("invalid input param, try again")
)

var (
	toSnake = regexp.MustCompile("[A-Z][a-z]")
	toCamel = map[string]*regexp.Regexp{
		" ": regexp.MustCompile(" +[a-zA-Z]"),
		"-": regexp.MustCompile("-+[a-zA-Z]"),
		"_": regexp.MustCompile("_+[a-zA-Z]"),
	}
)

// Quote returns an alias of the default strconv.Quote() function.
func Quote(s string) string {
	return strconv.Quote(s)
}

// Unquote removes starting and ending quotes from anything containing "" or ”.
//
// strconv.Unquote() is not able to unquote single-quoted strings.
func Unquote(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}

	start, end := s[0], s[l-1]

	var remove bool
	if start == '"' && end == '"' {
		remove = true
	}
	if start == '\'' && end == '\'' {
		remove = true
	}
	if remove {
		s = s[1 : l-1]
	}

	return s
}

// String converts the provided value to a string.
func String(v any) (string, error) {
	return FormatToString(v, true)
}

// ToString will return any value as a string.
func ToString(v any) (string, error) {
	return FormatToString(v, true)
}

// ToStringOrErr converts any value to a string or returns an error on failure.
func ToStringOrErr(v any) (string, error) {
	return FormatToString(v, true)
}

// QuietToString converts any value to a string and ignores errors on failure.
func QuietToString(v any) string {
	val, _ := FormatToString(v, false)
	return val
}

// MustToString converts any value to a string and panics on error.
func MustToString(v any) string {
	val, _ := FormatToString(v, false)
	return val
}

// FormatToString takes any input and returns it as a string. Offers custom error output too.
func FormatToString(v any, errorOut bool) (s string, err error) {
	if v == nil {
		return
	}

	switch value := v.(type) {
	case int:
		s = strconv.Itoa(value)
	case int8:
		s = strconv.Itoa(int(value))
	case int16:
		s = strconv.Itoa(int(value))
	case int32:
		s = strconv.Itoa(int(value))
	case int64:
		s = strconv.FormatInt(value, 10)
	case uint:
		s = strconv.FormatUint(uint64(value), 10)
	case uint8:
		s = strconv.FormatUint(uint64(value), 10)
	case uint16:
		s = strconv.FormatUint(uint64(value), 10)
	case uint32:
		s = strconv.FormatUint(uint64(value), 10)
	case uint64:
		s = strconv.FormatUint(value, 10)
	case float32:
		s = strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		s = strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		s = strconv.FormatBool(value)
	case string:
		s = value
	case []byte:
		s = string(value)
	case json.Number:
		s = value.String()
	default:
		if errorOut {
			// need to do something on error out
		} else {
			s = fmt.Sprint(value)
		}
	}

	return
}

/*
Convert byte values to a string and strings back to bytes.
See: https://github.com/valyala/fastjson/blob/master/util.go
*/

// ByteToString converts bytes to a string.
func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// ByteToStr converts bytes to a string, an alias of ByteToString()
func ByteToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StrToBytes converts a string to bytes.
func StrToBytes(s string) (b []byte) {
	header := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	sh.Data = header.Data
	sh.Len = header.Len
	sh.Cap = header.Len

	return b
}
