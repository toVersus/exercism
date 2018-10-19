package accumulate

func Accumulate(col []string, converter func(string) string) []string {
	if len(col) == 0 {
		return []string{}
	}

	var c string
	for i := 0; i < len(col); i++ {
		c, col = col[0], col[1:]
		col = append(col, converter(c))
	}

	return col
}
