package main

import (
	"github.com/google/uuid"
)

type Game struct {
	Id string
	Board *Board
	Players []*Player
}

func NewGame() *Game {
	gameID := uuid.NewString()
	safePos := []int{0, 8, 13, 21, 26, 34, 47}
	return &Game{
		Id:      gameID,
		Board:   NewBoard(safePos),
		Players: []*Player{},
	}
}
