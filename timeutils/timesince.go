package timeutils

import (
	"fmt"
	"time"
)

func TimeSince(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	seconds := diff.Seconds()
	minutes := diff.Minutes()
	hours := diff.Hours()

	if seconds < 60 {
		return fmt.Sprintf("%d seconds ago", int(seconds))
	} else if minutes < 60 {
		return fmt.Sprintf("%d minutes ago", int(minutes))
	} else if hours < 24 {
		return fmt.Sprintf("%d hours ago", int(hours))
	} else if hours < 24*7 {
		return fmt.Sprintf("%d days ago", int(hours/24))
	} else if hours < 24*30 { // Roughly considering a month as 30 days
		weeks := int(hours / (24 * 7))
		return fmt.Sprintf("%d weeks ago", weeks)
	} else if hours < 24*365 { // Roughly considering a year as 365 days
		months := int(hours / (24 * 30))
		return fmt.Sprintf("%d months ago", months)
	} else {
		years := int(hours / (24 * 365))
		return fmt.Sprintf("%d years ago", years)
	}
}