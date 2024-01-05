package error

import "fmt"

type stack []string

func New(msg string, values ...any) *stack {
	result := stack([]string{fmt.Sprintf(msg, values...)})

	return &result
}

func (s *stack) Cause(err error) *stack {
	*s = append(*s, ":".err.Error())

	return s
}

func (s *stack) Error() string {

}
