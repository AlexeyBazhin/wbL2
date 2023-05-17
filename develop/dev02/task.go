package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	result, err := repeatingRunes("a\\") // см. тесты
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func repeatingRunes(str string) (string, error) {
	var result string

	var curRune rune
	var curCount int
	var isEscape bool
	for _, v := range str {
		// если до этого встретили \, то запоминаем любую текущую руну
		if isEscape {
			isEscape = false
			curRune = v
			continue
		}
		// если встречаем цифры, то прибвляем к текущей сумме с учетом (прим.) e12 => eeeeeeeeeeee
		if number, err := strconv.Atoi(string(v)); err == nil {
			curCount = curCount*10 + number
			continue
		}

		// перезаписываем строку на основе встреченной руны и ее повторений
		var err error
		if result, err = addToResult(result, curRune, curCount); err != nil {
			return "", err
		}

		// обнуляемся, но учитываем \
		curCount = 0
		if v == rune('\\') {
			isEscape = true
			curRune = 0
			continue
		}
		curRune = v
	}
	// если escape-последовательность была объявлена, но не реализована (прим.) "av\"
	if isEscape {
		return "", fmt.Errorf("invalid escape sequence")
	}
	return addToResult(result, curRune, curCount) // возвращаем строку (и ошибку) с последней повторенной n раз руной
}

func addToResult(result string, curRune rune, curCount int) (string, error) {
	// если в начале встретилось число (прим.) "34av"
	if curRune == 0 && curCount != 0 {
		return "", fmt.Errorf("invalid string")
	}
	// если руна без повторений (должна повториться 1 раз)
	if curRune != 0 && curCount == 0 {
		curCount = 1
	}
	for i := 0; i < curCount; i++ {
		result = strings.Join([]string{result, string(curRune)}, "")
	}
	return result, nil
}
