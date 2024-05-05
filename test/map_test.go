package test

import (
	"testing"

	map_utils "github.com/YeMyoAung/go_utils/utils/map"
)

func TestMap(t *testing.T) {
	data := map[string]int{
		"a": 1,
		"b": 2,
	}

	t.Run("keys test", func(t *testing.T) {
		if len(data) != len(map_utils.Keys(data)) {
			t.Fail()
		}
	})
	t.Run("values test", func(t *testing.T) {
		if len(data) != len(map_utils.Values(data)) {
			t.Fail()
		}
	})

}
