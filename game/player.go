package game

type Player struct {
	PlayerID int
	Piece    XO
}

type XO int

const (
	X XO = iota
	O
)
