package main

import (
	"fmt"

	slicestack "../../dataStruct/SliceStack"
)

func main() {

	odd := slicestack.NewSliceStack()
	even := slicestack.NewSliceStack()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even.Push(i)
		} else {
			odd.Push(i)
		}
	}
	fmt.Println(" 스택 even push 5회 : ", even)
	fmt.Println(" 스택 odd  push 5회 : ", odd)
	fmt.Println(" 스택 even     peek : ", even.Peek())
	fmt.Println(" 스택 odd      peek : ", odd.Peek())
	for i := 0; i < 2; i++ {
		even.Pop()
	}
	for i := 0; i < 3; i++ {
		odd.Pop()
	}
	fmt.Println(" 스택 even pop  2회 : ", even)
	fmt.Println(" 스택 odd  pop  3회 : ", odd)
}
