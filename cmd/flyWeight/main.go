package main

import (
	"fmt"

	"design_pattern/pattern/flyWeight"
	)

func main() {
	game := flyWeight.NewGame()

	//Add Terrorist
	game.AddTerrorist(flyWeight.TerroristDressType)
	game.AddTerrorist(flyWeight.TerroristDressType)
	game.AddTerrorist(flyWeight.TerroristDressType)
	game.AddTerrorist(flyWeight.TerroristDressType)

	//Add CounterTerrorist
	game.AddCounterTerrorist(flyWeight.CounterTerroristDressType)
	game.AddCounterTerrorist(flyWeight.CounterTerroristDressType)
	game.AddCounterTerrorist(flyWeight.CounterTerroristDressType)

	dressFactoryInstance := flyWeight.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
