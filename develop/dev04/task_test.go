package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name   string
	input  []string
	output map[string][]string
}{
	{
		"default_test",
		[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
		map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}, "листок": {"листок", "слиток", "столик"}},
	},
	{
		"with_repeats_and_lowercase_test",
		[]string{"пятак", "пятка", "кОтСИЛ", "тяпка", "листок", "ПяТка", "слиток", "столик"},
		map[string][]string{"кОтСИЛ": {"кОтСИЛ", "листок", "слиток", "столик"}, "пятак": {"пятак", "пятка", "тяпка"}},
	},
	{
		"empty_test",
		[]string{},
		map[string][]string{},
	},
	{
		"one_element_empty_test",
		[]string{"пятак"},
		map[string][]string{},
	},
	{
		"one_element_test",
		[]string{"кОт", "пятак", "ток"},
		map[string][]string{"кОт": {"кОт", "ток"}},
	},
}

func TestRun(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			set := FindAnagramSets(test.input)
			if !reflect.DeepEqual(set, test.output) {
				t.Errorf("expected %v, got %v", test.output, set)
			}
		})
	}
}
