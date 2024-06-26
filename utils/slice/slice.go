package slice_utils

import "fmt"

// Map applies the given function to each item in the input slice and returns a
// new slice with the results.
//
// data []T, fn func(T) R
// []R, error
func Map[T any, R any](data []T, fn func(T) R) ([]R, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("cannot map nil or empty array")
	}
	newData := make([]R, len(data))
	for i, v := range data {
		newData[i] = fn(v)
	}
	return newData, nil
}
