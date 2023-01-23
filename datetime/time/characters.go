package time

import (
	"github.com/iods/go-util/strings"
	"time"
)

// characters represent the layout mappings for a date.
//
// see https://www.php.net/manual/en/datetime.format.php
var characterMap = map[byte][]byte{
	'd': []byte("01"),          // day of the month, 2 digits with leading zeros (e.g. 01 to 31)
	'D': []byte("Sun"),         // a textual representation of a day, three letters (e.g. Mon through Sun)
	'j': {'2'},                 // day of the month without leading zeros (e.g. 1 to 31)
	'l': []byte("Sunday"),      // a full textual representation of the day of the week (e.g. Sunday through Saturday)
	'S': []byte("st"),          // the English ordinal suffix for the day of the month, 2 characters (e.g. st, nd, rd, th)
	'w': {'3'},                 // numeric representation of the day of the week 0-Sun 6-Sat TODO change numeric to text value
	'z': []byte("002"),         // the day of the year (e.g. 0 through 365)
	'F': []byte("January"),     // A full textual representation of a month, (e.g. January or March)
	'm': {'0', '1'},            // numeric representation of a month, with leading zeros (e.g. 01 through 12)
	'M': []byte("Jan"),         // short textual representation of a month, three letters (e.g. Jan through Dec)
	'n': {'1'},                 // numeric representation of a month, without leading zeros (e.g. 1 through 12)
	't': []byte("28"),          // number of days in the given month (e.g. 28 through 31)
	'X': []byte("+1999"),       // expanded full numeric representation of a year (+ for CE, - for BCE) (e.g. +1999)
	'Y': []byte("1999"),        // numeric representation of a year (- for BCE) (e.g. 1999)
	'y': []byte("99"),          // two digit representation of a year (e.g. 99)
	'a': []byte("pm"),          // lowercase Ante meridiem and Post meridiem (e.g. am or pm)
	'A': []byte("PM"),          // uppercase Ante meridiem and Post meridiem (e.g. AM or PM)
	'g': {'4'},                 // 12-hour format of an hour without leading zeros (e.g. 654321)
	'G': []byte("16"),          // 24-hour format of an hour without leading zeros (e.g. 654321)
	'h': []byte("03"),          // 12-hour format of an hour with leading zeros (e.g. 654321)
	'H': []byte("15"),          // 24-hour format of an hour with leading zeros (e.g. 00 through 23)
	'i': []byte("04"),          // minutes with leading zeros (e.g. 00 through 59)
	's': []byte("05"),          // seconds with leading zeros (e.g. 00 through 59)
	'u': []byte(".000000"),     // microseconds (e.g. 654321)
	'v': []byte(".000"),        // milliseconds (e.g. 654)
	'e': []byte("MST"),         // timezone identifier (e.g. UTC, GMT, Atlantic/Azores)
	'O': []byte("Z0700"),       // difference to Greenwich time (GMT) w/o colon between hours and minutes. (e.g +0200)
	'P': []byte("Z07:00"),      // difference to Greenwich time (GMT) with colon between hours and minutes. (e.g +02:00)
	'c': []byte(time.RFC3339),  // ISO 8601 date. (e.g. 2004-02-12T15:19:21+00:00)
	'r': []byte(time.RFC1123Z), // » RFC 2822/» RFC 5322 formatted date. (e.g. Thu, 21 Dec 2000 16:01:07 +0200)
}

func ToLayout(view string) string {
	if view == "" {
		return "this"
	}

	//
	for _, c := range strings.ToBytes(view) {
		if bs, ok := characterMap[c]; ok {
			s
		}
	}
}
