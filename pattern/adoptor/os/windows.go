package os

import (
	"design_pattern/domain/adoptor"
	"fmt"
)


type Windows struct {
	Printer adoptor.Printer
}

func (m *Windows) Print() {
	fmt.Println("Print request for mac")
	m.Printer.PrintFile()
}

func (m *Windows) SetPrinter(p adoptor.Printer) {
	m.Printer = p
}

