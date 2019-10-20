package main

import (
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

	// Remove all non-command characters
	r := regexp.MustCompile("[^\\[\\]\\+-\\.,><]")
	file = r.ReplaceAllLiteral(openFile, []byte(""))
}

func main() {
	tape := make([]byte, 30000)
	ip := 0
	dp := 0

	var s []int // bracket stack
	skipLoop := false

	for ip < len(file) {
		c := file[ip]

		if skipLoop {
			ip++
			continue
		}

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
			fmt.Println(",")
		// [
		// TODO (BUG): I need to capture each loop construct in its own
		// struct and assign skipping instructions based on that. If we
		// encounter a "[" with a 0 value, and then encounter another
		// "[", followed by a "]", it will stop skipping. This is
		// incorrect behavior as the original outer loop should
		// determine skipping behavior
		case 91:
			if tape[dp] <= 0 {
				skipLoop = true
			}

			s = append(s, ip)
		// ]
		case 93:
			if len(s) <= 0 {
				fmt.Printf("ERROR: Mismatched brackets, ip: %d", ip)
				os.Exit(1)
			}

			var temp int
			// Pop from stack
			temp, s = s[len(s)-1], s[:len(s)-1]

			if tape[dp] <= 0 {
				skipLoop = false
				ip++
				continue
			}

			if !skipLoop {
				ip = temp
				continue
			}
		}

		ip++
	}

	if len(s) > 0 {
		fmt.Println("ERROR: Mismatched brackets")
		os.Exit(1)
	}
}
