package integers

import "testing"

func TestAdder(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, expected int) {
		t.Helper()
		if got != expected {
			t.Errorf("got %q expected %q", got, expected)
		}
	}

	t.Run("Basic Addition", func(t *testing.T) {
		got := Add(2, 2)
		expect := 4
		assertCorrectMessage(t, got, expect)
	})

	t.Run("Subtraction Addition", func(t *testing.T) {
		got := Add(2, -2)
		expect := 0
		assertCorrectMessage(t, got, expect)
	})

}
