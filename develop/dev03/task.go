package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
type SortOptions struct {
	column  int  // -k
	numeric bool // -n
	reverse bool // -r
	unique  bool // -u
}

func ParseFlags() *SortOptions {
	options := &SortOptions{}

	// Определение флагов командной строки
	flag.IntVar(&options.column, "k", 1, "Указание колонки для сортировки")
	flag.BoolVar(&options.numeric, "n", false, "Сортировать по числовому значению")
	flag.BoolVar(&options.reverse, "r", false, "Сортировать в обратном порядке")
	flag.BoolVar(&options.unique, "u", false, "Не выводить повторяющиеся строки")

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

	strs, err := ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading file: err:", err)
		os.Exit(1)
	}

	strs, err = SortStringsInFile(strs, options)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error sorting: err:", err)
		os.Exit(1)
	}

	if err := WriteFile(flag.Arg(1), strs); err != nil {
		fmt.Fprintln(os.Stderr, "error writing file: err:", err)
		os.Exit(1)
	}
}

func SortStringsInFile(strs []string, options *SortOptions) ([]string, error) {
	if options.unique {
		strs = MakeUnique(strs)
	}
	// с сохранением изначального порядка равных элементов
	sort.SliceStable(strs, func(i, j int) bool {
		if strings.TrimSpace(strs[i]) == "" {
			return false
		}
		iColumns := strings.Split(strs[i], " ")
		jColumns := strings.Split(strs[j], " ")
		// если у первого числа не существует колонки, по которой идет сортировка (по дефолту сортируется по 1)
		// то тогда оно будет больше и вставляться в конец (если сравнивать с другими пустыми, то порядок сохраняется)
		if len(iColumns) < options.column {
			return false
		}
		// аналогично
		if len(jColumns) < options.column {
			return true
		}

		if options.numeric {
			iNum, iErr := strconv.Atoi(iColumns[options.column-1])
			jNum, jErr := strconv.Atoi(jColumns[options.column-1])
			if iErr != nil {
				return false // не удалось запарсить - отправляем в конец
			}
			if jErr != nil {
				return true
			}
			return iNum < jNum
		}
		result := iColumns[options.column-1] < jColumns[options.column-1]
		return result
	})

	if options.reverse {
		for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
			strs[i], strs[j] = strs[j], strs[i]
		}
	}

	return strs, nil
}

func MakeUnique(strs []string) []string {
	unique := make(map[string]struct{})
	newStrs := []string{}
	for _, str := range strs {
		if _, ok := unique[str]; !ok {
			unique[str] = struct{}{}
			newStrs = append(newStrs, str)
		}
	}
	return newStrs
}
