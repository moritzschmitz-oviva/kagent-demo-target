package mathutil

import "errors"

// Divide returns the integer division of a by b, or an error if b is zero.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
