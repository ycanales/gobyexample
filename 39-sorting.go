package main

import (
	"fmt"
	"sort"
)

func main() {

	// in-place sorting
	strs := []string{"c", "a", "b", "ñ", "A", "B", "C", "n",
		"!", "3", "1"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// is a slice of Ints already sorted?
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}