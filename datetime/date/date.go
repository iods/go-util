package date

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

// Year Getter method to return the year value of the Date struct.
func (d *Date) Year() int {
	return d.year
}

// Month Getter method to return the month value of the Date struct.
func (d *Date) Month() int {
	return d.month
}

// Day Getter method to return the day value of the Date struct.
func (d *Date) Day() int {
	return d.day
}

// SetYear Sets the year value of the Date struct.
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

// SetMonth Sets the month value of the Date struct.
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

// SetDay Sets the day value of the Date struct.
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
