package utils

import "fmt"

func Map[T any, R any](data []T, fn func(T) R) ([]R, error) {
	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("cannot map nil or empty array")
	}
	newData := make([]R, len(data))
	for i, v := range data {
		newData[i] = fn(v)
	}
	return newData, nil
}
