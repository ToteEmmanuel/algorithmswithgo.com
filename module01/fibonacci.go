package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//   Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//
//
// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377
//
func FibonacciLame(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Fibonacci(n int) int {
	result, _ := FibonacciCooler(n)
	return result
}

func FibonacciCooler(n int) (int, int) {
	if n == 0 {
		return 0, 0
	}
	if n == 1 {
		return 1, 0
	}
	prev1, prev2 := FibonacciCooler(n - 1)
	return prev1 + prev2, prev1
}
