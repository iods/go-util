package strutil

import "strings"

type Person struct {
	Name       string
	FirstName  string
	MiddleName string
	LastName   string
	Suffix     string

	Suffixes string
	Prefixes string
}

/*
see https://github.com/grokify/mogo/blob/master/strconv/humannameparser/humannameparser.go
*/

func ParsePerson(name string) (*Person, error) {
	p := &Person{}
	parts := strings.Split(name, " ")
	if len(parts) == 2 {
		p.FirstName = parts[0]
		p.LastName = parts[1]
	} else if len(parts) == 3 {
		p.FirstName = parts[0]
		p.MiddleName = parts[1]
		p.LastName = parts[2]
	}
	return p, nil
}
