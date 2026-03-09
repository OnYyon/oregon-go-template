package summator

import "testing"

func TestDigitSummator(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "два положительных числа",
			a:        2,
			b:        3,
			expected: 5,
		},
		{
			name:     "два отрицательных числа",
			a:        -4,
			b:        -6,
			expected: -10,
		},
		{
			name:     "положительное и отрицательное",
			a:        10,
			b:        -3,
			expected: 7,
		},
		{
			name:     "сложение с нулём",
			a:        0,
			b:        5,
			expected: 5,
		},
		{
			name:     "оба нуля",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "большие числа",
			a:        1000000,
			b:        2000000,
			expected: 3000000,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := DigitSummator(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("DigitSummator(%d, %d) = %d, ожидалось %d",
					tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
