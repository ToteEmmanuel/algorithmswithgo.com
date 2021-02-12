package module01

// Sum will sum up all of the numbers passed
// in and return the result
func SumIt(numbers []int) int {
	acc := 0
	for _, val := range numbers {
		acc += val
	}
	return acc
}

func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum(numbers[1:])
}
