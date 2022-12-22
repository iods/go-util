package errorutil

import (
	"gopkg.in/errgo.v1"
)

func bottom(err error) error {
	for {
		e, ok := err.(ErrorContext)
		if ok {
			err = e.err
		}

		errgoErr, ok := err.(*errgo.Err)
		if !ok {
			return err
		}

		if errgoErr.Underlying() == nil {
			return err
		}

		err = errgoErr.Underlying()
	}
}
