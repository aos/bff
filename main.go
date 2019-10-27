package main

import (
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

var program *Program

func init() {
	if len(os.Args) <= 1 || len(os.Args) > 2 {
		fmt.Println("USAGE: ./bff <file-name>")
		os.Exit(1)
	}

	program = ParseProgram(os.Args[1])
}

func main() {
	Interpret(program, TRACEBF)
}
