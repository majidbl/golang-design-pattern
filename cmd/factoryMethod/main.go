package main

import (
	"fmt"

	"design_pattern/pattern/factoryMethod"

	factoryDomain "design_pattern/domain/factoryMethod"
	)

func main() {
	ak47, _ := factoryMethod.GetGun("ak47")
	musket, _ := factoryMethod.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g factoryDomain.Gun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
