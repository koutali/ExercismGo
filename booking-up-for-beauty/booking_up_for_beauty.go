package booking

import (
	"fmt"
	"os"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	const longForm = "7/13/2020 20:32:00"
	t, error := time.Parse(longForm, date)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	return t.UTC()
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	const longForm = "October 3 2019 20:32:00"
	givenDate, error := time.Parse(longForm, date)

	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	return time.Now().After(givenDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	const longForm = "October 3, 2019 20:32:00"
	givenDate, _ := time.Parse(longForm, date)

	hour := givenDate.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	const longForm = "7/13/2020 20:32:00"
	givenDate, error := time.Parse(longForm, date)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	givenDate.Format("October 3, 2019 20:32:00")

	appointment := "You have an appointment on " + givenDate.String()
	return appointment
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	anniversary := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	return anniversary
}
