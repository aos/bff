package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

// The main loop is a switch/case statement parsing each character by its
// character code. The character codes are:
//
//	+	43
//	-	45
//	>	62
//	<	60
//	[	91
//	]	93
//	.	46 - output
//	,	44 - input
//

var file []byte

func init() {
	if len(os.Args) <= 1 || len(os.Args) > 2 {
		fmt.Println("USAGE: ./bf <file-name>")
		os.Exit(0)
	}

	openFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Parse out commands
	// This is the lazy person's approach ;-)
	r := regexp.MustCompile("[^\\[\\]\\+-\\.,><]")
	file = r.ReplaceAllLiteral(openFile, []byte(""))
}

func main() {
	tape := make([]byte, 30000)
	ip := 0
	dp := 0

	var s []int // bracket stack

MainLoop:
	for ip < len(file) {
		c := file[ip]

		switch c {
		// +
		case 43:
			tape[dp]++
		// -
		case 45:
			tape[dp]--
		// >
		case 62:
			dp++
		// <
		case 60:
			dp--
		// .
		case 46:
			fmt.Print(string(tape[dp]))
		// ,
		case 44:
			in := bufio.NewReader(os.Stdin)
			c, err := in.ReadByte()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			tape[dp] = c
		// [
		case 91:
			// We ignore everything coming after the bracket, until
			// we hit the matching closing bracket.
			if tape[dp] <= 0 {
				nestedBracket := 1
				for ip < len(file) {
					ip++
					newC := file[ip]

					switch newC {
					case 91:
						nestedBracket++
					case 93:
						nestedBracket--
					}

					if nestedBracket == 0 {
						ip++
						continue MainLoop
					}
				}
			} else {
				// Push the current position of the bracket
				s = append(s, ip)
			}

		// ]
		case 93:
			if len(s) <= 0 {
				fmt.Printf("ERROR: Mismatched brackets, ip: %d", ip)
				os.Exit(1)
			}

			// Pop from stack and go back to last
			ip, s = s[len(s)-1], s[:len(s)-1]
			continue
		}

		ip++
	}

	if len(s) > 0 {
		fmt.Println("ERROR: Mismatched brackets")
		os.Exit(1)
	}
}
