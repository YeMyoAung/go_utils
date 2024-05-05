package test

import (
	"log"
	"testing"

	map_utils "github.com/YeMyoAung/go_utils/utils/map"
)

func TestMap(t *testing.T) {
	t.Run("map test", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		newData, _ := map_utils.Map(data, func(i int) int {
			return i + 1
		})
		log.Println(newData)
	})

}
