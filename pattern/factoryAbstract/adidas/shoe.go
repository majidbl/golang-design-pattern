package adidas

import (
	"design_pattern/domain/factoryAbstract"
)

type Shoe struct {
	factoryAbstract.Shoe
}
func (s *Shoe) SetLogo(logo string) {
	s.Logo = logo
}

func (s *Shoe) GetLogo() string {
	return s.Logo
}

func (s *Shoe) SetSize(size int) {
	s.Size = size
}

func (s *Shoe) GetSize() int {
	return s.Size
}