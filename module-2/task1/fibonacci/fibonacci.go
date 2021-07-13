package fibonacci

func Fibonacci(n interface{}) interface{} {
	switch v := n.(type) {
	case uint:
		return fibonnaciUint(v)
	case int:
		return fibonnaciInt(v)
	case float64:
		return fibonnaciFloat(n)
	default:
		panic("Fatal: incorrect type. Expecting uint, int or float64")
	}
}

func fibonnaciUint(n uint) uint {
	// uint
	return n
}

func fibonnaciInt(n int) int {
	// int
	return n
}

func fibonnaciFloat(n float64) float64 {
	// float
	return n
}
