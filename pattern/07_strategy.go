package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type (
	SortStrategy interface {
		Sort([]any)
	}

	BubbleSortStrategy struct{}
	QuickSortStrategy  struct{}

	Sorter struct {
		SortStrategy
	}
)

func (bs BubbleSortStrategy) Sort(s []any) {
	fmt.Println("Пузырьковая")
}

func (qs QuickSortStrategy) Sort(s []any) {
	fmt.Println("Быстрая")
}

func (sorter Sorter) SortArr(s []any) {
	sorter.SortStrategy.Sort(s)
}

func main7() {
	sorter := Sorter{BubbleSortStrategy{}}
	sorter.SortArr([]any{1, 3, 2})
	sorter.SortStrategy = QuickSortStrategy{}
	sorter.SortArr([]any{1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2})
}
