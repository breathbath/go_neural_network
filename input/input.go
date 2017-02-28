package input

type Input struct {
	X int
	Value float64
}

func (this Input) NewInput(x int, value float64){
	return Input{x, value}
}
