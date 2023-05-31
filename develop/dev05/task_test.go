package main

import "testing"

var input = []string{
	"Поддержать флаги:",
	"-A - after печатать +N строк после совпадения",
	"-B - before печатать +N строк до совпадения",
	"-C - context (A+B) печатать ±N строк вокруг совпадения",
	"-c - count (количество строк)",
	"-i - ignore-case (игнорировать регистр)",
	"-v - invert (вместо совпадения, исключать)",
	"-F - fixed, точное совпадение со строкой, не паттерн",
	"-n - line num, печатать номер строки",
} // tests = []struct {
// 	name    string
// 	options *SortOptions
// 	needle  string
// 	input   []string
// 	output  []string
// }{
// 	{"simple_test", &SortOptions{0, 0, 0, false, false, false, false, false}, "count", input, outputSimple},
// }

func isStringsEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// func TestSortRun(t *testing.T) {
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			strs, _ := Grep(test.needle, test.input, test.options)
// 			if !isStringsEqual(strs, test.output) {
// 				t.Fatalf("%v - output is not expected: want: %#v have: %#v", test.name, test.output, strs)
// 			}
// 		})
// 	}
// }

func TestSimple(t *testing.T) {
	output := []string{
		"-c - count (количество строк)",
		"",
	}
	strs, _ := Grep("count", input, &SortOptions{0, 0, 0, false, false, false, false, false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestBefore(t *testing.T) {
	output := []string{
		"-B - before печатать +N строк до совпадения",
		"-C - context (A+B) печатать ±N строк вокруг совпадения",
		"-c - count (количество строк)",
		"",
	}
	strs, _ := Grep("count", input, &SortOptions{0, 2, 0, false, false, false, false, false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestBigAfter(t *testing.T) {
	output := []string{
		"-c - count (количество строк)",
		"-i - ignore-case (игнорировать регистр)",
		"-v - invert (вместо совпадения, исключать)",
		"-F - fixed, точное совпадение со строкой, не паттерн",
		"-n - line num, печатать номер строки",
		"",
	}
	strs, _ := Grep("count", input, &SortOptions{10, 0, 0, false, false, false, false, false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestContext(t *testing.T) {
	output := []string{
		"-B - before печатать +N строк до совпадения",
		"-C - context (A+B) печатать ±N строк вокруг совпадения",
		"-c - count (количество строк)",
		"-i - ignore-case (игнорировать регистр)",
		"-v - invert (вместо совпадения, исключать)",
		"-F - fixed, точное совпадение со строкой, не паттерн",
		"",
	}
	strs, _ := Grep("count", input, &SortOptions{0, 0, 5, false, false, false, false, false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestSeveralMatches(t *testing.T) {
	output := []string{
		"-A - after печатать +N строк после совпадения",
		"",
		"-B - before печатать +N строк до совпадения",
		"",
		"-C - context (A+B) печатать ±N строк вокруг совпадения",
		"",
		"-v - invert (вместо совпадения, исключать)",
		"",
	}
	strs, _ := Grep("совпадения", input, &SortOptions{0, 0, 0, false, false, false, false, false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestSeveralContext(t *testing.T) {
	output := []string{
		"-A - after печатать +N строк после совпадения",
		"-B - before печатать +N строк до совпадения",
		"",
		"-B - before печатать +N строк до совпадения",
		"-C - context (A+B) печатать ±N строк вокруг совпадения",
		"",
		"-C - context (A+B) печатать ±N строк вокруг совпадения",
		"-c - count (количество строк)",
		"",
		"-v - invert (вместо совпадения, исключать)",
		"-F - fixed, точное совпадение со строкой, не паттерн",
		"",
	}
	strs, _ := Grep("совпадения", input, &SortOptions{0, 0, 1, false, false, false, false, false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestSeveralCount(t *testing.T) {
	output := []string{
		"4",
	}
	strs, _ := Grep("совпадения", input, &SortOptions{0, 0, 0, true, false, false, false, false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestIgnoreCase(t *testing.T) {
	output := []string{
		"-c - count (количество строк)",
		"",
	}
	strs, _ := Grep("OunT", input, &SortOptions{0, 0, 0, false, true, false, false, false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestIgnoreCase2(t *testing.T) {
	output := []string{
		"-C - context (A+B) печатать ±N строк вокруг совпадения",
		"",
	}
	strs, _ := Grep("(a+b", input, &SortOptions{0, 0, 0, false, true, false, false, false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestInvert(t *testing.T) {
	output := []string{
		"Поддержать флаги:",
		"-A - after печатать +N строк после совпадения",
		"-B - before печатать +N строк до совпадения",
		"-C - context (A+B) печатать ±N строк вокруг совпадения",
		"-c - ",
		" (количество строк)",
		"-i - ignore-case (игнорировать регистр)",
		"-v - invert (вместо совпадения, исключать)",
		"-F - fixed, точное совпадение со строкой, не паттерн",
		"-n - line num, печатать номер строки",
	}
	strs, _ := Grep("cOunt", input, &SortOptions{0, 0, 0, false, true, true, false, false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestSeveralInvert(t *testing.T) {
	output := []string{
		"Поддержать флаги:",
		"-A - after печатать +N строк после ",
		"я",
		"-B - before печатать +N строк до ",
		"я",
		"-C - context (A+B) печатать ±N строк вокруг ",
		"я",
		"-c - count (количество строк)",
		"-i - ignore-case (игнорировать регистр)",
		"-v - invert (вместо ",
		"я, исключать)",
		"-F - fixed, точное ",
		"е со строкой, не паттерн",
		"-n - line num, печатать номер строки",
	}
	strs, _ := Grep("совпадени", input, &SortOptions{0, 0, 0, false, true, true, false, false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestFixedAndLineNum(t *testing.T) {
	output := []string{
		"5:",
		"-A - after печатать +N строк после совпадения",
		"-B - before печатать +N строк до совпадения",
		"-C - context (A+B) печатать ±N строк вокруг совпадения",
		"-c - count (количество строк)",
		"-i - ignore-case (игнорировать регистр)",
		"",
	}
	strs, _ := Grep("-c - count (количество строк)", input, &SortOptions{1, 3, 0, false, false, false, true, true})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestInvertFixedAndLineNum(t *testing.T) {
	output := []string{
		"Line numbers: ",
		"5",
		"",
		"Поддержать флаги:",
		"-A - after печатать +N строк после совпадения",
		"-B - before печатать +N строк до совпадения",
		"-C - context (A+B) печатать ±N строк вокруг совпадения",
		"-i - ignore-case (игнорировать регистр)",
		"-v - invert (вместо совпадения, исключать)",
		"-F - fixed, точное совпадение со строкой, не паттерн",
		"-n - line num, печатать номер строки",
	}
	strs, _ := Grep("-c - count (количество строк)", input, &SortOptions{0, 0, 0, false, false, true, true, true})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestSeveralInvertFixedAndLineNum(t *testing.T) {
	input := []string{
		"-c - count (количество строк)",
		"Поддержать флаги:",
		"-A - after печатать +N строк после совпадения",
		"-B - before печатать +N строк до совпадения",
		"-C - context (A+B) печатать ±N строк вокруг совпадения",
		"-c - count (количество строк)",
		"-c - count (количество строк) SHISH",
		"-i - ignore-case (игнорировать регистр)",
		"-c - count (количество строк)",
		"-v - invert (вместо совпадения, исключать)",
		"-F - fixed, точное совпадение со строкой, не паттерн",
		"-c - count (количество строк)",
		"-n - line num, печатать номер строки",
		"-c - count (количество строк)",
	}
	output := []string{
		"Line numbers: ",
		"1",
		"6",
		"9",
		"12",
		"14",
		"",
		"Поддержать флаги:",
		"-A - after печатать +N строк после совпадения",
		"-B - before печатать +N строк до совпадения",
		"-C - context (A+B) печатать ±N строк вокруг совпадения",
		"-c - count (количество строк) SHISH",
		"-i - ignore-case (игнорировать регистр)",
		"-v - invert (вместо совпадения, исключать)",
		"-F - fixed, точное совпадение со строкой, не паттерн",
		"-n - line num, печатать номер строки",
	}
	strs, _ := Grep("-c - count (количество строк)", input, &SortOptions{0, 0, 0, false, false, true, true, true})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}
