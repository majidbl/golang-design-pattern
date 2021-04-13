package normal

import (
	normalDomain "design_pattern/domain/builder"
)

type Builder struct {
	normalDomain.House

}


func NewNormalBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) SetWindowType() {
	b.WindowType = "Wooden Window"
}

func (b *Builder) SetDoorType() {
	b.DoorType = "Wooden Door"
}

func (b *Builder) SetNumFloor() {
	b.Floor = 2
}

func (b *Builder) GetHouse() normalDomain.House {
	return normalDomain.House{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}
