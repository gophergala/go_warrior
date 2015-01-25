package characters

import (
	"errors"
	"github.com/go_warrior/abilities"
	"github.com/go_warrior/game"
)

type Warrior struct {
	Health       int
	AttackPoints int
	Space        *game.Space
	Abilities    *abilities.Abilities
}

func NewWarrior() *Warrior {
	warrior := &Warrior{
		Health:       20,
		AttackPoints: 3,
		Abilities:    &abilities.Abilities{},
	}

	warrior.Abilities.Performer = warrior

	return warrior
}

func (this *Warrior) SetSpace(space *game.Space) {
	this.Space = space
}

func (this *Warrior) GetSpace() *game.Space {
	return this.Space
}

func (this *Warrior) GetSprite() string {
	return "@"
}

func (this *Warrior) GetType() string {
	return "warrior"
}

func (this *Warrior) Move(direction game.Direction) error {
	space := this.Space
	nextSpace := space.GetNext(direction)

	if nextSpace == nil {
		return errors.New("Cannot move there!")
	}

	nextSpace.Element = this
	space.Element = nil
	return nil
}
