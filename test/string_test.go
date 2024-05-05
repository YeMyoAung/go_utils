package test

import (
	"testing"

	"github.com/YeMyoAung/go_utils/utils"
)

func TestString(t *testing.T) {
	t.Run("string test", func(t *testing.T) {
		fake := utils.IsValidEmail("h@.c")
		if fake {
			t.Errorf("expected %v, got %v", true, fake)
		}
	})
}
