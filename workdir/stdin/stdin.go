/*
 * Christian Visintin - Studying Go <https://github.com/ChristianVisintin/studying-go>
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		fmt.Println(stdin.Text())
	}
}
