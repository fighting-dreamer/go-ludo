package main

import (
	"errors"
	"fmt"
	"os"
)

type Player struct {
	Name string
}

type LudoBoard struct {
	Size    int
	Snakes  map[int]int
	Ladders map[int]int
}

func fileExist(filePath string) (bool, error) {
	_, err := os.Lstat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func readFile(filePath string) (string, error) {
	existsflag, err := fileExist(filePath)
	if err != nil {
		fmt.Println(fmt.Sprintf("error while checking if file %s exist", filePath))
	}
	if existsflag {
		// file exists
		data, err := os.ReadFile(filePath)
		if err != nil {
			return "", err
		}
		return string(data), nil

	}
	// file does not exist
	return "", errors.New("file does not exist")
}

func readBoardConfig() {
	var board LudoBoard
	
	snakesConfigFile := ""
	ladderConfigFile := ""
	fileDataJson, err := readFile(snakesConfigFile)
	if err != nil {
		panic(err);
	}
	
}

func NewLudoBoard() *LudoBoard {
	snakes, ladders := readBoardConfig()
	return &LudoBoard{
		Size:    10,
		Snakes:  map[int]int{},
		Ladders: map[int]int{},
	}
}

type Game struct{}

type IGame interface {
	Register(Player) (int, error)
	Play(User, input any) error
	Status() Game
	Start() (bool, error)
}
