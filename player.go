package main


type Player struct {
	Name string
	Colour string
	Pos int
}

func getStartingPosforcolour(colour string) int {
	switch(colour) {
	case "red":
		return 0
	case "green":
		return 13
	case "yellow":
		return 26
	case "blue" :
		return 39
	default:
		-1
	}
}

func NewPlayer(colour string, name string) (*Player, error) {
	if !isValidColour(colour) {
		return nil, errors.New("Invalid colour")
	}
	return &{
		Name: name,
		Colour: colour,
		Pos: getStartingPosForColour(colour),
	}
}
