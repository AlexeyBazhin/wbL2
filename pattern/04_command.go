package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/


/*
Определение
Поведенческий паттерн, который позволяет инкапсулировать в объект всю информацию необходимую для совершения какой-либо операции.

Простыми словами
Комманд представляет действие, а также заключает в себе его параметры. 
Клиент говорит инвокеру выполнить определенное действие. Инвокер запускает действие.
Само действие заключает в себе то, как нужно запуститься - передать выполнение ресиверу.

Основная идея - разделение клиента и получателя.
*/
type (
	Bank struct {
		userBalance int
	}
	Command interface {
		exec() error
	}
	WithdrawMoneyCommand struct {
		sum int
		*Bank
	}
	DepositMoneyCommand struct {
		sum int
		*Bank
	}
	// Invoker
	MobileApplication struct{}
	Client            struct {
		MobileApplication
	}
)

func (bank *Bank) WithdrawMoney(sum int) error {
	if bank.userBalance-sum < 0 {
		return errors.New("Not enough money")
	}
	bank.userBalance -= sum
	return nil
}

func (bank *Bank) DepositMoney(sum int) error {
	bank.userBalance += sum
	return nil
}

func (command WithdrawMoneyCommand) exec() error {
	return command.Bank.WithdrawMoney(command.sum)
}

func (command DepositMoneyCommand) exec() error {
	return command.Bank.DepositMoney(command.sum)
}

func (app MobileApplication) submit(command Command) {
	if err := command.exec(); err != nil {
		fmt.Println("Выкидываю предупреждение на экран")
		return
	}
	fmt.Println("Ререндер баланса")
}

// важно показать, что вызывается метод Invoker'a (mobApp), который принимает команду, которая передает действие receiver'у (bank)
func main4() {
	bank := &Bank{1000}
	client := Client{MobileApplication{}}
	client.MobileApplication.submit(DepositMoneyCommand{1000, bank})
	client.MobileApplication.submit(WithdrawMoneyCommand{2001, bank})
}
