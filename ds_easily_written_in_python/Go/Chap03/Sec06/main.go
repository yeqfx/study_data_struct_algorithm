package main

import (
	"../../dataStruct/Sets"
)

func main() {
	setA := Sets.NewSet()
	setA.Insert("휴대폰")
	setA.Insert("지갑")
	setA.Insert("손수건")
	setA.Display("Set A : ")

	setB := Sets.NewSet()
	setB.Insert("빗")
	setB.Insert("파이썬 자료구조")
	setB.Insert("야구공")
	setB.Insert("지갑")
	setB.Display("Set B : ")

	setB.Insert("빗")
	setA.Delete("손수건")
	setA.Delete("발수건")
	setA.Display("Set A : ")
	setB.Display("Set B : ")

	setA.Union(setB).Display("A U B : ")
	setA.Intersect(setB).Display("A ^ B : ")
	setA.Difference(setB).Display("A - B : ")
}
