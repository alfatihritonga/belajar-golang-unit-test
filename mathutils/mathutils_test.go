package mathutils

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on Windows OS")
	}
}

func TestAdd(t *testing.T) {
	result := Add(10, 5)
	expected := 15

	assert.Equal(t, expected, result, "Add(10, 5) = %d; want %d", result, expected)
}

func TestSubstract(t *testing.T) {
	result := Substract(10, 5)
	expected := 5

	assert.Equal(t, expected, result, "Subtract(10, 5) = %d; want %d", result, expected)
}

func TestMultiply(t *testing.T) {
	result := Multiply(10, 5)
	expected := 50

	assert.Equal(t, expected, result, "Multiply(10, 5) = %d; want %d", result, expected)
}

func TestDivide(t *testing.T) {
	result := Divide(10, 5)
	expected := float64(2)

	assert.Equal(t, expected, result, "Divide(10, 5) = %f; want %f", result, expected)
}
