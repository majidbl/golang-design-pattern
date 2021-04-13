package nike

import (
	"design_pattern/domain/factoryAbstract"
)

type Nike struct{
}

func (n *Nike) MakeShoe() factoryAbstract.IShoe {
	return &Shoe{
		Shoe: factoryAbstract.Shoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

func (n *Nike) MakeShirt() factoryAbstract.IShirt {
	return &Shirt{
		Shirt: factoryAbstract.Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}