package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	appEmitter := "default"
	for _, val := range log {
		switch fmt.Sprintf("%c", val) {
		case "❗":
			appEmitter = "recommendation"
		case "🔍":
			appEmitter = "search"
		case "☀":
			appEmitter = "weather"
		}
		if appEmitter != "default" {
			break
		}
	}
	return appEmitter
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newRuneStr := fmt.Sprintf("%c", newRune)
	oldRuneStr := fmt.Sprintf("%c", oldRune)
	newString := strings.ReplaceAll(log, oldRuneStr, newRuneStr)
	return newString
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return (utf8.RuneCountInString(log) <= limit)
}
