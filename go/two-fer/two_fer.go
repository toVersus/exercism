// Package twofer is short for two for one,
// and provides helper function to inform something is shared by two.
package twofer

import "fmt"

// ShareWith shares something with someone in given name.
func ShareWith(name string) string {
	if len(name) == 0 {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
