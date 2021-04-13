package flyWeight

type terroristDress struct {
	color string
}

func (t *terroristDress) GetColor() string {
	return t.color
}

func newTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}
