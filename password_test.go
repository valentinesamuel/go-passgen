package main

import (
	"testing"
)

func TestBuildPasswordPool(t *testing.T) {
	pool := upperCasePool + lowerCasePool + digitsPool + symbolsPool

	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "10 characters",
			input:    10,
			expected: 10,
		},
		{
			name:     "20 characters",
			input:    20,
			expected: 20,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			password, _ := buildPasswordPool(testCase.input, pool)
			if len(password) != testCase.expected {
				t.Errorf("Expected %d characters, got %d", testCase.expected, len(password))
			}
		})
	}
}

func TestGeneratePassword(t *testing.T) {

	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "12 characters",
			input:    12,
			expected: 12,
		},
		{
			name:     "10 characters",
			input:    10,
			expected: 10,
		},
		{
			name:     "20 characters",
			input:    20,
			expected: 20,
		},
		{
			name:     "30 characters",
			input:    30,
			expected: 30,
		},
		{
			name:     "40 characters",
			input:    40,
			expected: 40,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			password, _ := generatePassword(testCase.input, true, true, true, true)
			if len(password) != testCase.expected {
				t.Errorf("Expected %d characters, got %d", testCase.expected, len(password))
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for b.Loop() {
		generatePassword(10, true, true, true, true)
	}
}