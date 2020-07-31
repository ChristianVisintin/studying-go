/*
 * Christian Visintin - Studying Go <https://github.com/ChristianVisintin/studying-go>
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <url>\n", os.Args[0])
		os.Exit(255)
	}

	var url string = os.Args[1]
	// Perform GET
	var resp, err = http.Get(url)
	if err != nil {
		fmt.Printf("HTTP error: %s\n", err)
		os.Exit(1)
	}

	if resp.StatusCode == 200 {
		fmt.Println("OK")
	} else {
		fmt.Printf("ERR %s\n", resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read response: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(data))

	os.Exit(0)
}
