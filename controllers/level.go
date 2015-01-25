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

func NewLevel() *LevelController {
	board := game.NewBoard(3, 3)

	generator := NewBoardGenerator(board)

	config, err := getConfig("00")
	panicIfError(err)

	warrior := getWarrior(config)

	elements := getElements(config, warrior)

	generator.Generate(elements)

	return &LevelController{
		Printer: &Printer{
			Board: board,
		},
		Warrior: warrior,
	}
}

func (this *LevelController) Start(f UserFunction) {
	defer func() {
		if r := recover(); r != nil {
			endLevel(r)
		}
	}()

	for {
		this.Printer.PrintBoard()
		f()
		if isLevelSucceded() {
			endLevel("Level successful!")
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func endLevel(cause interface{}) {
	var message string

	switch causeType := cause.(type) {
	case string:
		message = causeType
		break
	case error:
		message = causeType.Error()
		break
	default:
		message = "The game ended for unknown reasons"
	}

	fmt.Println(message)
}

func isLevelSucceded() bool {
	return game.STATE == game.SUCCESS
}

func getConfig(levelNumber string) (*Level, error) {
	file, err := os.Open("levels/level" + levelNumber + ".json")
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

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
