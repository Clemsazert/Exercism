package booking

import (
	"fmt"
	"time"
)

// Mon Jan 2 15:04:05 -0700 MST 2006

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return parsedDate
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return parsedDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return parsedDate.Hour() >= 12 && parsedDate.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("You have an appointment on %s.", parsedDate.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return Schedule("9/15/2022 00:00:00")
}
