/*
 * Christian Visintin - Studying Go <https://github.com/ChristianVisintin/studying-go>
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readFile(filename string) (string, bool) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error while reading file %s: %s\n", filename, err)
		return err.Error(), true
	}
	return string(data), false
}

func queryJSON(jsonObj map[string]interface{}, query []string) (string, bool) {
	var node string = query[0]

	// Get node
	if found, exists := jsonObj[node]; exists {
		// Base case
		if len(query) == 1 {
			// Return value
			return fmt.Sprint(found), true
		}
		// Recursive case
		// Pop from query
		query = query[1:]
		return queryJSON(found.(map[string]interface{}), query)
	}
	return "", false
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <jsonfile> <jsonQuery>\n", os.Args[0])
		os.Exit(255)
	}

	// Parse json query
	var queryArray = strings.Split(os.Args[2], ".")
	var query []string = queryArray

	// Read file
	var jsonFile string = os.Args[1]
	jsonData, err := readFile(jsonFile)
	if err {
		fmt.Println("Could not read JSON file:", jsonData)
		os.Exit(1)
	}
	var jsonMap map[string]interface{}
	var jsonErr = json.Unmarshal([]byte(jsonData), &jsonMap)
	if jsonErr != nil {
		fmt.Println("Could not decode JSON file:", jsonErr.Error())
		os.Exit(1)
	}
	// Query
	var result, found = queryJSON(jsonMap, query)
	if found {
		fmt.Println(result)
	} else {
		fmt.Println("null")
		os.Exit(1)
	}
}
