package lesson1

import "testing"

func TestHypotenuse(t *testing.T) {
	result := Hypotenuse(3, 4)
	if result != 5.0 {
		t.Fatalf("expected 5.0, got %v", result)
	}
}
