package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    // layout := "Mon, 01/02/2006, 15:04"
    layout := "1/2/2006 15:04:05"
    t, _ := time.Parse(layout,date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
    d, _ := time.Parse(layout,date)
	n := time.Now()
	return n.After(d)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	d, _ := time.Parse(layout, date)
	app := false
	hour := d.Hour()
	if hour >= 12 && hour < 18{
		app = true
	}
	return app
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}
