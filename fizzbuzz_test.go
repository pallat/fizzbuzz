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
	t.Run("2 should say 2", func(t *testing.T) {
		want := "2"
		get := fizzbuzz.Say(2)
		if want != get {
			t.Errorf("it should say %s but get %s", want, get)
		}
	})
	t.Run("3 should say Fizz", func(t *testing.T) {
		want := "Fizz"
		get := fizzbuzz.Say(3)
		if want != get {
			t.Errorf("it should say %s but get %s", want, get)
		}
	})
}
