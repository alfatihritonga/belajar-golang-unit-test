package mathutils

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkTableAdd(b *testing.B) {
	benchmarks := []struct {
		name string
		x    int
		y    int
	}{
		{"Add(10, 5)", 10, 5},
		{"Add(95, 20)", 95, 20},
		{"Add(190, 35)", 190, 35},
		{"Add(0, 105)", 0, 105},
		{"Add(105, 95)", 105, 95},
		{"Add(130, 215)", 130, 215},
		{"Add(1000, 225)", 1000, 225},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			Add(benchmark.x, benchmark.y)
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	b.Run("Add(100, 90)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Add(100, 90)
		}
	})
	b.Run("Add(205, 100)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Add(205, 100)
		}
	})
	b.Run("Add(50, 90)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Add(50, 90)
		}
	})
}

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

func TestTableAdd(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		y        int
		expected int
	}{
		{
			name:     "10 plus 5",
			x:        10,
			y:        5,
			expected: 15,
		},
		{
			name:     "10 plus -5",
			x:        10,
			y:        -5,
			expected: 5,
		},
		{
			name:     "-10 plus -5",
			x:        -10,
			y:        -5,
			expected: -15,
		},
		{
			name:     "10 plus 0",
			x:        10,
			y:        0,
			expected: 10,
		},
		{
			name:     "0 plus 5",
			x:        0,
			y:        5,
			expected: 5,
		},
		{
			name:     "0 plus 0",
			x:        0,
			y:        0,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Add(test.x, test.y)
			assert.Equal(t, test.expected, result, "Add(%d, %d) = %d; want %d", test.x, test.y, result, test.expected)
		})
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
