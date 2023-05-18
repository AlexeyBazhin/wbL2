package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func FindAnagramSets(words []string) map[string][]string {
	anagramSets := make(map[string][]string)

LOOP:
	for _, word := range words {
		wordLower := strings.ToLower(word)
		sortedWord := sortString(wordLower)

		for key, set := range anagramSets {
			sortedKey := sortString(strings.ToLower(key))
			// если отсортированный ключ равен отсоритрованному слову (оба в нижнем регистре)
			if sortedWord == sortedKey {
				for _, value := range set {
					if wordLower == strings.ToLower(value) {
						continue LOOP // если слово уже есть в множестве (регистр не важен - проверяем в нижнем), то переходим к следующему слову
					}
				}
				// если есть подходящий ключ, то добавляем в множество и переходим к следующему слову
				anagramSets[key] = append(anagramSets[key], word)
				continue LOOP
			}
		}
		// устанавливаем ключ
		anagramSets[word] = []string{word}
	}

	// удаляем множества из одного элемента
	for key, value := range anagramSets {
		if len(value) < 2 {
			delete(anagramSets, key)
		}
	}

	return anagramSets
}

func sortString(s string) string {
	sortedRunes := []rune(s)
	sort.Slice(sortedRunes, func(i, j int) bool {
		return sortedRunes[i] < sortedRunes[j]
	})
	return string(sortedRunes)
}

func main() {
	words := []string{"пятак", "пятка", "кОтСИЛ", "тяпка", "листок", "ПяТка", "слиток", "столик"}

	anagramSets := FindAnagramSets(words)

	for key, value := range anagramSets {
		fmt.Printf("Множество анаграмм по ключу [%v]: %v\n", key, value)
	}
}
