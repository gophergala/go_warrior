package controllers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestLevel(t *testing.T) {
	file, err := os.Open("../levels/level00.json")
	if err != nil {
		t.Error("An error has ocurred:", err)
		return
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("An error has ocurred:", err)
		return
	}

	level := &Level{}

	err = json.Unmarshal(fileBytes, level)
	if err != nil {
		t.Error("An error has ocurred:", err)
		return
	}

	if len(level.WarriorAbilities) < 1 {
		t.Error("No habilities were added")
		return
	}

	if level.WarriorAbilities[0] != "feel" {
		t.Error("Incorrect habilitie")
		return
	}

	if len(level.Elements) < 3 {
		t.Error("Elements missing")
		return
	}
}
