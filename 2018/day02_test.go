package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2LetterCount(t *testing.T) {
	tests := []struct {
		id                   string
		expected2, expected3 bool
	}{
		{"abcdef", false, false},
		{"bababc", true, true},
		{"abbcde", true, false},
		{"abcccd", false, true},
		{"aabcdd", true, false},
		{"abcdee", true, false},
		{"ababab", false, true},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			seen2, seen3 := day2LetterCount(test.id)
			assert.Equal(t, seen2, test.expected2, "Expected to report having 2 same characters")
			assert.Equal(t, seen3, test.expected3, "Expected to report having 3 same characters")
		})
	}
}
