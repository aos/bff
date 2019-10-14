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
		var s []rune
		for _, c := range example {
			if c == 91 {
				s = append(s, c)
				continue
			}

			if c == 93 && len(s) > 0 && s[len(s)-1] == 91 {
				// Pop from slice
				_, s = s[len(s)-1], s[:len(s)-1]
				continue
			}
		}

		if (len(s) == 0) != ok {
			fmt.Println(example)
		}
	}
}
