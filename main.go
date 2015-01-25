package main

import (
	//"fmt"
	"github.com/go_warrior/characters"
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

	warrior := &characters.Warrior{}

	elements := map[string]game.Element{
		//"1-2": &characters.Warrior{},
		"0-2": warrior,
	}

	generator.Generate(elements)

	printer := controllers.Printer{
		Board: board,
	}

	printer.PrintBoard()

	//nextspace := board.Spaces["0-2"].GetNext("right")

	//fmt.Println(nextspace.Element.GetSprite())

	warrior.Move("right")

	printer.PrintBoard()
}
