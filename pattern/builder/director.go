package builder

import builderDomain "design_pattern/domain/builder"

type director struct {
	builder builderDomain.Builder
}

func NewDirector(b builderDomain.Builder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) SetBuilder(b builderDomain.Builder) {
	d.builder = b
}

func (d *director) BuildHouse() builderDomain.House {
	d.builder.SetDoorType()
	d.builder.SetWindowType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}
