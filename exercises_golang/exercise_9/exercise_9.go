package exercise9

import (
	"fmt"
	"sort"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func Exercise9(words string) (string, error) {
	if len(words) == 0 {
		return "", fmt.Errorf("Input string cannot be empty")
	}
	wordsSlice := SplitAndSort(words)
	wordsNew := strings.Join(wordsSlice, " ")
	return wordsNew, nil
}

func SplitAndSort(words string) []string {
	wordsSlice := strings.Fields(words)
	wordsSet := mapset.NewSet(wordsSlice...)
	wordsCleanSlice := wordsSet.ToSlice()
	sort.Strings(wordsCleanSlice)
	return wordsCleanSlice
}
