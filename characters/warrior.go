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
		Abilities: &abilities.Abilities{
			Map: map[string]abilities.Ability{},
		},
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

func (this *Warrior) Walk(direction game.Direction) {
	space := this.Space
	nextSpace := space.GetNext(direction)

	if nextSpace == nil {
		panic(errors.New("Cannot move there!"))
	} else if nextSpace.Element != nil && nextSpace.Element.GetType() == "stairs" {
		panic(errors.New("Level ended"))
	}

	nextSpace.Element = this
	this.Space = nextSpace

	space.Element = nil
}
