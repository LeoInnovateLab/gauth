package utils

import "testing"

func TestSafeStringConvert(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"normal string", "hello", "hello"},
		{"empty string", "", ""},
		{"nil value", nil, ""},
		{"number type", 123, ""},
		{"boolean type", true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SafeStringConvert(tt.input)
			if result != tt.expected {
				t.Errorf("SafeStringConvert(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSafeIntConvert(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int
	}{
		{"integer", 123, 123},
		{"float", 123.45, 123},
		{"zero value", 0, 0},
		{"nil value", nil, 0},
		{"string", "123", 0},
		{"boolean type", true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SafeIntConvert(tt.input)
			if result != tt.expected {
				t.Errorf("SafeIntConvert(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
