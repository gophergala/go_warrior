package characters

import "github.com/go_warrior/game"

type Warrior struct {
	Health       int
	AttackPoints int
	Space        *game.Space
}

func GetWarrior(space, x, y int) *Warrior {
	return &Warrior{
		Health:       20,
		AttackPoints: 3,
	}
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

func (this *Warrior) Move(direction string) error {
	space := this.Space
	nextSpace := space.GetNext(direction)
	nextSpace.Element = this
	space.Element = nil
	return nil
}
