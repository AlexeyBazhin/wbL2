package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».

Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Facade_pattern
*/
type (
	SubsystemA struct {
		s string
	}
	SubsystemB struct {
		a,
		b int
	}

	Facade struct {
		SubsystemA
		SubsystemB
	}
)

func (a *SubsystemA) OperationA() {
	fmt.Println(a.s)
}

func (a *SubsystemA) OperationB() string {
	return a.s
}

func (b *SubsystemB) OperationA() {
	fmt.Println(b.a + b.b)
}

func (b *SubsystemB) OperationB() int {
	return b.a + b.b
}

func (facade *Facade) Operation1() {
	facade.SubsystemA.OperationB()
	facade.SubsystemB.OperationA()
	facade.SubsystemB.OperationB()
}

func (facade *Facade) Operation2() {
	facade.SubsystemB.a = facade.SubsystemB.OperationB()
	facade.SubsystemB.OperationA()
}

func main() {
	// Структурный паттерн, который позволяет скрыть сложность системы, сводя все внешние вызовы к объекту, который делегирует их соответствующим компонентам системы.
	//Простыми словами: Данный паттерн позволяет получить простой интерфейс доступа к сложной системе.
	facade := &Facade{
		SubsystemA: SubsystemA{"A"},
		SubsystemB: SubsystemB{5,3},
	}

	facade.Operation1()
	facade.Operation2()
}
