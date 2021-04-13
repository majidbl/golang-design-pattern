package main

import (
	"fmt"

	"design_pattern/pattern/adoptor/os"
	"design_pattern/pattern/adoptor/printer"
)

func main() {

	hpPrinter := &printer.HP{}
	epsonPrinter := &printer.Epson{}

	macComputer := &os.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &os.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
