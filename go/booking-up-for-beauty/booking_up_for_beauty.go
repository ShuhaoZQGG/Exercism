package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// Parsing layout matches the input format
	parseLayout := "1/2/2006 15:04:05"

	// Output layout matches the desired output format
	outputLayout := "2006-01-02 15:04:05 -0700 MST"

	t, err := time.Parse(parseLayout, date)
	if err != nil {
		panic(err)
	}
	// Return the formatted date string
	t, err = time.Parse(outputLayout, t.Format(outputLayout))
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	parseLayout := ("January 2, 2006 15:04:05")
	t, _ := time.Parse(parseLayout, date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parseLayout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(parseLayout, date)
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	prefix := "You have an appointment on "
	parseLayout := "1/2/2006 15:04:05"
	t, _ := time.Parse(parseLayout, date)

	outputLayout := "Monday, January 2, 2006, at 15:04."
	formattedTime := t.Format(outputLayout)
	return prefix + formattedTime
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
