package controllers

import "github.com/go_warrior/game"

type BoardGenerator struct {
	Board *game.Board
}

func NewBoardGenerator(board *game.Board) *BoardGenerator {
	return &BoardGenerator{
		Board: board,
	}
}

func (this *BoardGenerator) Generate(elementMap map[string]game.Element) {
	for y := 0; y < this.Board.Height; y++ {

		for x := 0; x < this.Board.Width; x++ {
			key := GenerateKey(x, y)

			space := game.NewSpace(this.Board, x, y)

			if element, exists := elementMap[key]; exists {
				element.SetSpace(space)
				space.Element = element
			}

			this.Board.Spaces[key] = space
		}

	}
}
