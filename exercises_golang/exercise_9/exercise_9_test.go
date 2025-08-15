package exercise9

import (
	"testing"
)

func TestExercise9(t *testing.T) {
	result, err := Exercise9("banana apple cherry apple")
	if err != nil {
		t.Errorf("Exercise9 returned an error: %v", err)
		return
	}
	expected := "apple banana cherry"
	if result != expected {
		t.Errorf("Exercise9 returned %q, want %q", result, expected)
	}
	t.Logf("Exercise9(\"banana apple cherry apple\") = %q; want %q", result, expected)
	result, err = Exercise9("")
	if err == nil {
		t.Error("Exercise9 should return an error for empty input")
		return
	}
	expectedErr := "Input string cannot be empty"
	if err.Error() != expectedErr {
		t.Errorf("Exercise9 returned error %q, want %q", err.Error(), expectedErr)
	}
}
