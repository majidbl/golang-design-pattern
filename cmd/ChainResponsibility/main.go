package main

import cr "design_pattern/pattern/ChainResponsibility"
import "design_pattern/domain/ChainResponsibility"
func main() {

	cashier := &cr.Cashier{}

	//Set next for medical department
	medical := &cr.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &cr.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &cr.Reception{}
	reception.SetNext(doctor)

	patient := &ChainResponsibility.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
