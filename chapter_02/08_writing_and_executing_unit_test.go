package mathops

import "testing"

// Function to be tested
func Add(a, b int) int {
	return a + b
}

// TestAdd tests the Add function
func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 3
	if result != expected {
		t.Errorf("Add(1, 2) = %d; want %d", result, expected)
	}
}