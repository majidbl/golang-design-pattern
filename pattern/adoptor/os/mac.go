package os

import (
	"fmt"

	"design_pattern/domain/adoptor"
	)


type Mac struct {
	Printer adoptor.Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.Printer.PrintFile()
}

func (m *Mac) SetPrinter(p adoptor.Printer) {
	m.Printer = p
}
