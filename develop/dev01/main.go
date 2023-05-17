package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AlexeyBazhin/wbL2/develop/dev01/task"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/
func main() {
	ntpTime, err := task.GetNTPTime("0.ru.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("time.Now(): %v\n", time.Now())
	fmt.Printf("NTP time: %v\n", ntpTime)
	fmt.Printf("time.Now(): %v\n", time.Now())
}
