package characters

//import "github.com/go_warrior/abilities"

type Warrior struct {
	Health       int
	AttackPoints int
	x            int
	y            int
}

func GetWarrior(x, y int) *Warrior {
	return &Warrior{
		Health:       20,
		AttackPoints: 3,
		x:            x,
		y:            y,
	}
}

//func (this *Warrior) Health() {
//return this.Health
//}

func (this *Warrior) GetPosition() (int, int) {
	return this.x, this.y
}

func (this *Warrior) GetSprite() string {
	return "@"
}

func (this *Warrior) GetType() string {
	return "warrior"
}
