package game

import "strconv"

type Space struct {
	Element Element
	Board   *Board
	X       int
	Y       int
}

func NewSpace(board *Board, x, y int) *Space {
	return &Space{
		X:     x,
		Y:     y,
		Board: board,
	}
}

func (this *Space) Empty() bool {
	return this.Element == nil
}

func (this *Space) Stairs() bool {
	return !this.Empty() && this.Element.GetType() == "stairs"
}

var (
	RightCoords    = []int{1, 0}
	LeftCoords     = []int{-1, 0}
	ForwardCoords  = []int{0, 1}
	BackwardCoords = []int{0, -1}
)

func (this *Space) GetNext(direction string) *Space {
	switch direction {
	case "right":
		key := strconv.Itoa(this.X+RightCoords[0]) + "-" + strconv.Itoa(this.Y+LeftCoords[1])
		if space, exists := this.Board.Spaces[key]; exists {
			return space
		}
	}
	return nil
}
