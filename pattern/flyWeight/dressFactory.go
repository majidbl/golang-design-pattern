package flyWeight

import (
	"fmt"

	"design_pattern/domain/flyWeight"
	)

const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerroristDressType terrorist dress type
	CounterTerroristDressType = "ctDress"
)

var (
	DressFactorySingleInstance = &DressFactory{
		DressMap: make(map[string]flyWeight.Dress),
	}
)

type DressFactory struct {
	DressMap map[string]flyWeight.Dress
}

func (d *DressFactory) getDressByType(dressType string) (flyWeight.Dress, error) {
	if d.DressMap[dressType] != nil {
		return d.DressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.DressMap[dressType] = newTerroristDress()
		return d.DressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		d.DressMap[dressType] = newCounterTerroristDress()
		return d.DressMap[dressType], nil
	}

	return nil, fmt.Errorf("wrong dress type passed")
}

func GetDressFactorySingleInstance() *DressFactory {
	return DressFactorySingleInstance
}
