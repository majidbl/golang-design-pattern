package main

import (
	"design_pattern/pattern/mediator"
)
func main() {
	stationManager := mediator.NewStationManger()

	passengerTrain := &mediator.PassengerTrain{
		Mediator: stationManager,
	}
	freightTrain := &mediator.FreightTrain{
		Mediator: stationManager,
	}

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}
