package thefarm

import "errors"

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	scaleMalfunc := errors.New("ErrScaleMalfunction")
	i, err := weightFodder.FodderAmount()
	if cows == 0 {
		return 0, errors.New("division by zero")
	} else if err == scaleMalfunc && i > 0 {
		total := i * 2 / float64(cows)
		return total, nil
	} else if err != nil || err != scaleMalfunc {
		return 0, errors.New("non-scale error")
	} else if i < 0 {
		return 0, errors.New("negative fodder")
	}
	total := i / float64(cows)
	return total, nil
}
