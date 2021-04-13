package factoryMethod

import "design_pattern/domain/factoryMethod"

type musket struct {
	gun
}

func NewMusket() factoryMethod.Gun {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
