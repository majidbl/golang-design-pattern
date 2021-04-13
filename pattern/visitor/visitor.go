package visitor

type Visitor interface {
	VisitForSquare(c *Square)
	VisitForCircle(c *Circle)
	VisitForRectangle(c *Rectangle)
}
