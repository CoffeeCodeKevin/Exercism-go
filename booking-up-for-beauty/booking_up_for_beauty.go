package booking

import "time"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const shortForm = "1/2/2006 15:04:05"
	value, _ := time.Parse(shortForm, date)
	return value
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	const dateFormat = "January 2, 2006 15:04:05"
	today := time.Now()
	convertedDate, _ := time.Parse(dateFormat, date)
	return convertedDate.Before(today)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	const dateFormat = "Monday, January 2, 2006 15:04:05"
	convertedDate, _ := time.Parse(dateFormat, date)
	return (convertedDate.Hour() >= 12 && convertedDate.Hour() < 18)
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	const dateFormat = "1/2/2006 15:04:05"
	convertedDate, _ := time.Parse(dateFormat, date)
	const newFormat = "You have an appointment on Monday, January 2, 2006, at 15:04."
	description := convertedDate.Format(newFormat)
	return description
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	const dateFormat = "January 2th, 2006"
	convertedDate, _ := time.Parse(dateFormat, "September 15th, 2012")
	currentYear := time.Now().Year()
	anniversary := convertedDate.AddDate(currentYear-convertedDate.Year(), 0, 0)
	return anniversary
}
