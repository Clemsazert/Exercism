package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File = []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard = map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
	if ok := cb[file]; ok == nil {
		return 0
	}
	for _, tile := range cb[file] {
		if tile {
			count += 1
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank > len(cb["A"]) || rank < 1 {
		return 0
	}
	count := 0
	for _, file := range cb {
		if file[rank-1] {
			count += 1
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb["A"])
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for key := range cb {
		count += CountInFile(cb, key)
	}
	return count
}
