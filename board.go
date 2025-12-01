package main


type Board struct {
	SafePos []int
}

func NewBoard(safePos []int) *Board {
	return &Board{
		SafePos: safePos,
	}
}