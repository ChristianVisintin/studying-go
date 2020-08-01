/*
 * Christian Visintin - Studying Go <https://github.com/ChristianVisintin/studying-go>
 */

package main

import (
	"fmt"
	"strconv"
)

// Typedef
type UNIXEpoch int64

// Enums in go; NOTE: iota increment starting from 0 by 1 for each entry
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// @! Struct
type Person struct {
	uuid    string
	Name    string
	Surname string
	Address // Anonymous field
	age     uint
}

type Address struct {
	Street      string
	CivicNumber int
	City        string
	Country     string
}

func main() {

	// Simple variable
	var simpleVar int = 56
	fmt.Println(simpleVar)

	// Numeric values
	var byteVal uint8 = 127
	fmt.Println(byteVal)
	// Float values
	var f float32 = 1.82
	fmt.Println(f)

	// There are imaginary numbers too, but I don't think I'm ever going to use them...

	// Boolean
	var booleanVal bool = true
	fmt.Println(booleanVal)

	// Strings
	var str string = "Hello world!"
	fmt.Println(str)
	// String slices
	fmt.Println(str[:5])             // Prints 'Hello'
	fmt.Println(str[6 : len(str)-1]) // Prints 'world'

	// Swap values
	var i int = 78
	var j int = 12
	i, j = j, i
	fmt.Println(i, j)

	// Pointers
	var ptr *int = &simpleVar
	*ptr = 48
	fmt.Println(simpleVar)

	// Tuple assignment
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// @! Type conversion

	// Number types casting

	var uintInt uint16 = uint16(i)
	fmt.Println(uintInt)

	// String to numbers

	str = "128"
	var numStr, err = strconv.Atoi(str)
	if err != nil {
		fmt.Println("Could not convert", str, "to number")
	} else {
		fmt.Println(str, "=>", numStr)
	}

	// Numbers to string

	var myNum int = 65535
	str = strconv.Itoa(myNum)
	fmt.Println(myNum, "=>", str)

	// @! Composite types

	// Arrays NOTE: fixed-length sequence of zero or more elements of a particular type
	var array [3]int = [3]int{1, 2, 3}
	fmt.Println(array)
	// Iterate over array
	for idx, val := range array {
		fmt.Printf("%d: %d\n", idx, val)
	}
	// Array declaration can also be simplified using this syntax
	var q = [...]int{1, 2, 3, 4}
	fmt.Println(q)

	// Slices: variabler length sequences whose elementts all have the same type; are portions of arrays
	var months = [...]string{"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}
	var holidayMonths = months[6:8]
	for _, val := range holidayMonths {
		fmt.Println(val)
	}

	// Maps
	var countryPrefix = make(map[string]int) // Key is string, value is integer
	countryPrefix["IT"] = 39
	countryPrefix["FR"] = 33
	countryPrefix["DE"] = 49
	countryPrefix["US"] = 1
	// Delete element
	delete(countryPrefix, "US")
	// Iterate map
	for country, prefix := range countryPrefix {
		fmt.Printf("%s: %d\n", country, prefix)
	}
	// If key doesn't exist, return 0 value
	fmt.Println(countryPrefix["SK"])
	// This method can be used instead
	if prefix, exists := countryPrefix["SK"]; exists {
		fmt.Println(prefix)
	} else {
		fmt.Println("Key 'SK' doesn't exist in countryPrefix")
	}

	// @! Structs (see declaration at 29)
	var omar Person
	omar.uuid = "d8bceaf5-fbf9-4ca8-bef7-9adbb21f7347"
	omar.Name = "Omar"
	omar.Surname = "De la Lune"
	omar.Address = Address{"Rue de Lagny", 67, "Paris", "FR"}
	omar.age = 32

	// Access to fields and anonymous field
	fmt.Printf("%s lives at %d, %s, %s (%s)\n", omar.Name, omar.CivicNumber, omar.Street, omar.City, omar.Country)

	var jane Person
	jane.uuid = "a2dbde73-6598-4833-b59d-dfea4ff550e7"
	jane.Name = "Jane"
	jane.Surname = "Doe"
	jane.age = 24
	jane.Address = Address{"Baker Street", 221, "London", "UK"}

	// Struct are comparable!
	fmt.Println(omar == jane)

}
