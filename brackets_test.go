package main

import (
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

var tests = []struct {
	configuration string
	correct       bool
}{
	{"[]", true},
	{"][", false},
	{"[[[][]][]]", true},
	{"[[][]]", true},
	{"[[]", false},
	{"][[]", false},
}

func TestBracketParser(t *testing.T) {
	t.Log("Given the need to teset different bracket configurations...")
	{
		for num, tt := range tests {
			t.Logf("\tTest %d:\t%v", num, tt.configuration)
			{
				// Example bracket parser
				var s []int
				for i, c := range tt.configuration {
					switch {
					case c == 91:
						s = append(s, i)
					case c == 93 && len(s) > 0:
						_, s = s[len(s)-1], s[:len(s)-1]
					}
				}

				if (len(s) == 0) != tt.correct {
					t.Errorf("\t%s should be %v.", failed, tt.correct)
				} else {
					t.Logf("\t%s should be %v.", succeed, tt.correct)
				}

			}
		}
	}
}
