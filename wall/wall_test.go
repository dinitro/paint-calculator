package wall

import (
	"reflect"
	"testing"
)

func TestDisplayWalls(t *testing.T) {
	width := []float64{1.0, 3.0, 5.0}
	height := []float64{2.0, 4.0, 6.0}
	expected := []string{"1x2", "3x4", "5x6"}

	// Call function
	display := displayWalls(width, height)

	// Call & Compare
	if !reflect.DeepEqual(display, expected) {
		t.Errorf("Expected %v but got %v", expected, display)
	}
}
