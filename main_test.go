package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	want := 3

	if got != want {
		t.Errorf("Add(2,3) = %d, want %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(2, 1)
	want := 1

	if got != want {
		t.Errorf("Subtract(2,3) = %d, want %d", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(2, 2)
	want := 4

	if got != want {
		t.Errorf("Multiply(2,3) = %d, want %d", got, want)
	}
}

func TestDivide(t *testing.T) {
	got := Divide(2, 2)
	want := 1

	if got != want {
		t.Errorf("Divide(2,2) = %d, want %d", got, want)
	}
}

func TestChoice(t *testing.T) {
	got := Choice(1, 2, 2)
	want := 4

	if got != want {
		t.Errorf("Choice(1, 2, 2) = %d, want %d", got, want)
	}
}
