package main

import (
	"fmt"
	"log"
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

const helloWorld = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

func main() {
	tape := make([]byte, 30000)
	ip := 0
	dp := 0

	var s []int // bracket stack
	skipLoop := false

	for ip < len(helloWorld) {
		c := helloWorld[ip]

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
		case 91:
			if tape[dp] <= 0 {
				skipLoop = true
			}

			s = append(s, ip)
		// ]
		case 93:
			if len(s) <= 0 {
				log.Fatalf("ERROR: Mismatched brackets, ip: %d", ip)
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
		log.Fatal("ERROR: Mismatched brackets")
	}
}
