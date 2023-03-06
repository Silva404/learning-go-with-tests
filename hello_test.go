package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Breno")
    want := "hello, Breno"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
