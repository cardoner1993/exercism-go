package exercise3

import (
	"fmt"
)

func Exercise3(num int) (map[int]int, error) {
	if num < 0 {
		return nil, fmt.Errorf("Input must be a non-negative integer")
	}
	result := make(map[int]int, num)
	for i := 1; i <= num; i++ {
		result[i] = i * i
	}
	return result, nil
}
