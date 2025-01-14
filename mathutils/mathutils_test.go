package mathutils

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// before unit test running
	fmt.Println("BEFORE UNIT TEST")
	m.Run()

	// after unit test running
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on Windows OS")
	}
}

func TestAdd(t *testing.T) {
	t.Run("10 plus 5", func(t *testing.T) {
		result := Add(10, 5)
		expected := 15

		assert.Equal(t, expected, result, "Add(10, 5) = %d; want %d", result, expected)
	})
	t.Run("20 plus 15", func(t *testing.T) {
		result := Add(20, 15)
		expected := 35

		assert.Equal(t, expected, result, "Add(20, 15) = %d; want %d", result, expected)
	})
	t.Run("10 plus 0", func(t *testing.T) {
		result := Add(10, 0)
		expected := 10

		assert.Equal(t, expected, result, "Add(10, 0) = %d; want %d", result, expected)
	})
}

func TestSubstract(t *testing.T) {
	t.Run("10 minus 5", func(t *testing.T) {
		result := Subtract(10, 5)
		expected := 5

		assert.Equal(t, expected, result, "Subtract(10, 5) = %d; want %d", result, expected)
	})
	t.Run("-20 minus 5", func(t *testing.T) {
		result := Subtract(-20, 5)
		expected := -25

		assert.Equal(t, expected, result, "Subtract(-20, 5) = %d; want %d", result, expected)
	})
	t.Run("-20 minus -5", func(t *testing.T) {
		result := Subtract(-20, -5)
		expected := -15

		assert.Equal(t, expected, result, "Subtract(-20, -5) = %d; want %d", result, expected)
	})
}

func TestMultiply(t *testing.T) {
	t.Run("10 times 5", func(t *testing.T) {
		result := Multiply(10, 5)
		expected := 50

		assert.Equal(t, expected, result, "Multiply(10, 5) = %d; want %d", result, expected)
	})
	t.Run("10 times 0", func(t *testing.T) {
		result := Multiply(10, 0)
		expected := 0

		assert.Equal(t, expected, result, "Multiply(10, 0) = %d; want %d", result, expected)
	})
	t.Run("-2 times 5", func(t *testing.T) {
		result := Multiply(-2, 5)
		expected := -10

		assert.Equal(t, expected, result, "Multiply(-2, 5) = %d; want %d", result, expected)
	})
	t.Run("-2 times -5", func(t *testing.T) {
		result := Multiply(-2, -5)
		expected := 10

		assert.Equal(t, expected, result, "Multiply(-2, -5) = %d; want %d", result, expected)
	})
}

func TestDivide(t *testing.T) {
	t.Run("10 divided by 5", func(t *testing.T) {
		result := Divide(10, 5)
		expected := float64(2)

		assert.Equal(t, expected, result, "Divide(10, 5) = %f; want %f", result, expected)
	})
	t.Run("10 divided by -5", func(t *testing.T) {
		result := Divide(10, -5)
		expected := float64(-2)

		assert.Equal(t, expected, result, "Divide(10, -5) = %f; want %f", result, expected)
	})
	t.Run("-10 divided by -5", func(t *testing.T) {
		result := Divide(-10, -5)
		expected := float64(2)

		assert.Equal(t, expected, result, "Divide(-10, -5) = %f; want %f", result, expected)
	})
}
