package main

import "testing"

var tests = []struct {
	name    string
	options *SortOptions
	input   []string
	output  []string
}{
	{"simple_test", &SortOptions{1, false, false, false}, []string{"cab 123", "asz 123", "azf 123"}, []string{"asz 123", "azf 123", "cab 123"}},
	{"reverse_test", &SortOptions{1, false, true, false}, []string{"cab 123", "asz 123", "azf 123"}, []string{"cab 123", "azf 123", "asz 123"}},
	{"with_invalid_columns_test", &SortOptions{2, false, false, false}, []string{"cab 1234", "asz", "azf 23", "b"}, []string{"cab 1234", "azf 23", "asz", "b"}},
	{"numeric_test", &SortOptions{2, true, false, false}, []string{"cab 1234", "asz", "azf 23", "b"}, []string{"azf 23", "cab 1234", "asz", "b"}},
	{"numeric_test_reverse", &SortOptions{2, true, true, false}, []string{"cab 1234", "asz", "azf 23", "b"}, []string{"b", "asz", "cab 1234", "azf 23"}},
	{"unique_test", &SortOptions{1, false, false, true}, []string{"c", "a", "d", "c", "a", "d"}, []string{"a", "c", "d"}},
}

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

func TestSortRun(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			strs, _ := SortStringsInFile(test.input, test.options)
			if !isStringsEqual(strs, test.output) {
				t.Fatalf("%v - output is not expected: want: %#v have: %#v", test.name, test.output, strs)
			}
		})
	}
}

// func TestSimple(t *testing.T) {
// 	strs, _ := SortStringsInFile(tests[0].input, tests[0].options)
// 	if !isStringsEqual(strs, tests[0].output) {
// 		t.Fatalf("output is not expected: want: %+#v have: %+#v", tests[0].output, strs)
// 	}
// }

// func TestReverse(t *testing.T) {
// 	strs, _ := SortStringsInFile(tests[1].input, tests[1].options)
// 	if !isStringsEqual(strs, tests[1].output) {
// 		t.Fatalf("output is not expected: want: %+#v have: %+#v", tests[1].output, strs)
// 	}
// }

// func TestWithInvalidColumns(t *testing.T) {
// 	strs, _ := SortStringsInFile(tests[2].input, tests[2].options)
// 	if !isStringsEqual(strs, tests[2].output) {
// 		t.Fatalf("output is not expected: want: %+#v have: %+#v", tests[2].output, strs)
// 	}
// }
