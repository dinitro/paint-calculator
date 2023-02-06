package userin

import (
	"strconv"
	"testing"
)

func TestInFloat(t *testing.T) {
	// Input
	input := "5.5"

	// Expected
	expected := 5.5

	// Compare result with expected output
	uFloat, _ := strconv.ParseFloat(input, 64)
	if uFloat != expected {
		t.Errorf("Expected %f but got %f", expected, uFloat)
	}
}

func TestInInt(t *testing.T) {
	// Input
	input := "5"

	// Expected
	expected := 5

	// Compare result with expected output
	uInt, _ := strconv.Atoi(input)
	if uInt != expected {
		t.Errorf("Expected %d but got %d", expected, uInt)
	}
}
