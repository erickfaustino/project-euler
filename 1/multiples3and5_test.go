package main

import "testing"

var testCases = []struct {
	in  int
	out int
}{
	{10, 23},
	{11, 33},
	{20, 78},
	{100, 2318},
}

func TestMultiples3and5(t *testing.T) {
	for _, tt := range testCases {
		v := sumOfMultiples(tt.in)
		if v != tt.out {
			t.Fatalf("result %d is different from expected: %d", tt.out, v)
		}
	}
}
