package main

import "testing"

func TestSayHello(t *testing.T) {
	got := sayHello()
	want := "Hello"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

}
