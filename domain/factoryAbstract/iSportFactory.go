package factoryAbstract
type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}
