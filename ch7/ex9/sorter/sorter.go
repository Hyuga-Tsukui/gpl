package sorter

import "golang.org/x/exp/constraints"

type (
	SortKey         string
	LessFunc[T any] func(x, y *T) result
)

type result struct {
	isSorted bool
	isLess   bool
}

func LessByField[T any, U constraints.Ordered](getField func(*T) U) LessFunc[T] {
	return func(x, y *T) result {
		xv, yv := getField(x), getField(y)
		if xv != yv {
			return result{
				isSorted: true,
				isLess:   xv < yv,
			}
		}
		return result{}
	}
}

type Sorter[T any] struct {
	SortKeys    []SortKey
	LessFuncMap map[SortKey]LessFunc[T]
}

func NewSorter[T any](keys []SortKey, lessFuncMap map[SortKey]LessFunc[T]) *Sorter[T] {
	return &Sorter[T]{
		LessFuncMap: lessFuncMap,
		SortKeys:    keys,
	}
}

func (s Sorter[T]) Less(x, y *T) bool {
	for _, k := range s.SortKeys {
		r := s.LessFuncMap[k](x, y)
		if r.isSorted {
			return r.isLess
		}
		continue
	}
	return false
}
