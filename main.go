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
		fmt.Println("USAGE: ./gofk <file-name>")
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
	jumpTable := computeJumptable(commands)

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
			if tape[dp] == 0 {
				ip = jumpTable[ip]
			}
		// ]
		case 93:
			if tape[dp] != 0 {
				ip = jumpTable[ip]
			}
		default:
			fmt.Printf("Error: bad character, ip: %d", ip)
			os.Exit(1)
		}

		ip++
	}
}

func computeJumptable(commands []byte) []int {
	table := make([]int, len(commands))
	ip := 0

	for ip < len(commands) {
		if commands[ip] == 91 {
			nestedBracket := 1
			seek := ip + 1

			for nestedBracket > 0 && seek < len(commands) {
				switch commands[seek] {
				case 91:
					nestedBracket++
				case 93:
					nestedBracket--
				}

				if nestedBracket == 0 {
					table[ip] = seek
					table[seek] = ip
					break
				}

				seek++
			}

			if seek > len(commands) {
				fmt.Println("Error: Mismatched brackets")
				os.Exit(1)
			}
		}
		ip++
	}

	return table
}
