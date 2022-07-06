package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// write a  new file from byte string
	name := "test.txt"

	txt := []byte("Not much in this file.")

	if err := ioutil.WriteFile(name, txt, 0755); err != nil {
		panic(err)
	}
	// Read the contents of the file
	results, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(results))

	// Or use os.Open(filename)

	reader, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	results, err = ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	reader.Close()
	fmt.Println(string(results))

}
