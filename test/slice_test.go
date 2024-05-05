package test

import (
	"testing"

	slice_utils "github.com/YeMyoAung/go_utils/utils/slice"
)

func TestSlic(t *testing.T) {
	t.Run("map test", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		_, err := slice_utils.Map(data, func(i int) int {
			return i + 1
		})
		if err != nil {
			t.Fail()
		}
	})
}
