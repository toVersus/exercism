package diffsquares

// SquareOfSum computes the square of the sum of the sequence
// of natural numbers.
func SquareOfSum(seq int) int {
	var sum int
	for ; seq > 0; seq-- {
		sum += seq
	}
	return sum * sum
}

// SumOfSquares computes the sum of the squares of the sequence
// of natural numbers.
func SumOfSquares(seq int) int {
	var total int
	for ; seq > 0; seq-- {
		total += seq * seq
	}
	return total
}

// Difference finds the difference between the square of the sum
// and the sum of the squares of the first N natural numbers.
func Difference(seq int) int {
	return SquareOfSum(seq) - SumOfSquares(seq)
}
