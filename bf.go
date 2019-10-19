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
//	.	46
//	,	44
//

const example = ">><-++].,"

func main() {
	var s []int // bracket stack

	for i, c := range example {
		switch c {
		// +
		case 43:
			fmt.Println("+")
		// -
		case 45:
			fmt.Println("-")
		// >
		case 62:
			fmt.Println(">")
		// <
		case 60:
			fmt.Println("<")
		// .
		case 46:
			fmt.Println(".")
		// ,
		case 44:
			fmt.Println(",")
		// [
		case 91:
			s = append(s, i)
		// ]
		case 93:
			if len(s) <= 0 {
				log.Fatalf("ERROR: Mismatched brackets, index: %d", i)
			}

			// Pop from slice
			_, s = s[len(s)-1], s[:len(s)-1]
		}
	}

	if len(s) > 0 {
		log.Fatal("ERROR: Mismatched brackets")
	}
}
