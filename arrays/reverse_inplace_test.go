package main

import (
	"reflect"
	"testing"
)

func TestReverseInPlace(t *testing.T) {
	tests := []struct {
		start    int
		end      int
		list     []int
		reversed []int
	}{
		{
			start:    0,
			end:      4,
			list:     []int{1, 2, 3, 4, 5},
			reversed: []int{5, 4, 3, 2, 1},
		},
		{
			start:    1,
			end:      2,
			list:     []int{1, 2, 3, 4, 5},
			reversed: []int{1, 3, 2, 4, 5},
		},
	}
	for i, test := range tests {
		reverseInPlace(test.list, test.start, test.end)
		if !reflect.DeepEqual(test.list, test.reversed) {
			t.Fatalf("Test %d failed. Expected %v, got %v", i, test.reversed, test.list)
		}
	}
}
