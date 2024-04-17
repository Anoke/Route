package main

import (
	"strings"
	"testing"
)

func TestEmptyInput(t *testing.T) {
	input := ""
	expected := ""
	actual := formatText(input)
	if actual != expected {
		t.Errorf("Empty input test failed, expected '%s', got '%s'", expected, actual)
	}
}

func TestNoCommasInput(t *testing.T) {
	input := "This is a test input without commas"
	expected := "This is a test input without commas"
	actual := formatText(input)
	if actual != expected {
		t.Errorf("No commas input test failed, expected '%s', got '%s'", expected, actual)
	}
}

func TestOneWordInput(t *testing.T) {
	input := "Hello"
	expected := "Hello"
	actual := formatText(input)
	if actual != expected {
		t.Errorf("One word input test failed, expected '%s', got '%s'", expected, actual)
	}
}

func TestOneCommaInput(t *testing.T) {
	input := ","
	expected := ","
	actual := formatText(input)
	if actual != expected {
		t.Errorf("One comma input test failed, expected '%s', got '%s'", expected, actual)
	}
}

func TestLongWordInput(t *testing.T) {
	longWord := strings.Repeat("a", 1000)
	input := longWord
	expected := longWord
	actual := formatText(input)
	if actual != expected {
		t.Errorf("Long word input test failed, expected '%s', got '%s'", expected, actual)
	}
}

func TestOnlySpacesInput(t *testing.T) {
	input := "          "
	expected := ""
	actual := formatText(input)
	if actual != expected {
		t.Errorf("Only spaces input test failed, expected '%s', got '%s'", expected, actual)
	}
}

func TestMultipleCommasInput(t *testing.T) {
	input := ",,,"
	expected := ", , ,"
	actual := formatText(input)
	if actual != expected {
		t.Errorf("Multiple commas input test failed, expected '%s', got '%s'", expected, actual)
	}
}

func TestMultipleSpacesInput(t *testing.T) {
	input := "This    is      a    test    input"
	expected := "This is a test input"
	actual := formatText(input)
	if actual != expected {
		t.Errorf("Multiple spaces input test failed, expected '%s', got '%s'", expected, actual)
	}
}

func TestInputWithWordsAndCommas(t *testing.T) {
	input := "Hello, this is a test, input"
	expected := "Hello, this is a\ntest, input"
	actual := formatText(input)
	if actual != expected {
		t.Errorf("Input with words and commas test failed, expected '%s', got '%s'", expected, actual)
	}
}
