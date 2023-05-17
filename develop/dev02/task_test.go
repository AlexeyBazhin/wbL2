package main

import (
	"testing"
)

func TestWithoutRepeats(t *testing.T) {
	str := "abcd"
	newStr, err := repeatingRunes(str)
	if err != nil {
		t.Error(err)
	}
	if str != newStr {
		t.Errorf("expected %s, got %s", str, newStr)
	}
}

func TestWithoutRepeatsOneRune(t *testing.T) {
	str := "a"
	newStr, err := repeatingRunes(str)
	if err != nil {
		t.Error(err)
	}
	if str != newStr {
		t.Errorf("expected %s, got %s", str, newStr)
	}
}

func TestWithRepeats(t *testing.T) {
	str := "a4bc2d5e12"
	newStr, err := repeatingRunes(str)
	if err != nil {
		t.Error(err)
	}
	expectedStr := "aaaabccdddddeeeeeeeeeeee"
	if expectedStr != newStr {
		t.Errorf("expected %s, got %s", expectedStr, newStr)
	}
}

func TestInvalidString(t *testing.T) {
	if _, err := repeatingRunes("32534"); err == nil {
		t.Error("expected error")
	} else if err.Error() != "invalid string" {
		t.Error("another error")
	}
}

func TestInvalidStringOneRune(t *testing.T) {
	if _, err := repeatingRunes("3"); err == nil {
		t.Error("expected error")
	} else if err.Error() != "invalid string" {
		t.Error("another error")
	}
}

func TestWithEscape(t *testing.T) {
	str := "a\\4"
	newStr, err := repeatingRunes(str)
	if err != nil {
		t.Error(err)
	}
	expectedStr := "a4"
	if expectedStr != newStr {
		t.Errorf("expected %s, got %s", expectedStr, newStr)
	}
}

func TestWithEscapes(t *testing.T) {
	str := "a\\4\\5\\6"
	newStr, err := repeatingRunes(str)
	if err != nil {
		t.Error(err)
	}
	expectedStr := "a456"
	if expectedStr != newStr {
		t.Errorf("expected %s, got %s", expectedStr, newStr)
	}
}

func TestWithEscapesWithRepeats(t *testing.T) {
	str := "a\\45\\53\\62"
	newStr, err := repeatingRunes(str)
	if err != nil {
		t.Error(err)
	}
	expectedStr := "a4444455566"
	if expectedStr != newStr {
		t.Errorf("expected %s, got %s", expectedStr, newStr)
	}
}

func TestWithSlashEscape(t *testing.T) {
	str := "a\\\\3" // a\\3
	newStr, err := repeatingRunes(str)
	if err != nil {
		t.Error(err)
	}
	expectedStr := "a\\\\\\" // a\\\
	if expectedStr != newStr {
		t.Errorf("expected %s, got %s", expectedStr, newStr)
	}
	// t.Log(newStr)
}

func TestWithInvalidEscape(t *testing.T) {
	if _, err := repeatingRunes("a\\"); err == nil {
		t.Error("expected error")
	} else if err.Error() != "invalid escape sequence" {
		t.Error("another error")
	}
}

func TestEmptyString(t *testing.T) {
	str := ""
	newStr, err := repeatingRunes(str)
	if err != nil {
		t.Error(err)
	}
	expectedStr := ""
	if expectedStr != newStr {
		t.Errorf("expected %s, got %s", expectedStr, newStr)
	}
}
