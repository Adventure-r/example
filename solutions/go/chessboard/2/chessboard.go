package chessboard

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var counter int
	for _, v := range cb[file] {
		if v {
			counter++
		}
	}
	return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	rankTablet := [8]string{"A", "B", "C", "D", "E", "F", "G", "H"}
	if rank < 1 || rank > 8 {
		fmt.Println("Out of range")
		return 0
	}
	var chessCounter int
	file, exists := cb[rankTablet[rank-1]]
	if !exists {
		fmt.Println("Not exists")
		return 0
	}
	for _, isOccupied := range file {
		if isOccupied {
			chessCounter++
		}
	}

	return chessCounter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var counter int
	for _, file := range cb {
		for range file {
			counter++
		}
	}
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var counter int
	for _, file := range cb {
		for _, isOccupied := range file {
			if isOccupied {
				counter++
			}
		}
	}
	return counter
}
