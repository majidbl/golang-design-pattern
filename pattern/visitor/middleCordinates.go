package visitor

import "fmt"

type MiddleCoordinates struct {
	X int
	y int
}

func (a *MiddleCoordinates) VisitForSquare(c *Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the X and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) VisitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *MiddleCoordinates) VisitForRectangle(c *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
