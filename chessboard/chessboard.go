package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	totalCount := 0
	for _, square := range cb[rank] {
		if square {
			totalCount++
		}
	}
	return totalCount
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	totalCount := 0
	if len(cb) >= file {
		for _, rank := range cb {
			if rank[file-1] {
				totalCount++
			}
		}
	}
	return totalCount
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	totalCount := 0
	for range cb {
		totalCount += len(cb)
	}
	return totalCount
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	totalCount := 0
	for _, rank := range cb {
		for i := 0; i < len(rank); i++ {
			if rank[i] {
				totalCount++
			}
		}
	}
	return totalCount
}
