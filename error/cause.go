package error

import (
	"reflect"
)

var errorType = reflect.TypeOf((*error)(nil)).Elem()

type Error struct {
	StatusCode int         `json:"-"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Details    interface{} `json:"details,omitempty"`
	trace      string
}

// return true if the cause of the error is the same type as X

// return error stack

//func IsOfType(err error, match interface{}) bool {
//	errType := reflect.TypeOf(match)
//	errCause :=
//}
//
//func IsRootCause(err error, errmatch interface{}) bool {
//
//}
