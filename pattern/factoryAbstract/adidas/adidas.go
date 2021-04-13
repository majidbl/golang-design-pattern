package adidas

import (
	"design_pattern/domain/factoryAbstract"
)

type Adidas struct {

}
func (a *Adidas) MakeShoe() factoryAbstract.IShoe {
	return &Shoe{
		Shoe: factoryAbstract.Shoe{
			Logo: "adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShirt() factoryAbstract.IShirt {
	return &Shirt{
		Shirt: factoryAbstract.Shirt{
			Logo: "adidas",
			Size: 14,
		},
	}
}