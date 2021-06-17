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
	fmt.Println("CheckBrackets    (string) : main.go ----->", result)

	result = myutils.CheckBracketsV2(string(statement))
	fmt.Println("CheckBrackets V2 ([]byte) : main.go ----->", result)
}
