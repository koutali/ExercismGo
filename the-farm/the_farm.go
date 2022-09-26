package thefarm

import (
	"errors"
	"strconv"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	numberOfCows int
}

func (e *SillyNephewError) Error() string {
	message := "silly nephew, there cannot be " + strconv.Itoa(e.numberOfCows) + " cows"
	return message
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, error := weightFodder.FodderAmount()

	if error != nil && error != ErrScaleMalfunction {
		return 0.0, error
	}

	if error == ErrScaleMalfunction {
		if fodderAmount > 0.0 {
			fodderAmount *= 2.0
		}
	}

	if fodderAmount < 0 {
		return 0.0, errors.New("negative fodder")
	}

	// division by zero
	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}

	// negative cows
	if cows < 0 {
		return 0.0, &SillyNephewError{numberOfCows: cows}
	}

	return fodderAmount / float64(cows), nil
}
