package input

type InputMatrix map[int][]*Node

func NewInputMatrix() InputMatrix {
	return make(map[int]InputMatrix)
}

func (this InputMatrix) AddNode(node *Node)  {
	this[node.Address.Y] = append(this[node.Address.Y], node)
}