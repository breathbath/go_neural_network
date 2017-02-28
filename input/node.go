package input

type Node struct {
	Address Coord
	Value float64
}

func NewNode(x, y int, value float64) Node {
	return Node{NewCoord(x, y), value}
}