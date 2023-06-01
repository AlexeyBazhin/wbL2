package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Определение
Поведенческий паттерн, который предназначен для организации ответственности между объектами. Состоит из источника и ряда обработчиков.

Простыми словами
Представляет из себя цепочку, в которой каждый объект (кроме последнего) вызывает встроенный в себя объект, пока не справится с обработкой запроса. 
Сам объект и встроенный объект реализуют один интерфейс.
Позволяет уменьшить связанность между объектами.
*/

type (
	Accounter interface {
		setNext(Accounter)
		pay(sum int)
	}

	Account struct {
		next    Accounter
		balance int
	}
	PayPal struct {
		Accounter
	}
	BitCoin struct {
		Accounter
	}
)

func (acc *Account) setNext(accounter Accounter) {
	acc.next = accounter
}

func (acc *Account) pay(sum int) {
	if acc.balance >= sum {
		acc.balance -= sum
		fmt.Printf("Оплата %v. Остаток: %v", sum, acc.balance)
		return
	}
	fmt.Println("Не удалось\n")
	if acc.next != nil {
		acc.next.pay(sum)
		return
	}
	fmt.Println("Обработчик не найден")
}

func (payPal PayPal) setNext(acc Accounter) {
	payPal.Accounter.setNext(acc)
}

func (payPal PayPal) pay(sum int) {
	fmt.Println("PayPal:")
	payPal.Accounter.pay(sum)
}

func (bitCoin BitCoin) setNext(acc Accounter) {
	bitCoin.Accounter.setNext(acc)
}

func (bitCoin BitCoin) pay(sum int) {
	fmt.Println("Bitcoin:")
	bitCoin.Accounter.pay(sum)
}

func main5() {
	payPal := PayPal{&Account{balance: 2000}}
	bitCoin := BitCoin{&Account{balance: 1000}}

	bitCoin.setNext(payPal)
	bitCoin.pay(1500)
}
