package robot

import "fmt"

type Robot interface {
	GetPosition() (int, int, rune)
	SingleMove(rune)
}

type robot struct {
	X           int
	Y           int
	Orientation rune
}

func NewRobot(x, y int, orientation rune) Robot {
	return &robot{
		X:           x,
		Y:           y,
		Orientation: orientation,
	}
}

var moveMap = map[string]rune{
	"EL": 'N',
	"ER": 'S',
	"WL": 'S',
	"WR": 'N',
	"SL": 'E',
	"SR": 'W',
	"NL": 'W',
	"NR": 'E',
}

var moveString = make([]rune, 2)

func (r *robot) GetPosition() (int, int, rune) {
	return r.X, r.Y, r.Orientation
}

func (r *robot) SingleMove(command rune) {
	if command == 'F' {
		switch r.Orientation {
		case 'N':
			r.Y++
		case 'S':
			r.Y--
		case 'E':
			r.X++
		case 'W':
			r.X--
		}
	} else if command == 'L' || command == 'R' {
		moveString[0] = r.Orientation
		moveString[1] = command
		r.Orientation = moveMap[string(moveString)]
	}
}

func (r *robot) String() string {
	return fmt.Sprintf("(%d, %d, %c)", r.X, r.Y, r.Orientation)
}
