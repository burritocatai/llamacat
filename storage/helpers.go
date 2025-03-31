package storage

import (
	"fmt"
	"strings"
	"time"
)

func ReplaceStringWithDate(stringWithDate string) string {
	currentTime := time.Now()

	// Formats
	year := fmt.Sprintf("%04d", currentTime.Year())
	monthDigits := fmt.Sprintf("%02d", int(currentTime.Month()))
	monthAbbrev := currentTime.Format("Jan")
	monthFull := currentTime.Format("January")
	dayFull := currentTime.Format("Monday")
	dayAbbrev := currentTime.Format("Mon")
	dayDigits := currentTime.Format("02")

	// Replacements
	replaced := strings.ReplaceAll(stringWithDate, "yyyy", year)
	replaced = strings.ReplaceAll(replaced, "MMMM", monthFull)
	replaced = strings.ReplaceAll(replaced, "MMM", monthAbbrev)
	replaced = strings.ReplaceAll(replaced, "MM", monthDigits)
	replaced = strings.ReplaceAll(replaced, "EEEE", dayFull)
	replaced = strings.ReplaceAll(replaced, "EEE", dayAbbrev)
	replaced = strings.ReplaceAll(replaced, "dd", dayDigits)

	return replaced

}
