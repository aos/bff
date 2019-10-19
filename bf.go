package main

import "fmt"

// Example bracket configurations
// TODO: Turn this into a test
var examples = map[string]bool{
	"[]":         true,
	"][":         false,
	"[[[][]][]]": true,
	"[[][]]":     true,
	"[[]":        false,
	"][[]":       false,
}

func main() {
	for example, ok := range examples {
		// Bracket parser
		var s []int
		for i, c := range example {
			switch {
			case c == 91:
				s = append(s, i)
			case c == 93 && len(s) > 0:
				// Pop from slice
				_, s = s[len(s)-1], s[:len(s)-1]
			}
		}

		if (len(s) == 0) != ok {
			fmt.Println(example)
		}
	}
}
