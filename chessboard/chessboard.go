package chessboard

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	fileValues, exists := cb[file]
	if !exists {
		return 0
	}

	numOccupied := 0
	for _, fileValue := range fileValues {
		if fileValue {
			numOccupied++
		}
	}

	return numOccupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank <= 0 || rank > 8 {
		return 0
	}

	occupiedRanks := 0
	for key, fileValues := range cb {
		fmt.Println(key, fileValues)
		if fileValues[rank-1] {
			occupiedRanks++
		}
	}

	return occupiedRanks
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb["A"])
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	numOccupied := 0
	for _, fileValues := range cb {
		for _, fileValue := range fileValues {
			if fileValue {
				numOccupied++
			}
		}
	}
	return numOccupied
}
