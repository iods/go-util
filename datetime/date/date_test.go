package date

import "testing"

func TestDate_Year(t *testing.T) {
	d := Date{year: 1988}
	if d.year != 1988 {
		t.Error("the year entered does not equal 1988.")
	}
}

func TestDate_SetYear(t *testing.T) {
	d := Date{}
	err := d.SetYear(1988)
	if err != nil && d.year < 1 {
		t.Error("0 is not a valid year")
	}
}

func TestDate_Month(t *testing.T) {
	d := Date{month: 04}
	if d.month != 04 {
		t.Error("the month entered was not what was returned")
	}
}

func TestDate_SetMonth(t *testing.T) {
	d := Date{}
	err := d.SetMonth(04)
	if err != nil && d.month < 1 || d.month > 12 {
		t.Error("the month set must be between 1 and 12")
	}
}

func TestDate_Day(t *testing.T) {
	d := Date{day: 25}
	if d.day != 25 {
		t.Error("the day entered does not equal 25")
	}
}

func TestDate_SetDay(t *testing.T) {
	d := Date{}
	err := d.SetDay(25)
	if err != nil && d.day < 1 || d.day > 31 {
		t.Error("the day set must be between 1 and 31")
	}
}
