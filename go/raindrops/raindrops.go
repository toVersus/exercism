package raindrops

import (
	"bytes"
	"fmt"
)

// converter is the pre-defined rules for converting.
var converter = map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}

// Convert converts the input number to appropriate string
// by following the pre-defined rules above.
func Convert(n int) string {
	var factors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}

	var out bytes.Buffer
	var matched bool
	for _, f := range factors {
		if val, ok := converter[f]; ok {
			out.WriteString(val)
			matched = true
		}
	}

	if !matched {
		out.WriteString(fmt.Sprint(n))
	}

	return out.String()
}
