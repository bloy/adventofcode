package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIsReal tests the IsRealFunction
func TestIsReal(t *testing.T) {
	tests := []struct {
		name     string
		room     Room
		expected bool
	}{
		{
			name:     "aaaa-bbb-z-y-x-123[abxyz]",
			room:     Room{"aaaa-bbb-z-y-x", 123, "abxyz"},
			expected: true,
		},
		{
			name:     "a-b-c-d-e-f-g-h-987[abcde]",
			room:     Room{"a-b-c-d-e-f-h", 987, "abcde"},
			expected: true,
		},
		{
			name:     "not-a-real-room-404[oarel]",
			room:     Room{"not-a-real-room", 404, "oarel"},
			expected: true,
		},
		{
			name:     "totally-real-room-200[decoy]",
			room:     Room{"totally-real-room", 200, "decoy"},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.room.IsReal(), test.expected)
		})
	}
}

// TestDecryptName tests the name decription
func TestDecryptName(t *testing.T) {
	r := Room{"qzmt-zixmtkozy-ivhz", 343, "abcd"}
	assert.Equal(t, r.DecryptedName(), "very encrypted name")
}
