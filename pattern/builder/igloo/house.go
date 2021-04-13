package igloo

import iglooDomain "design_pattern/domain/builder"

type Builder struct {
	iglooDomain.House
}


func NewIglooBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) SetWindowType() {
	b.WindowType = "Snow Window"
}

func (b *Builder) SetDoorType() {
	b.DoorType = "Snow Door"
}

func (b *Builder) SetNumFloor() {
	b.Floor = 1
}

func (b *Builder) GetHouse() iglooDomain.House {
	return iglooDomain.House{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}
