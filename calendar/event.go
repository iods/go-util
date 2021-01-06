package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

/*
Title Retrieves the title value from the Event struct. */
func (e *Event) Title() string {
	return e.title
}

/*
SetTitle Sets the value of title from the Event struct. */
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}