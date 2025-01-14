package mathutils

import "testing"

func TestAdd(t *testing.T) {
	result := Add(10, 5)
	expected := 15
	if result != expected {
		t.Errorf("Add(10, 5) = %d; want %d", result, expected)
	}
}

func TestSubstract(t *testing.T) {
	result := Substract(10, 5)
	expected := 5
	if result != expected {
		t.Errorf("Subtract(10, 5) = %d; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(10, 5)
	expected := 50
	if result != expected {
		t.Errorf("Multiply(10, 5) = %d; want %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(10, 5)
	expected := float64(2)
	if result != expected {
		t.Errorf("Divide(10, 5) = %f; want %f", result, expected)
	}
}
