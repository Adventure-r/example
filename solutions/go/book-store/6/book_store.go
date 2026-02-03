package bookstore

import "fmt"

func Cost(books []int) int {
	fmt.Println("///////////////////////////////////////////////////////////////////////////")
	books = FromNumericNameToCountedNames(books)
	booksSave := make([]int, len(books))
	copy(booksSave, books)
	fmt.Println("booksSave:", booksSave)
	fmt.Println("books:", books)
	// ind 0 - counter 1, ind 1 - counter 2, ind 2 - counter 3,
	// ind 3 - counter 4, ind 4 - counter 5
	var booksGroupCounter [5]int
	// try to collect 5 in a row
	for NotNull(books) {
		counter := 0
		for i := range books {
			if books[i] > 0 {
				fmt.Print(books[i], " => ")
				books[i] -= 1
				counter += 1
				fmt.Println(books[i])
			}

			fmt.Println("iter")
		}
		booksGroupCounter[counter-1]++
	}
	totalCost5InARow := CalculateTotal(booksGroupCounter)

	fmt.Println("5 in a row collection")
	for _, value := range booksGroupCounter {
		fmt.Print(value, " ")
	}
	fmt.Print("\n5 in a row cost:", totalCost5InARow, "\n\n")

	// trying to collect 4 in a row, cause from 3 to 4 is the biggest discount, and compare
	books = booksSave
	fmt.Println("booksSave:", booksSave)
	fmt.Println("books:", books)
	booksGroupCounter = [5]int{0, 0, 0, 0, 0}
	fmt.Println("booksGroupCounter:", booksGroupCounter)
	// try to collect 5 in a row
	for NotNull(books) {
		counter := 0
		for i := range books {
			if books[i] > 0 && counter <= 4 {
				fmt.Print(books[i], " => ")
				books[i] -= 1
				counter += 1
				fmt.Println(books[i])
			}

			fmt.Println("iter")
		}
		booksGroupCounter[counter-1]++
	}
	totalCost4InARow := CalculateTotal(booksGroupCounter)

	fmt.Println("4 in a row collection")
	for _, value := range booksGroupCounter {
		fmt.Print(value, " ")
	}
	fmt.Println("\n4 in a row:", totalCost4InARow)

	return max(totalCost5InARow, totalCost4InARow)
}

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
	return int(totalCost * 100)
}

func FromNumericNameToCountedNames(numericName []int) []int {
	countedNames := make([]int, 5)
	for _, v := range numericName {
		if v != 0 {
			countedNames[v-1]++
		}
	}
	return countedNames
}
