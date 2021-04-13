package main

import (
	"fmt"

	"design_pattern/pattern/decorator"
	)

func main() {

	pizza := &decorator.VeggeMania{}

	//Add cheese topping
	pizzaWithCheese := &decorator.CheeseTopping{
		Pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &decorator.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}