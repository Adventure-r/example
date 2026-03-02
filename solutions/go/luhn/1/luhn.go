package luhn

import (
	"fmt"
	"strings"
)

func Valid(id string) bool {
	//fmt.Println(id, "=\n-------------")
	id = strings.TrimSpace(id)
	var startIndex int
	if len(id) == 1 {
		return false
	}
	if len(id)%2 == 0 {
		startIndex = 0
	} else {
		startIndex = 1
	}
	var sum int
	var check string
	for i := startIndex; i <= len(id)-1; i += 2 {
		multBy2 := id[i] * 2
		if multBy2 > 9 {
			multBy2 -= 9
		}
		sum += int(multBy2)
		check += fmt.Sprint(multBy2)
		if i <= len(id)+1 {
			sum += int(id[i+1])
			check += fmt.Sprint(sum)
		}
	}
	//fmt.Println(check, "\nEnd")
	return sum%10 == 0
}
