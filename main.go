package main

import (
	"fmt"
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

	warrior := characters.NewWarrior()
	slug := &characters.Slug{}
	stairs := &environment.Stairs{}

	elements := map[string]game.Element{
		"2-0": stairs,
		"1-2": warrior,
		"1-1": slug,
	}

	generator.Generate(elements)

	printer := controllers.Printer{
		Board: board,
	}

	printer.PrintBoard()

	space := warrior.Abilities.Feel(game.Forward)

	fmt.Println(space.Element.GetSprite())
	//warrior.Move(game.Forward)

	//printer.PrintBoard()

}
