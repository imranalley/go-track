package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.
var ErrNegativeFodder = errors.New("negative fodder")
var ErrDivisionByZero = errors.New("division by zero")

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	i, err := weightFodder.FodderAmount()

	if cows == 0 {
		return 0.0, ErrDivisionByZero
	}
	if err == ErrScaleMalfunction {
		i = 2 * i
		err = nil
	}
	if err == nil && i < 0 {
		return 0.0, ErrNegativeFodder
	}
	if err != nil {
		return 0.0, err
	}
	if cows < 0 {
		return 0.0, &SillyNephewError{cows}
	}
	return i / float64(cows), err

}
