package controllers

import (
	"github.com/go_warrior/game"
	"strconv"
)

type BoardGenerator struct {
	Board *game.Board
}

func (this *BoardGenerator) Generate() {
	for y := 0; y < this.Board.Height; y++ {

		for x := 0; x < this.Board.Width; x++ {
			key := generateKey(x, y)

			space := game.NewSpace()

			this.Board.Spaces[key] = space
		}

	}
}

func generateKey(x, y int) string {
	return strconv.Itoa(x) + "-" + strconv.Itoa(y)
}
