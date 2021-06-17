package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"../../myutils"
)

func main() {
	statement, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatal(err)
	}
	result := myutils.CheckBrackets(string(statement))
	fmt.Println("main.go ----->", result)
}
