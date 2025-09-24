package main

import (
	"fmt"
)

func main() {
	//fmt.Println(myAtoi("    -4612s12"))
	//fmt.Println(myAtoi("    +4612.12"))
	//fmt.Println(myAtoi("✊-461212"))
	//fmt.Println(myAtoi("-91283472332"))
	//fmt.Println(myAtoi("9223372036854775808"))
	//fmt.Println(myAtoi("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
	fmt.Println(myAtoi("2147483648"))
}

func myAtoi(s string) int {
	isNeg := 1
	isSign := false
	isDigit := false
	notZero := false
	max := int(2 << 30)

	digits := make([]int, 0)
Sum:
	for _, char := range s {
		switch char {
		case 32:
			// space
			if isSign || isDigit {
				break Sum
			}
		case 45:
			// -
			if isSign || isDigit {
				break Sum
			}
			isNeg = -1
			isSign = true
		case 43:
			// +
			if isSign || isDigit {
				break Sum
			}
			isSign = true
		case 48:
			isDigit = true
			if notZero {
				digits = append(digits, 0)
			}
		case 49:
			isDigit = true
			notZero = true
			digits = append(digits, 1)
		case 50:
			isDigit = true
			notZero = true
			digits = append(digits, 2)
		case 51:
			isDigit = true
			notZero = true
			digits = append(digits, 3)
		case 52:
			isDigit = true
			notZero = true
			digits = append(digits, 4)
		case 53:
			isDigit = true
			notZero = true
			digits = append(digits, 5)
		case 54:
			isDigit = true
			notZero = true
			digits = append(digits, 6)
		case 55:
			isDigit = true
			notZero = true
			digits = append(digits, 7)
		case 56:
			isDigit = true
			notZero = true
			digits = append(digits, 8)
		case 57:
			isDigit = true
			notZero = true
			digits = append(digits, 9)
		default:
			break Sum
		}

	}
	//Sum:
	sum := 0
	pow := 0
	lenDig := len(digits)
	if lenDig > 10 {
		return returnMax(max, isNeg)
	}
	for i := lenDig; i > 0; i-- {
		dig := digits[i-1]
		for j := 0; j < pow; j++ {
			dig *= 10
		}
		sum += dig
		if sum >= max {
			return returnMax(max, isNeg)
		}
		pow++
	}
	sum *= isNeg
	return sum
}

func returnMax(max, n int) int {
	if n > 0 {
		return max - 1
	}
	return max * -1
}
