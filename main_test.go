package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 2)
	if result != 5 {
		t.Errorf("Add(2, 2) = %d; want 5", result)
	}
}
