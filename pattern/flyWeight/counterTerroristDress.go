package flyWeight

type counterTerroristDress struct {
	color string
}

func (c *counterTerroristDress) GetColor() string {
	return c.color
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "green"}
}
