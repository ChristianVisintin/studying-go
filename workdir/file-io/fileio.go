/*
 * Christian Visintin - Studying Go <https://github.com/ChristianVisintin/studying-go>
 */

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	OptRead    = iota
	OptWrite   = iota
	OptUnknown = iota
)

func read_file(filename string) bool {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error while reading file %s: %s\n", filename, err)
		return false
	}
	fmt.Println(string(data))
	return true
}

func write_file(filename string) {
	var data string
	var stdin = bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		data += stdin.Text() + "\n"
	}
	ioutil.WriteFile(filename, []byte(data), 0644)
}

func main() {

	var command int = OptUnknown
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <READ/WRITE> <filename>\n", os.Args[0])
		os.Exit(255)
	}
	switch os.Args[1] {
	case "READ":
		command = OptRead
	case "WRITE":
		command = OptWrite
	default:
		fmt.Printf("Unknown command '%s'\n", os.Args[1])
		os.Exit(255)
	}
	var filename string = os.Args[2]
	// Perform read
	if command == OptRead {
		if !read_file(filename) {
			os.Exit(1)
		}
	} else {
		// Write
		write_file(filename)
	}

	os.Exit(0)
}
