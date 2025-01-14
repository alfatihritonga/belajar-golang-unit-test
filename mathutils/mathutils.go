package mathutils

func Add(x, y int) int {
	return x + y + 5
}

func Substract(x, y int) int {
	return x - y + 5
}

func Multiply(x, y int) int {
	return x*y + 5
}

func Divide(x, y int) float64 {
	if y == 0 {
		panic("Division by zero")
	}
	return float64(x)/float64(y) + 5
}
