package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n == 0 {
		return 0, errors.New("0 can't go to 1")
	}
	counter := 0
	for ; n != 1; counter++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
	}
	return counter, nil
}
