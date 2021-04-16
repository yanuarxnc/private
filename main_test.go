package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubble(t *testing.T) {

	testCases := map[string]struct {
		a      []int
		n      int
		expect []int
	}{
		"Test 1": {
			a:      []int{7, 4, 5, 2},
			n:      4,
			expect: []int{2, 4, 5, 7},
		},
	}

	for name, test := range testCases {

		t.Run(name, func(t *testing.T) {
			output := Bubble(test.a, test.n)
			assert.Equal(t, test.expect, output)
		})
	}
}
