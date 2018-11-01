package strain

type Ints []int
type Lists [][]int
type Strings []string

func (is Ints) Keep(pred func(int) bool) Ints {
	var kept []int
	for _, i := range is {
		if pred(i) {
			kept = append(kept, i)
		}
	}
	return kept
}

func (is Ints) Discard(pred func(int) bool) Ints {
	return is.Keep(func(x int) bool { return !pred(x) })
}

func (ls Lists) Keep(pred func([]int) bool) Lists {
	var kept Lists
	for _, l := range ls {
		if pred(l) {
			kept = append(kept, l)
		}
	}
	return kept
}

func (ss Strings) Keep(pred func(string) bool) Strings {
	var kept Strings
	for _, s := range ss {
		if pred(s) {
			kept = append(kept, s)
		}
	}
	return kept
}
