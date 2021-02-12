package module01

import (
	"fmt"
	"math"
	"strconv"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	acc := 0
	for i, val := range value {
		acc += ParseInt(val) * int(math.Pow(float64(base), float64((len(value)-i-1))))
	}
	return acc
}

func ParseInt(val rune) int {
	response := 0
	switch val {
	case 'A':
		response = 10
	case 'B':
		response = 11
	case 'C':
		response = 12
	case 'D':
		response = 13
	case 'E':
		response = 14
	case 'F':
		response = 15
	default:
		response, _ = strconv.Atoi(string(val))
	}
	return response
}

func fancierBaseToDec(value string, base int) int {
	acc := 0
	multiplier := 1
	for i := len(value) - 1; i >= 0; i-- {
		var val int
		fmt.Sscanf(string(value[i]), "%X", &val)
		acc += multiplier * val
		multiplier *= base
		fmt.Println(i)
	}
	return acc
}
