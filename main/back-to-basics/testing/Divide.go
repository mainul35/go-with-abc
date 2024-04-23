package testing

import "errors"

func Divide(x, y float64) (float64, error) {
	var result float64
	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}

	result = x / y
	return result, nil
}
