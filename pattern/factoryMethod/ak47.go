package factoryMethod

import "design_pattern/domain/factoryMethod"

type ak47 struct {
	gun
}

func NewAk47() factoryMethod.Gun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
