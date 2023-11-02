package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	steps := 0
	if n <= 0 {
		err := errors.New("integers must be greater than zero")
		return 0, err
	}
	for ; n != 1; steps++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
	}
	return steps, nil
}
