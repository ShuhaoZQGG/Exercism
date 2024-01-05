package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(calculator FodderCalculator, num int) (float64, error) {
	fodderAmount, err := calculator.FodderAmount(num)
	if err != nil {
		return 0, err
	}

	flattingFactor, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return fodderAmount / float64(num) * flattingFactor, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calculator FodderCalculator, num int) (float64, error) {
	if num > 0 {
		result, err := DivideFood(calculator, num)
		if err != nil {
			return 0, err
		}
		return result, nil
	}

	return 0, errors.New("invalid number of cows")
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	number  int
	message string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %s", err.number, err.message)
}

func ValidateNumberOfCows(num int) error {
	if num < 0 {
		return &InvalidCowsError{num, "there are no negative cows"}
	} else if num == 0 {
		return &InvalidCowsError{num, "no cows don't need food"}
	} else {
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
