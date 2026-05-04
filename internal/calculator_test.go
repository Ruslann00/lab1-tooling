// Package internal contains tests for calculator functions.
package internal

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        2,
			b:        3,
			expected: 5,
		},
		{
			name:     "negative numbers",
			a:        -2,
			b:        -3,
			expected: -5,
		},
		{
			name:     "positive and negative number",
			a:        10,
			b:        -4,
			expected: 6,
		},
		{
			name:     "zero values",
			a:        0,
			b:        0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Add(tt.a, tt.b)
			if actual != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, actual, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        10,
			b:        4,
			expected: 6,
		},
		{
			name:     "negative result",
			a:        4,
			b:        10,
			expected: -6,
		},
		{
			name:     "negative numbers",
			a:        -10,
			b:        -4,
			expected: -6,
		},
		{
			name:     "subtract zero",
			a:        7,
			b:        0,
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Subtract(tt.a, tt.b)
			if actual != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, actual, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a           int
		b           int
		expected    int
		expectError bool
	}{
		{
			name:        "positive numbers",
			a:           10,
			b:           2,
			expected:    5,
			expectError: false,
		},
		{
			name:        "negative dividend",
			a:           -10,
			b:           2,
			expected:    -5,
			expectError: false,
		},
		{
			name:        "negative divisor",
			a:           10,
			b:           -2,
			expected:    -5,
			expectError: false,
		},
		{
			name:        "division by zero",
			a:           10,
			b:           0,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Divide(tt.a, tt.b)

			if tt.expectError && err == nil {
				t.Fatal("expected error, got nil")
			}

			if !tt.expectError && err != nil {
				t.Fatalf("did not expect error, got %v", err)
			}

			if actual != tt.expected {
				t.Errorf("Divide(%d, %d) = %d; want %d", tt.a, tt.b, actual, tt.expected)
			}
		})
	}
}
