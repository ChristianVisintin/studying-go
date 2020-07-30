/*
 * Christian Visintin - Studying Go <https://github.com/ChristianVisintin/studying-go>
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Let's introduce C like printf
	fmt.Printf("You provided %d arguments\n", len(os.Args))
	var argsString, sep string
	// Iterate over arguments
	for i := 1; i < len(os.Args); i++ {
		argsString += sep + os.Args[i]
		sep = " "
	}
	// Print arguments in a single string
	fmt.Printf("Program name: %s\n", os.Args[0])
	fmt.Printf("CLI options: %s\n", argsString)
	// We could rewrite this using Join obviously
	//NOTE: here we have a slice too
	argsString = strings.Join(os.Args[1:], " ")
	fmt.Printf("Joined CLI options: %s\n", argsString)
}
