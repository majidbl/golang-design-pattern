package prototype

type Node interface {
	Print(string)
	Clone() Node
}
