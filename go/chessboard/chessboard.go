package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	files := cb[file]
	counter := 0
	for _, v := range files {
		if v == true {
			counter += 1
		}
	}
	return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank > len(cb) || rank < 1 {
		return 0
	}
	counter := 0
	for _, v := range cb {
		if v[rank-1] == true {
			counter += 1
		}
	}

	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	counter := 0
	for _, v := range cb {
		for range v {
			counter += 1
		}
	}

	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	counter := 0
	for _, v := range cb {
		for _, v2 := range v {
			if v2 == true {
				counter += 1
			}
		}
	}

	return counter
}
