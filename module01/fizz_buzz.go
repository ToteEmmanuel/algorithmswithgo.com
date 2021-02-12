package module01

import "fmt"

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	var result string
	for i := 1; i <= n; i++ {
		three := i%3 == 0
		five := i%5 == 0
		switch {
		case three && five:
			result += "Fizz Buzz"
		case three:
			result += "Fizz"
		case five:
			result += "Buzz"
		default:
			result += fmt.Sprint(i)
		}
		if i != n {
			result += ", "
		}
	}
	fmt.Println(result)
}
