package diffsquares

func SquareOfSum(n int) int {
	var sum int
	for v := range n + 1 {
		sum += v
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	var sum int
	for v := range n + 1 {
		sum += v * v
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
