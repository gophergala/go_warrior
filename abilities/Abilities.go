package abilities

import "github.com/go_warrior/game"

type Abilities struct {
	Performer Performer
}

func (this *Abilities) Feel(direction game.Direction) *game.Space {
	feel := &Feel{}

	return feel.Perform(this.Performer, direction)
}
