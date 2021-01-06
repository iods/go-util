package calendar

import (
	"testing"
	"unicode/utf8"
)

func TestEvent_Title(t *testing.T) {
	e := Event{title: "The Event Title"}
	if e.title != "The Event Title" {
		t.Error("the title returned does not equal the title set")
	}
}

func TestEvent_SetTitle(t *testing.T) {
	e := Event{}
	err := e.SetTitle("The Event Title")
	if err != nil && utf8.RuneCountInString(e.title) > 30 {
		t.Error("invalid title length returned")
	}
}