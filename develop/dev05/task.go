package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type SortOptions struct {
	after      int  // -A
	before     int  // -B
	context    int  // -C
	count      bool // -c
	ignoreCase bool // -i
	invert     bool // -v
	fixed      bool // -F
	lineNum    bool // -n
}

func ParseFlags() *SortOptions {
	options := &SortOptions{}

	// Определение флагов командной строки
	flag.IntVar(&options.after, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&options.before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&options.context, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&options.count, "c", false, "количество строк")
	flag.BoolVar(&options.ignoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&options.invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&options.fixed, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&options.lineNum, "n", false, "печатать номер строки")

	// Парсинг аргументов командной строки
	flag.Parse()

	return options
}

func ReadFile(filename string) ([]string, error) {
	strs := make([]string, 0)
	if filename != "" {
		file, err := os.Open(filename)
		if err != nil {
			return strs, err
		}
		defer file.Close()

		buf := bufio.NewScanner(file)
		for buf.Scan() {
			strs = append(strs, buf.Text())
		}
		if err := buf.Err(); err != nil {
			return strs, err
		}
	}
	return strs, nil
}

func WriteFile(filename string, strs []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, v := range strs {
		_, err = file.WriteString(v + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	options := ParseFlags()

	text, err := ReadFile(flag.Arg(1))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading file: ", err)
		os.Exit(1)
	}

	text, err = Grep(flag.Arg(0), text, options)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error grep:", err)
		os.Exit(1)
	}

	if err := WriteFile(flag.Arg(2), text); err != nil {
		fmt.Fprintln(os.Stderr, "error writing file: ", err)
		os.Exit(1)
	}
}

func Grep(pattern string, text []string, options *SortOptions) ([]string, error) {
	var kmpPattern string
	kmpText := make([]string, len(text))
	if options.ignoreCase {
		for i, v := range text {
			kmpText[i] = strings.ToLower(v) // приводим все строки к lowercase
		}
		kmpPattern = strings.ToLower(pattern) // приводим pattern k lowercase
	} else {
		kmpPattern = pattern
		copy(kmpText, text)
	}

	starts := make(map[int][]int)
	if options.fixed {
		for i, str := range text {
			if str == pattern {
				starts[i] = []int{0} // если полное совпадаение по строке
			}
		}
	} else {
		starts = useKMP(kmpPattern, kmpText) // получаем мапу, где ключ - номер строки в начальном тексте, а значение - массив стартовых индексов совпадений
	}

	if options.count {
		var count int
		for _, v := range starts {
			count = count + len(v)
		}
		return []string{fmt.Sprintf("%d", count)}, nil // возвращаем только число совпадений
	}

	if options.invert {
		return InvertSearch(pattern, text, starts, options.lineNum) //инвертированный вывод
	}

	var up, down int
	if options.context > 0 {
		up = options.context / 2
		down = options.context - up
	} else {
		if options.before > 0 {
			up = options.before
		}
		if options.after > 0 {
			down = options.after
		}
	}

	//вывод с учетом before, after, context
	newText := make([]string, 0)
	for i := 0; i < len(text); i++ {
		if _, ok := starts[i]; ok {
			if options.lineNum {
				newText = append(newText, fmt.Sprintf("%d:", i+1))
			}
			for j := i - up; j < i; j++ {
				if j >= 0 {
					newText = append(newText, text[j])
				}
			}
			newText = append(newText, text[i])
			for j := i + 1; j <= i+down; j++ {
				if j < len(text) {
					newText = append(newText, text[j])
				}
			}
			newText = append(newText, "")
		}
	}
	return newText, nil
}

func InvertSearch(pattern string, text []string, starts map[int][]int, lineNum bool) ([]string, error) {
	newText := make([]string, 0)
	lineNums := []string{"Line numbers: "}
	for i := 0; i < len(text); i++ {
		if strStarts, ok := starts[i]; ok {
			if lineNum {
				lineNums = append(lineNums, fmt.Sprintf("%d", i+1))
			}
			counter := 0
			for _, startIndex := range strStarts {
				if counter != startIndex {
					newText = append(newText, text[i][counter:startIndex])
				}
				counter = startIndex + len(pattern)
			}
			if counter != len(text[i]) {
				newText = append(newText, text[i][counter:])
			}
		} else {
			newText = append(newText, text[i])
		}
	}
	if len(lineNums) > 1 {
		lineNums = append(lineNums, "")
		newText = append(lineNums, newText...)
	}
	return newText, nil
}

// использовался КМП алгоритм поиска подстроки в строке. Писал сам.
func useKMP(pattern string, text []string) map[int][]int {
	lps := makeLPSArr(pattern)
	starts := make(map[int][]int)
	for i, str := range text {
		strStarts := KMP(str, pattern, lps)
		if len(strStarts) > 0 {
			starts[i] = append(starts[i], strStarts...)
		}
	}
	return starts
}

func KMP(s, needle string, lps []int) []int {
	starts := make([]int, 0)
	i, j := 0, 0
	for len(s)-i >= len(needle)-j {
		if s[i] == needle[j] {
			i++
			j++
			if j == len(needle) {
				starts = append(starts, i-j)
				j = lps[j-1]
			}
		} else if j == 0 {
			i++
		} else {
			j = lps[j-1]
		}
	}
	return starts
}

func makeLPSArr(s string) []int {
	lps := []int{0}
	for i := 1; i < len(s); i++ {
		j := lps[i-1]
		for j > 0 && s[i] != s[j] {
			j = lps[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		lps = append(lps, j)
	}
	return lps
}
