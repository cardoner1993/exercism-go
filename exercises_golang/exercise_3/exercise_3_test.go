package exercise3

import (
	"reflect"
	"testing"
)

func TestExercise3(t *testing.T) {
	result, err := Exercise3(3)
	expected := map[int]int{1: 1, 2: 4, 3: 9}

	if reflect.DeepEqual(result, expected) && err == nil {
		t.Logf("Exercise3(3) = %v; want %v", result, expected)
	} else {
		t.Errorf("Exercise3(3) = %v, %v; want %v, nil", result, err, expected)
	}
}
