package characters

type CharacterCommons struct {
	health       int
	attackPoints int
}

type Character interface {
	Health() int
}
