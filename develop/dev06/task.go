package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type CutOptions struct {
	fields    string // -f
	delim     string // -d
	separated bool   // -s
}

func ParseFlags() *CutOptions {
	options := &CutOptions{}

	// Определение флагов командной строки
	flag.StringVar(&options.fields, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&options.delim, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&options.separated, "s", false, "только строки с разделителем")

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
	for i, v := range strs {
		if i != len(strs)-1 {
			v = v + "\n"
		}
		_, err = file.WriteString(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	options := ParseFlags()

	text, err := ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading file: ", err)
		os.Exit(1)
	}

	text, err = Cut(text, options)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error cut:", err)
		os.Exit(1)
	}

	if err := WriteFile(flag.Arg(1), text); err != nil {
		fmt.Fprintln(os.Stderr, "error writing file: ", err)
		os.Exit(1)
	}
}

func Cut(text []string, options *CutOptions) ([]string, error) {
	delim, err := GetDelimeter(options.delim)
	if err != nil {
		return nil, err
	}

	if options.fields == "" {
		return nil, fmt.Errorf("must specify a list of fields")
	}

	newText := make([]string, 0)
	// если у нас диапазон
	if fieldsRange := strings.Split(options.fields, "-"); len(fieldsRange) == 2 {
		left, err := strconv.Atoi(fieldsRange[0])
		if fieldsRange[0] != "" && err != nil {
			return nil, fmt.Errorf("invalid field range - left")
		}
		right, err := strconv.Atoi(fieldsRange[1])
		if fieldsRange[1] != "" && err != nil {
			return nil, fmt.Errorf("invalid field range - right")
		}
		if left == 0 {
			left = 1
		}
		if right != 0 && left > right {
			return nil, fmt.Errorf("invalid field range")
		}

		for _, str := range text {
			splitedStr := strings.Split(str, delim)

			if len(splitedStr) == 1 {
				if !options.separated {
					newText = append(newText, splitedStr...)
				}
				continue
			}

			if right == 0 || right > len(splitedStr) {
				right = len(splitedStr)
			}
			newStr := ""
			if left <= len(splitedStr) {
				newStr = strings.Join(splitedStr[left-1:right], delim)
			}
			newText = append(newText, newStr)
		}
		return newText, nil
	}

	// значит поля записаны через запятую, получаем множество уникальных значений
	fields, err := GetFields(options.fields)
	if err != nil {
		return nil, err
	}

	for _, str := range text {
		splitedStr := strings.Split(str, delim)

		if len(splitedStr) == 1 {
			if !options.separated {
				newText = append(newText, splitedStr...)
			}
			continue
		}

		newStr := ""
		// смотрим указана ли каждая колонка в множестве
		for i, value := range splitedStr {
			if _, ok := fields[i+1]; ok {
				if newStr != "" {
					newStr = newStr + delim
				}
				newStr = newStr + value
			}
		}
		newText = append(newText, newStr)
	}
	return newText, nil
}

func GetDelimeter(delimStr string) (string, error) {
	if utf8.RuneCountInString(delimStr) != 1 {
		return "", fmt.Errorf("invalid delimiter")
	}
	delim := string([]rune(delimStr)[0])
	return delim, nil
}

func GetFields(fieldsStr string) (map[int]struct{}, error) {
	fields := make(map[int]struct{}, 0)
	splitedFields := strings.Split(fieldsStr, ",")
	for _, field := range splitedFields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, fmt.Errorf("invalid field value")
		}
		if num <= 0 {
			return nil, fmt.Errorf("invalid field - negative value")
		}
		fields[num] = struct{}{}
	}
	return fields, nil
}
