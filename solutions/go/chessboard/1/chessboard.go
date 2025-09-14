package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) (count int) {

	for _, value := range cb[file] {
		if value {
			count++
		}
	}

	return

}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (count int) {

	// Guard
	if rank > 8 || rank < 1 {
		return 0
	}

	for _, value := range cb {
		if value[rank-1] {
			count++
		}
	}

	return

}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb["A"]) * len(cb)
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (count int) {

	for file, _ := range cb {
		count += CountInFile(cb, file)
	}

	return

}
