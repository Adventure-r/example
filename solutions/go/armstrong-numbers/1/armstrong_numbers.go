package armstrongnumbers

func IsNumber(n int) bool {
	count := 0
	digits := make([]int, 0, 10)
	buffer := n
	for buffer > 10 {
		digits = append(digits, buffer%10)
		buffer = buffer / 10
		count++
	}
	digits = append(digits, buffer%10)
	count++

	var sum int
	for _, digit := range digits {
		sum += power(digit, count)
	}
	return sum == n
}

func power(num, count int) int {
	result := 1
	for range count {
		result *= num
	}
	return result
}
