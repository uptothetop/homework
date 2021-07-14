package fibonacci

import "math"

func Fibonacci(n interface{}) interface{} {
	switch v := n.(type) {
	case uint:
		return fibonnaciUint(v)
	case int:
		return fibonnaciInt(v)
	case float64:
		return fibonnaciFloat(v)
	case float32:
		// Convert F64 -> F32
		return float32(fibonnaciFloat(float64(v)))
	default:
		panic("Fatal: incorrect type. Expecting real number")
	}
}

func fibonnaciUint(n uint) uint {
	// uint
	return uint(fibonacci(int(n)))
}

func fibonnaciInt(n int) int {
	// int
	return int(fibonacci(n))
}

func fibonnaciFloat(n float64) float64 {
	// float
	return fibonacci(int(n))
}

func fibonacci(n int) float64 {
	// Using Binet's formula
	s5 := math.Sqrt(5)
	posn := (1 + s5) / 2
	negn := (1 - s5) / 2
	return (1 / s5) * (math.Pow(posn, float64(n)) - math.Pow(negn, float64(n)))
}
