package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/02/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	d, _ := time.Parse("January 2, 2006 15:04:05", date)
	return d.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	d, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	if d.Hour() >= 12 && d.Hour() < 18 {
		return true
	}
	return false
}


func Description(date string) string {
	d, _ := time.Parse("1/2/2006 15:04:05", date)
	return fmt.Sprintf("You have an appointment on %s", d.Format("Monday, January 2, 2006, at 15:04."))
}

func AnniversaryDate() time.Time {
	t, _ := time.Parse("2006-01-2", fmt.Sprintf("%d-09-15", time.Now().Year()))
	return t
}
