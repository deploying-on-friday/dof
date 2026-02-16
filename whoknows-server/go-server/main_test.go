package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Fatalf("expected %d but got %d", expected, result)
	}
}
