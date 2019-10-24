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

const tapeSize = 30000

var commands []byte

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
	commands = r.ReplaceAllLiteral(openFile, []byte(""))
}

func main() {
	tape := make([]byte, tapeSize)
	ip := 0 // instruction pointer
	dp := 0 // data pointer
	input := bufio.NewReader(os.Stdin)
	var s []int // bracket stack

MainLoop:
	for ip < len(commands) {
		c := commands[ip]

		switch c {
		// +
		case 43:
			tape[dp]++
		// -
		case 45:
			tape[dp]--
		// >
		// Loop around if at end
		case 62:
			if dp == tapeSize-1 {
				dp = 0
			} else {
				dp++
			}
		// <
		case 60:
			if dp == 0 {
				dp = tapeSize - 1
			} else {
				dp--
			}
		// .
		case 46:
			fmt.Printf("%c", tape[dp])
		// ,
		case 44:
			c, err := input.ReadByte()
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
				for ip < len(commands) {
					ip++
					newC := commands[ip]

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
