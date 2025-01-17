package main

import (
	"fmt"
	"strings"
)

func FormatDuration(seconds int64) string {
	if seconds == 0 {
		return "now"
	}

	const (
		secondsInMinute = 60
		secondsInHour   = 60 * secondsInMinute
		secondsInDay    = 24 * secondsInHour
		secondsInYear   = 365 * secondsInDay
	)

	years := seconds / secondsInYear
	seconds %= secondsInYear

	days := seconds / secondsInDay
	seconds %= secondsInDay

	hours := seconds / secondsInHour
	seconds %= secondsInHour

	minutes := seconds / secondsInMinute
	seconds %= secondsInMinute

	components := []string{}
	if years > 0 {
		components = append(components, fmt.Sprintf("%d year%s", years, pluralize(years)))
	}
	if days > 0 {
		components = append(components, fmt.Sprintf("%d day%s", days, pluralize(days)))
	}
	if hours > 0 {
		components = append(components, fmt.Sprintf("%d hour%s", hours, pluralize(hours)))
	}
	if minutes > 0 {
		components = append(components, fmt.Sprintf("%d minute%s", minutes, pluralize(minutes)))
	}
	if seconds > 0 {
		components = append(components, fmt.Sprintf("%d second%s", seconds, pluralize(seconds)))
	}

	if len(components) > 1 {
		last := components[len(components)-1]
		components = components[:len(components)-1]
		return strings.Join(components, ", ") + " and " + last
	}
	return components[0]
}

func pluralize(value int64) string {
	if value > 1 {
		return "s"
	}
	return ""
}

func main() {
	fmt.Println(FormatDuration(62))
	fmt.Println(FormatDuration(3662))
	fmt.Println(FormatDuration(31536000))
}
