package main

import "testing"

func TestHello(t *testing.T) {

	assertFunc := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Saying hello to names", func(t *testing.T) {
		got := Hello("Swathi")
		want := "Hello, Swathi"

		assertFunc(t, got, want)
	})

	t.Run("saying hello to World", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertFunc(t, got, want)
	})

}
