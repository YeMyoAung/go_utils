package test

import (
	"testing"

	string_utils "github.com/YeMyoAung/go_utils/utils/string"
)

func TestString(t *testing.T) {
	t.Run("string test", func(t *testing.T) {
		fake := string_utils.IsValidEmail("h@.c")
		if fake {
			t.Errorf("expected %v, got %v", true, fake)
		}
	})
}
