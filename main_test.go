package main

import "testing"

func TestHello(t *testing.T) {
	result := getText()
	expected := "Hello world"

	if result != expected {
		t.Errorf("Result: '%s'; Expected '%s'", result, expected)
	}
}
