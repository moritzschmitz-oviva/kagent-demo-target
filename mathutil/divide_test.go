package mathutil

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{6, 3, 2},
		{10, 2, 5},
		{10, 3, 3}, // Test integer division truncation
		{0, 5, 0},
		{-6, 3, -2},
		{-6, -3, 2},
	}

	for _, tt := range tests {
		result := Divide(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Divide(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
		}
	}
}
