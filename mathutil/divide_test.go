package mathutil

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
		wantErr  bool
	}{
		{6, 3, 2, false},
		{10, 2, 5, false},
		{10, 3, 3, false}, // integer division truncation
		{0, 5, 0, false},
		{-6, 3, -2, false},
		{-6, -3, 2, false},
		{5, 0, 0, true}, // division by zero
	}

	for _, tt := range tests {
		result, err := Divide(tt.a, tt.b)
		if tt.wantErr {
			if err == nil {
				t.Errorf("Divide(%d, %d): expected error, got nil", tt.a, tt.b)
			}
			continue
		}
		if err != nil {
			t.Errorf("Divide(%d, %d): unexpected error: %v", tt.a, tt.b, err)
			continue
		}
		if result != tt.expected {
			t.Errorf("Divide(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
		}
	}
}
