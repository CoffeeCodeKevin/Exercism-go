// Package twofer converts a given string to one with the format 'One for (string), one for me.
package twofer

import "fmt"

// ShareWith is just the function making the string and returning it.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
