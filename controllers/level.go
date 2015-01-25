package controllers

import (
	"github.com/go_warrior/abilities"
	"github.com/go_warrior/characters"
	"github.com/go_warrior/environment"
	"github.com/go_warrior/game"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Level struct {
	WarriorAbilities []string      `json:"abilities"`
	Elements         []ElementData `json:"elements"`
}

type ElementData struct {
	Type string `json:"type"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

type UserFunction func()

type LevelController struct {
	Printer *Printer
	Warrior *characters.Warrior
}

func NewLevel() (*LevelController, error) {
	board := game.NewBoard(3, 3)

	generator := NewBoardGenerator(board)

	config, err := getConfig()
	if err != nil {
		return nil, err
	}

	warrior := getWarrior(config)

	elements := getElements(config, warrior)

	generator.Generate(elements)

	return &LevelController{
		Printer: &Printer{
			Board: board,
		},
		Warrior: warrior,
	}, nil
}

func (this *LevelController) Start(f UserFunction) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				fmt.Println(x)
			case error:
				fmt.Println(x.Error())
			default:
				fmt.Println("Game ended")
			}
		}
	}()

	for {
		this.Printer.PrintBoard()
		f()
		time.Sleep(1000 * time.Millisecond)
	}
}

func getConfig() (*Level, error) {
	file, err := os.Open("levels/level00.json")
	if err != nil {
		return nil, err
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	level := &Level{}

	err = json.Unmarshal(fileBytes, level)
	if err != nil {
		return nil, err
	}

	return level, nil
}

func getWarrior(config *Level) *characters.Warrior {
	warrior := characters.NewWarrior()

	for _, abilityStr := range config.WarriorAbilities {
		switch abilityStr {

		case "feel":
			warrior.Abilities.Map["feel"] = &abilities.Feel{}
			break

		}
	}

	return warrior
}

func getElements(config *Level, warrior *characters.Warrior) map[string]game.Element {
	elements := map[string]game.Element{}

	for _, element := range config.Elements {
		key := GenerateKey(element.X, element.Y)

		switch element.Type {
		case "warrior":
			elements[key] = warrior
			break

		case "stairs":
			elements[key] = &environment.Stairs{}
			break

		case "slug":
			elements[key] = &characters.Slug{}
			break
		}
	}

	return elements
}
