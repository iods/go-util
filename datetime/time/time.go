package time

import "time"

// TimeAttributes represents attributes in the struct for `time.Date`
type TimeAttributes struct {
	Year       int
	Month      int
	Day        int
	Hour       int
	Minute     int
	Second     int
	Nanosecond int
	Location   time.Location
}
