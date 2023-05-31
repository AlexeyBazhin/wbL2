package main

import "testing"

var (
	input1 = []string{
		"a	d	f	c",
		"za	4ds	fd	ss	dsdsd",
		"ds:ds  :ds ds",
		"a	d",
	}
	input2 = []string{
		"su:*:0:0:User with special privileges:/:/usr/bin/sh",
		"daemon:*:1:1::/etc:",
		"bin:*:2:2::/usr/bin:",
		"sys:*:3:3::/usr/src:",
		"adm:*:4:4:system administrator:/var/adm:/usr/bin/sh",
		"pierre:*:200:200:Pierre Harper:/home/pierre:/usr/bin/sh",
		"joan:*:202:200:Joan Brown:/home/joan:/usr/bin/sh",
	}
)

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

func TestSimple(t *testing.T) {
	output := []string{
		"d",
		"4ds",
		"ds:ds  :ds ds",
		"d",
	}
	strs, _ := Cut(input1, &CutOptions{"2", "\t", false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestSimpleSeparated(t *testing.T) {
	output := []string{
		"d",
		"4ds",
		"d",
	}
	strs, _ := Cut(input1, &CutOptions{"2", "\t", true})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestRange(t *testing.T) {
	output := []string{
		"su:*:0:0:User with special privileges",
		"daemon:*:1:1:",
		"bin:*:2:2:",
		"sys:*:3:3:",
		"adm:*:4:4:system administrator",
		"pierre:*:200:200:Pierre Harper",
		"joan:*:202:200:Joan Brown",
	}
	strs, _ := Cut(input2, &CutOptions{"1-5", ":", false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestLeftRange(t *testing.T) {
	output := []string{
		"su:*:0:0:User with special privileges",
		"daemon:*:1:1:",
		"bin:*:2:2:",
		"sys:*:3:3:",
		"adm:*:4:4:system administrator",
		"pierre:*:200:200:Pierre Harper",
		"joan:*:202:200:Joan Brown",
	}
	strs, _ := Cut(input2, &CutOptions{"-5", ":", false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestRightRangeAndSeparated(t *testing.T) {
	output := []string{
		"special privileges:/:/usr/bin/sh",
		"",
		"",
		"",
	}
	strs, _ := Cut(input2, &CutOptions{"3-", " ", true})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestRightOutOfRange(t *testing.T) {
	output := []string{
		"/usr/bin/sh",
		"",
		"",
		"",
		"/usr/bin/sh",
		"/usr/bin/sh",
		"/usr/bin/sh",
	}
	strs, _ := Cut(input2, &CutOptions{"7-12", ":", false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestFullOutOfRange(t *testing.T) {
	output := []string{
		"",
		"daemon:*:1:1::/etc:",
		"bin:*:2:2::/usr/bin:",
		"sys:*:3:3::/usr/src:",
		"",
		"",
		"",
	}
	strs, _ := Cut(input2, &CutOptions{"13-20", " ", false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestSetOfFields(t *testing.T) {
	output := []string{
		"0:User with special privileges:/usr/bin/sh",
		"1::",
		"2::",
		"3::",
		"4:system administrator:/usr/bin/sh",
		"200:Pierre Harper:/usr/bin/sh",
		"202:Joan Brown:/usr/bin/sh",
	}
	strs, _ := Cut(input2, &CutOptions{"3,5,7,12,20", ":", false})
	if !isStringsEqual(strs, output) {
		t.Fatalf("output is not expected: want: %#v have: %#v", output, strs)
	}
}

func TestErrorInvalidRangeLeft(t *testing.T) {
	_, err := Cut(input2, &CutOptions{"f-4", ":", false})
	if err.Error() != "invalid field range - left" {
		t.Fatal()
	}
}

func TestErrorInvalidRangeRight(t *testing.T) {
	_, err := Cut(input2, &CutOptions{"4-f", ":", false})
	if err.Error() != "invalid field range - right" {
		t.Fatal()
	}
}

func TestErrorInvalidRange(t *testing.T) {
	_, err := Cut(input2, &CutOptions{"4-1", ":", false})
	if err.Error() != "invalid field range" {
		t.Fatal()
	}
}

func TestErrorInvalidField(t *testing.T) {
	_, err := Cut(input2, &CutOptions{"4,a,6", ":", false})
	if err.Error() != "invalid field value" {
		t.Fatal()
	}
}

func TestErrorInvalidFieldNegative(t *testing.T) {
	_, err := Cut(input2, &CutOptions{"4,-5,2 - 6, 3, -1", ":", false})
	if err.Error() != "invalid field - negative value" {
		t.Fatal()
	}
}

func TestErrorListOfFields(t *testing.T) {
	_, err := Cut(input2, &CutOptions{"", ":", false})
	if err.Error() != "must specify a list of fields" {
		t.Fatal()
	}
}
