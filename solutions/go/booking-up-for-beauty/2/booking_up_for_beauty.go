package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	v, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		fmt.Println(err)
	}
	return v
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	v, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		fmt.Println(err)
	}
	return v.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	v, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		fmt.Println(err)
	}
	hour := v.Hour()
	if hour >= 12 && hour <= 18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	myTime := Schedule(date)
	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.",
		myTime.Weekday(), myTime.Month(), myTime.Day(),
		myTime.Year(), myTime.Hour(), myTime.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	date := time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
	return date
}
