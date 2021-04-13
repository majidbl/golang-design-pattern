package factoryAbstract

import (
	"design_pattern/domain/factoryAbstract"
	"fmt"

	"design_pattern/pattern/factoryAbstract/adidas"
	"design_pattern/pattern/factoryAbstract/nike"
)

func GetSportsFactory(brand string) (factoryAbstract.ISportsFactory, error) {
	if brand == "adidas" {
		return &adidas.Adidas{}, nil
	}

	if brand == "nike" {
		return &nike.Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}