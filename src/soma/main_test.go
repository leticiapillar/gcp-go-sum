package main

import "testing"

func TestMainSuccess(t *testing.T) {
	result := sum(5, 5)
	if result != 10 {
		t.Errorf("Function sum failed, expected %v, got %v", 10, result)
	} else {
		t.Logf("Function sum success")
	}
}
