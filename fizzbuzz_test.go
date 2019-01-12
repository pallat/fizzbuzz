package fizzbuzz_test

import (
	"testing"

	"github.com/pallat/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("1 should say 1", func(t *testing.T) {
		want := "1"
		get := fizzbuzz.Say(1)
		if want != get {
			t.Errorf("it should say %s but get %s", want, get)
		}
	})
}
