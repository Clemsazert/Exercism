package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	recommendation := strings.IndexRune(log, '❗')
	search := strings.IndexRune(log, '🔍')
	weather := strings.IndexRune(log, '☀')
	if recommendation == -1 && search == -1 && weather == -1 {
		return "default"
	}
	if recommendation == -1 {
		recommendation = utf8.RuneCountInString(log) + 1
	}
	if search == -1 {
		search = utf8.RuneCountInString(log) + 1
	}
	if weather == -1 {
		weather = utf8.RuneCountInString(log) + 1
	}
	if recommendation < search && recommendation < weather {
		return "recommendation"
	}
	if search < recommendation && search < weather {
		return "search"
	}
	if weather < recommendation && weather < search {
		return "weather"
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
