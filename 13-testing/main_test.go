package main

import "testing"

func TestAdd(t *testing.T) {
	tables := []struct {
		name      string
		numberOne float64
		numberTwo float64
		expected  float64
	}{
		{"valid-addition-1", 4, 6, 10},
		{"valid-addition-2", 11, 15, 26},
	}

	for _, tt := range tables {
		result, _ := Add(tt.numberOne, tt.numberTwo)

		if result != tt.expected {
			t.Error("Incorrect addition calculation. Expected:", tt.expected)
		}
	}
}

func TestDivision(t *testing.T) {
	tables := []struct {
		name     string
		dividend float64
		divisor  float64
		expected float64
		isError  bool
	}{
		{"valid-division-1", 10, 5, 2, false},
		{"valid-division-2", 16, 4, 4, false},
		{"division-by-0", 234, 0, 0, true},
	}

	for _, tt := range tables {
		result, err := Divide(tt.dividend, tt.divisor)
		// if expecting an error
		if tt.isError {
			if err == nil {
				t.Error("Expected an error, but did not get one")
			}
		} else {
			if tt.expected != result {
				t.Error("Expected:", tt.expected, "But got:", result)
			}
		}
	}
}

// func TestAdd(t *testing.T) {
// 	result, _ := Add(5, 6) // 11

// 	if result != 11 {
// 		t.Error("Incorrect addition calculation. Expected:", 11)
// 	}

// }

func TestCorrectDivision(t *testing.T) {
	result, err := Divide(10, 5)

	if err != nil {
		t.Error("Did not expect an error, but got one:", err.Error())
	} else {
		if result != 2.0 {
			t.Error("Incorrect division calculation")
		}
	}
}
func TestIncorrectDivision(t *testing.T) {
	_, err := Divide(10, 0)

	if err == nil {
		t.Error("Expected an error, but did not get one")
	}
}
