package error

import (
	"reflect"

	"gopkg.in/errgo.v1"
)

// return true if the cause of the error is the same type as X

// return error stack

func IsOfType(err error, match interface{}) bool {
	errType := reflect.TypeOf(match)
	errCause :=
}

func IsRootCause(err error, errmatch interface{}) bool {

}