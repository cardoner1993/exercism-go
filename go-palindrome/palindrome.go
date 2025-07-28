package palindrome

import (
	"fmt"
	"sort"
	"strconv"
)

// Result holds the palindrome and its factor pairs
type Product struct {
	Value          int
	Factorizations [][2]int
}

func IsPalindrome(n int) bool {
	str := strconv.Itoa(n)
	begin, end := 0, len(str)-1
	for begin < end {
		if str[begin] != str[end] {
			return false
		}
		begin++
		end--
	}
	return true
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, fmt.Errorf("fmin > fmax")
	}
	palindromes := make(map[int][][2]int)
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			product := i * j
			if IsPalindrome(product) {
				palindromes[product] = append(palindromes[product], [2]int{i, j}) // [2]int{i, j} is an array literal in Go. Let me break this down:
			}
		}
	}
	if len(palindromes) == 0 {
		return Product{}, Product{}, fmt.Errorf("no palindromes")
	}

	// Find min and max palindromes
	var keys []int
	for k := range palindromes {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	minPalindrome := keys[0]
	maxPalindrome := keys[len(keys)-1]
	return Product{Value: minPalindrome, Factorizations: palindromes[minPalindrome]}, Product{Value: maxPalindrome, Factorizations: palindromes[maxPalindrome]}, nil
}

// NOTE: palindromes[product] - This accesses the slice stored at key product in the map
// If the key doesn't exist, Go returns the zero value for slices, which is nil
// If the key exists, it returns the existing slice
// append(palindromes[product], Pair{i, j}) - This appends a new Pair to the slice
// append can work with nil slices! When you append to nil, it creates a new slice
// Returns a new slice with the element added
