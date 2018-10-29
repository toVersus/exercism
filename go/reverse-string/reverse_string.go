package reverse

import (
	"strings"
)

// String reverses input string.
func String(input string) string {
	t := strings.Split(input, "")
	l := len(t)
	for i := 0; i < l/2; i++ {
		t[i], t[l-i-1] = t[l-i-1], t[i]
	}
	return strings.Join(t, "")
}
