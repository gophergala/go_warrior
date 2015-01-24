package abilities

type Abilitie interface {
	Perform() (string, error)
	GetName() string
}
