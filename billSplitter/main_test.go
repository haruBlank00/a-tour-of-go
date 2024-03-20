package main

import (
	"reflect"
	"testing"
)

var input = "momo/150, fanta/100"
var expected = []itemType{
	{name: "momo", price: 150},
	{name: "fanta", price: 100},
}

func ExtractItems(t *testing.T) {
	result := extractItems(input)
	expected := expected

	notDeepEqual := !reflect.DeepEqual(result, expected)

	if notDeepEqual {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSum(t *testing.T) {
	input := expected
	result := sum(input)
	expected := 250.0

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSplit(t *testing.T) {
	input := 250.0
	result := split(input, 2)
	expected := 125.0

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
