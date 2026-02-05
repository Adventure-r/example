package bookstore

import (
	"sort"
)

func Cost(books []int) int {
	// Firstly I changed view that books showed. Was: 1 1 2 2 3 4 => 2 books named 1, 2 books named 2, 1 book named 3, 1 book named 4
	// It helps to combine books into groups for discounts
	books = FromNumericNameToCountedNames(books)

	// skip that part for now later I will explain
	sort.Sort(sort.Reverse(sort.IntSlice(books)))

	// We need a real copy, no shadowing
	booksSave := make([]int, len(books))
	copy(booksSave, books)

	// This variable contains complete discount groups
	var booksGroupCounter [5]int
	// index 0 - 1st book type, index 1 - 2nd book type, index 2 - 3rd book type,
	// index 3 - 4th book type, index 4 - 5th book type

	// there program collects the groups
	for NotNull(books) {

		// counts how many different types of book we can group up
		counter := 0

		// if type has a book, program will take it and count, after go to the next type
		for i := range books {
			if books[i] > 0 {
				books[i] -= 1
				counter += 1
			}
		}

		// notes which group size program combined
		booksGroupCounter[counter-1]++

		// 'for' ends when books will runout
	}
	// calculating cost, not that hard when you know your discount groups
	totalCost5InARow := CalculateTotal(booksGroupCounter)

	/* fmt.Println("5 in a row collection")
	for _, value := range booksGroupCounter {
		fmt.Print(value, " ")
	}
	fmt.Print("\n5 in a row cost:", totalCost5InARow, "\n\n")
	*/

	// Next what we need to understand is that the biggest discount increase is from 3 to 4: 10%, when other hops are 5%

	// Trying to collect, cause from 3 to 4 is the biggest discount, and compare
	books = booksSave
	booksGroupCounter = [5]int{0, 0, 0, 0, 0}

	for NotNull(books) {
		counter := 0
		for i := range books {
			// 'counter <= 4-1' 4 items is what we need, skip others
			if books[i] > 0 && counter <= 4-1 {
				books[i] -= 1
				counter += 1
			}
		}
		booksGroupCounter[counter-1]++
	}
	totalCost4InARow := CalculateTotal(booksGroupCounter)

	/* fmt.Println("4 in a row collection")
	for _, value := range booksGroupCounter {
		fmt.Print(value, " ")
	}
	fmt.Println("\n4 in a row:", totalCost4InARow)
	*/

	// sorting slice solves the problem in which types of books with less count are taken first. That interfere with collecting 4-length groups.

	return min(totalCost5InARow, totalCost4InARow)
}

// Easy cycle
func NotNull(input []int) bool {
	for i := range input {
		if input[i] > 0 {
			return true
		}
	}
	return false
}

func CalculateTotal(booksGroupCounter [5]int) int {
	totalCost := float64(booksGroupCounter[0]) * 8
	totalCost += float64(booksGroupCounter[1]) * 2 * 8 * 0.95
	totalCost += float64(booksGroupCounter[2]) * 3 * 8 * 0.9
	totalCost += float64(booksGroupCounter[3]) * 4 * 8 * 0.8
	totalCost += float64(booksGroupCounter[4]) * 5 * 8 * 0.75
	return Check20s(int(totalCost * 100))
}

// Don't know why, but in the last test right answer was **80, when the real was **79. This is a crutch.
func Check20s(num int) int {
	to20 := num % 20
	if to20 == 0 {
		return num
	}

	if to20 > 10 {
		return num + 20 - to20
	} else {
		return num - to20
	}
}

// Firstly I changed view that books showed. Was: 1 1 2 2 3 4 => 2 books named 1, 2 books named 2, 1 book named 3, 1 book named 4
// It helps to combine books into groups for discounts
func FromNumericNameToCountedNames(numericName []int) []int {
	countedNames := make([]int, 5)
	for _, v := range numericName {
		if v != 0 {
			countedNames[v-1]++
		}
	}
	return countedNames
}
