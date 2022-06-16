// WorldCount
// The code below should return a map of the counts of each “word” in the string s.
// Input: "Hello, my name is John"
// Output: map[string]int{"Hello": 1, ",": 1, "my": 1, "name": 1, "is": 1, "John": 1}

package main

import (
	"strings"
)

var m map[string]int

func WordCount(s string) map[string]int {
	worlds := strings.Fields(s)

	m = make(map[string]int)

	for _, value := range worlds {
		count, ok := m[value]
		if ok {
			m[value] = count + 1
		} else {
			m[value] = 1
		}
	}

	return m
}
