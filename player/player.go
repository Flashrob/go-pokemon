package player

type player struct {
	Name string
}

func New(name string) player {
	return player{Name: name}
}
