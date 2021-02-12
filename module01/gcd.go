package module01

import "fmt"

func GCD(a, b int) int {

	fmt.Printf("%d %d \n", a, b)
	if b == 0 {
		return a
	}
	a, b = b, a%b

	return GCD(a, b)

}

func GCDManual(a, b int) int {
	var min, max int
	if a <= b {
		min, max = a, b
	} else {
		max, min = a, b
	}
	factors := []int{}
	factor := 2
	for min != 1 {
		if min%factor == 0 {
			min /= factor
			if max%factor == 0 {
				max /= factor
				factors = append(factors, factor)
			}
		} else {
			factor++
		}
	}
	fmt.Println(factors)
	acc := 1
	for _, val := range factors {
		acc *= val
	}
	return acc
}
