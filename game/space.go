package game

type Space struct {
	Element Element
}

func NewSpace() *Space {
	return &Space{}
}

func (this *Space) Empty() bool {
	return this.Element == nil
}

func (this *Space) Stairs() bool {
	return !this.Empty() && this.Element.GetType() == "stairs"
}
