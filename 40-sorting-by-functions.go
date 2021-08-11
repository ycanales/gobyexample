package main

import (
	"fmt"
	"sort"
)

// To sort by a custom function we need a 
// corresponding type. This one is just an
// alias for the []string builtin.
type byLength []string

// Implement sort.Interface: Len, Less and Swap.
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	// Sort by ascending order
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// Convert original fruits slice to "byLength"
	sort.Sort(byLength(fruits))
	
	fmt.Println(fruits)
}