package visitor

import "design_pattern/pattern/visitor"

type Shape interface {
	GetType() string
	Accept(visitor.Visitor)
}
