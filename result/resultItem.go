package result

type ResultItem struct {
	Column int
	Sum float64
	ActivationResult float64
	FinalResult float64
}

func NewResultItem(res int) *ResultItem {
	return &ResultItem{Column:res}
}
