package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

// TAPESIZE is the size of our tape
const TAPESIZE = 30000

// Program is our brainfuck program holding the instructions
type Program struct {
	instructions []byte
}

// ParseProgram takes in the bf filename and returns the program removing
// all non-supported characters (everything except +-><.,[])
func ParseProgram(filename string) *Program {
	openFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Parse out instructions
	// This is the lazy person's approach ;-)
	r := regexp.MustCompile("[^\\[\\]\\+-\\.,><]")
	return &Program{
		r.ReplaceAllLiteral(openFile, []byte("")),
	}
}

// Interpret takes in the program and interprets it
// This is where the magic happens
func Interpret(p *Program, TRACE bool) {
	instCount := make(map[byte]int) // Used for tracing only
	tape := make([]byte, TAPESIZE)
	ip := 0 // instruction pointer
	dp := 0 // data pointer
	input := bufio.NewReader(os.Stdin)
	jumpTable := computeJumptable(p.instructions)

	for ip < len(p.instructions) {
		c := p.instructions[ip]

		if TRACE {
			instCount[c]++
		}

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
			if dp == TAPESIZE-1 {
				dp = 0
			} else {
				dp++
			}
		// <
		case 60:
			if dp == 0 {
				dp = TAPESIZE - 1
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

	if TRACE {
		total := 0
		fmt.Printf("\n*** Tracing activated, printing instruction count: ***\n")
		for k, v := range instCount {
			total += v
			fmt.Printf("%c  --  %d\n", k, v)
		}
		fmt.Println("-----")
		fmt.Printf("TOTAL: %d\n", total)
	}
}

func computeJumptable(instructions []byte) []int {
	table := make([]int, len(instructions))
	ip := 0

	for ip < len(instructions) {
		if instructions[ip] == 91 {
			nestedBracket := 1
			seek := ip + 1

			for nestedBracket > 0 && seek < len(instructions) {
				switch instructions[seek] {
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

			if seek > len(instructions) {
				fmt.Println("Error: Mismatched brackets")
				os.Exit(1)
			}
		}
		ip++
	}

	return table
}
