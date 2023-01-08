package main

import (
	"fmt"

	"github.com/iods/go-util/debug"
)

func main() {
	debug.TestOne()
	// fmt.Println("this is just a test function. 2")
	testTwo("number two.\n")
	debug.TestThree()
}

func testTwo(v string) {
	fmt.Printf("this is a string with %s", v)
}
