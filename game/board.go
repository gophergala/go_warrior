package game

type Board struct {
	Width  int
	Height int
	Spaces map[string]*Space
}
