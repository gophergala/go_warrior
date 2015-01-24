package game

type Element interface {
	GetPosition() (int, int)
	GetSprite() string
	GetType() string
}
