// hello_test.go

package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ray")
	want := "Hello, Ray"

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
