package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"runtime/pprof"
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
	// Get a CPU profile to see where time is spent the most
	// go tool pprof cpu.out
	// top 40, list main.
	f, _ := os.Create("cpu.out")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	tape := make([]byte, tapeSize)
	ip := 0 // instruction pointer
	dp := 0 // data pointer
	input := bufio.NewReader(os.Stdin)

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

					switch commands[ip] {
					case 91:
						nestedBracket++
					case 93:
						nestedBracket--
					}

					// Matched closing bracket, continue
					// main program
					if nestedBracket == 0 {
						ip++
						continue MainLoop
					}
				}

				if ip > len(commands) {
					fmt.Println("Error: Mismatched brackets")
					os.Exit(1)
				}
			}
		// ]
		case 93:
			if tape[dp] != 0 {
				nesting := 1
				for ip > 0 && nesting > 0 {
					ip--

					switch commands[ip] {
					case 93:
						nesting++
					case 91:
						nesting--
					}
				}
			}

		default:
			fmt.Printf("Error: Bad character at ip: %d\n", ip)
			os.Exit(1)
		}

		ip++
	}
}
