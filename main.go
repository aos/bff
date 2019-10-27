package main

import (
	"flag"
	"fmt"
	"os"
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

// TAPESIZE is the size of our tape
const TAPESIZE = 30000

var flagTrace bool
var program *Program

func init() {
	flag.BoolVar(&flagTrace,
		"trace",
		false,
		"Traces the number of instructions",
	)
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 || len(args) > 1 {
		fmt.Println("USAGE: ./gofk [--trace] <file-name>")
		os.Exit(1)
	}

	program = ParseProgram(args[0])
}

func main() {
	Interpret(program)
}
