package result

type ResultMatrix map[int]*ResultItem

func NewResultMatrix() *ResultMatrix {
	return &make(map[int]ResultMatrix)
}

func (this *ResultMatrix) AddSum(Column int, sum float64) {
	item := this.getResultItem(Column)
	item.Sum = sum
}

func (this *ResultMatrix) AddActivationResult(Column int, actResult float64) {
	item := this.getResultItem(Column)
	item.ActivationResult = actResult
}

func (this *ResultMatrix) AddFinalResult(Column int, finalResult float64) {
	item := this.getResultItem(Column)
	item.FinalResult = finalResult
}

func (this *ResultMatrix) getResultItem(Column int) *ResultItem {
	item, ok := this[Column]
	if !ok {
		item = NewResultItem(Column)
		this[Column] = item
	}

	return item
}