package main

import (
	"github.com/go_warrior/characters"
	"github.com/go_warrior/controllers"
	"github.com/go_warrior/environment"
	"github.com/go_warrior/game"
)

func main() {
	board := game.Board{
		Width:  3,
		Height: 3,
		Spaces: map[string]*game.Space{},
	}

	generator := controllers.BoardGenerator{
		Board: &board,
	}

	warrior := &characters.Warrior{}
	slug := &characters.Slug{}
	stairs := &environment.Stairs{}

	elements := map[string]game.Element{
		"2-0": stairs,
		"0-2": warrior,
		"1-1": slug,
	}

	generator.Generate(elements)

	printer := controllers.Printer{
		Board: board,
	}

	printer.PrintBoard()

	warrior.Move(game.Forward)

	printer.PrintBoard()
}
