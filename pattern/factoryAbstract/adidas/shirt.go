package adidas

import (
	"design_pattern/domain/factoryAbstract"
)

type Shirt struct {
	factoryAbstract.Shirt
}

func (s *Shirt) SetLogo(logo string) {
	s.Logo = logo
}

func (s *Shirt) GetLogo() string {
	return s.Logo
}

func (s *Shirt) SetSize(size int) {
	s.Size = size
}

func (s *Shirt) GetSize() int {
	return s.Size
}