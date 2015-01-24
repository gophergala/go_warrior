package main

import (
	"github.com/go_warrior/controllers"
	"github.com/go_warrior/game"
)

func main() {
	board := game.Board{
		Width:  2,
		Height: 3,
		Spaces: map[string]*game.Space{},
	}

	generator := controllers.BoardGenerator{
		Board: &board,
	}

	generator.Generate()

	printer := controllers.Printer{
		Board: board,
	}

	printer.PrintBoard()

}
