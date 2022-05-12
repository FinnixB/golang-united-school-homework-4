package string_sum

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

var (
	space       = rune(' ')
	minus       = rune('-')
	plus        = rune('+')
	positive    = 1
	negative    = -1
	digits      = [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	readDigits1 = 1
	readOp      = 2
	readDigits2 = 3
	finished    = 4
)

func isSpace(value rune) bool {
	if value == space {
		return true
	}

	return false
}

func isDigit(value rune) bool {
	for i := 0; i < 10; i++ {
		if digits[i] == value {
			return true
		}
	}

	return false
}

func isSign(value rune) bool {
	if value == plus || value == minus {
		return true
	}

	return false
}

func skipSpaces(runes []rune) []rune {
	for len(runes) > 0 {
		if !isSpace(runes[0]) {
			return runes
		}

		runes = runes[1:]
	}

	return runes
}

func readDigits(runes []rune) (remain []rune, digits []rune) {
	digits = runes
	remain = runes

	for len(remain) > 0 {
		if isDigit(remain[0]) {
			remain = remain[1:]
		} else {
			digits = digits[0 : len(digits)-len(remain)]
			return
		}
	}

	return
}

func readSign(runes []rune) (remain []rune, sign rune) {
	remain = runes

	if isSign(remain[0]) {
		sign = remain[0]
		remain = remain[1:]
	}

	return
}

func StringSum(input string) (output string, err error) {
	var runes = []rune(input)

	runes = skipSpaces(runes)

	var num int
	var tmprune rune
	var op1 = 0
	var op2 = 0
	var sign = positive
	var digit []rune
	var op rune
	var stage = readDigits1

	if len(runes) == 0 {
		return "", errorEmptyInput
	}

	for len(runes) > 0 {
		fmt.Println(runes)
		fmt.Println(stage)

		switch {
		case isSign(runes[0]):
			runes, tmprune = readSign(runes)
			if stage == readDigits1 || stage == readDigits2 {
				if tmprune == plus {
					sign = positive
				} else {
					sign = negative
				}
			} else if stage == finished {
				return "", errorNotTwoOperands
			} else {
				op = tmprune
				stage = readDigits2
			}
		case isDigit(runes[0]):
			if stage == finished {
				return "", errorNotTwoOperands
			}
			runes, digit = readDigits(runes)
			fmt.Println(digit)
			num, err = strconv.Atoi(string(digit))
			if err != nil {
				return "", fmt.Errorf("Unable to parse int from str: %w", err)
			}
			switch {
			case stage == readDigits1:
				op1 = num * sign
				stage = readOp
			case stage == readDigits2:
				op2 = num * sign
				stage = finished
			}
		case isSpace(runes[0]):
			runes = skipSpaces(runes)
		}
	}

	if stage != finished {
		return "", errorNotTwoOperands
	}

	var result int

	switch {
	case op == plus:
		result = op1 + op2
	case op == minus:
		result = op1 - op2
	}

	return strconv.FormatInt(int64(result), 10), nil
}
