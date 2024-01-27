package test

import (
	"log"
	"testing"

	"github.com/YeMyoAung/go_utils/utils"
)

func TestMap(t *testing.T) {
	t.Run("map test", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		newData, _ := utils.Map(data, func(i int) int {
			return i + 1
		})
		log.Println(newData)
	})

}
