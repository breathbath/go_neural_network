package process

import (
	"github.com/breathbath/nn2/result"
	"github.com/breathbath/nn2/input"
)

type Processor struct {
	Results result.ResultMatrix
	Inputs input.InputMatrix
}

func NewProcessor(rm result.ResultMatrix, ints input.InputMatrix) *Processor {
	return &Processor{rm, ints}
}

func (this *Processor) ForwardPropagation(inputs ...input.Input) {
	this.addInputsToMatrix(inputs)

}

func (this *Processor) addInputsToMatrix(inputs[]input.Input)  {
	var newNode *input.Node
	for _, inputItem := range inputs {
		newNode = input.NewNode(0, inputItem.X, inputItem.Value)
		this.Inputs.AddNode(newNode)
	}
}