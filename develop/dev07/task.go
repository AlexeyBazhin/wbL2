package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	single := make(chan interface{})

	wg.Add(len(channels))
	for _, v := range channels {
		go func(channel <-chan interface{}) {
			for val := range channel {
				single <- val
			}
			wg.Done()
		}(v)
	}
	go func() {
		wg.Wait()
		close(single)
	}()
	return single
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
			// c <- after
			// fmt.Println("closed", after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(3*time.Second),
	)

	fmt.Printf("fone after %v\n", time.Since(start))

	// реализация, в которой данные записываются в общий канал (в первой реализации при записи хоть 1 элемента выстрелит <-or)
	// а здесь можно слушать в общий канал, куда стекаются все данные
	// sig2 := func(after time.Duration, i int) <-chan interface{} {
	// 	c := make(chan interface{})
	// 	go func() {
	// 		defer close(c)
	// 		time.Sleep(after)
	// 		c <- fmt.Sprint(after, i)
	// 		// fmt.Println("closed", after)
	// 	}()
	// 	return c
	// }

	// start2 := time.Now()
	// orChan := or(
	// 	sig2(1*time.Second, 1),
	// 	sig2(1*time.Second, 2),
	// 	sig2(1*time.Second, 3),
	// 	sig2(2*time.Second, 4),
	// )
	// for v := range orChan {
	// 	fmt.Println(v)
	// }
	// fmt.Printf("fone after 2 %v\n", time.Since(start2))
}
