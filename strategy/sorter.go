package strategy

type SortStrategy interface {
	Sort([]int) []int
}

type Sorter struct {
	strategy SortStrategy
}

func NewSorter(s SortStrategy) *Sorter {
	return &Sorter{strategy: s}
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func (s *Sorter) Sort(arr []int) []int {
	return s.strategy.Sort(arr)
}
