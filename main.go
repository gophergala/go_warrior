package main

import (
	"github.com/go_warrior/characters"
	"github.com/go_warrior/controllers"
	"github.com/go_warrior/game"
)

func main() {
	level := controllers.NewLevel()

	level.Start(getUserFunction(level.Warrior))
}

func getUserFunction(warrior *characters.Warrior) controllers.UserFunction {
	return func() {
		Warrior(warrior)
	}
}

func Warrior(warrior *characters.Warrior) {
	warrior.Walk(game.Forward)
}
