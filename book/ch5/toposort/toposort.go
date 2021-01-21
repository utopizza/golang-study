package ch5

import (
	"fmt"
	"sort"
)

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	// anonymous function recursion
	// must first declare a variable
	// then assign the function to that variable
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func main() {
	var prereqs = map[string][]string{
		"algorithms":      {"data structures"},
		"calculus":        {"linear algebra"},
		"data structures": {"discrete math"},
	}
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
