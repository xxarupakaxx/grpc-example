package game

type Player struct {
	PlayerID int
	Piece    XO
}

type XO int

const (
	X XO = iota
	O
	UNKNOWN
)

type Result int
const (
	WINNER = iota
	LOSE
	DRAW
)

func ConvertToXO(a XO) string {
	switch a {
	case X:
		return "X"
	case O:
		return "O"
	case UNKNOWN:
		return ""
	}

	return ""
}

func ConvertWinner(a Result) string {
	switch a {
	case WINNER:
		return "勝ち"
	case LOSE:
		return "負け"
	case DRAW:
		return "引き分け"
	}

	return ""
}