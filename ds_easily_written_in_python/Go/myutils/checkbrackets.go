package myutils

import (
	slicestack "../dataStruct/SliceStack"
)

func CheckBrackets(statement string) bool {
	stack := slicestack.NewSliceStack()
	for _, ch := range statement {
		if ch == '{' || ch == '[' || ch == '(' {
			stack.Push(ch)
		} else if ch == '}' || ch == ']' || ch == ')' {
			if stack.IsEmpty() {
				return false
			} else {
				left := stack.Pop()
				if (ch == '}' && left != '{') || (ch == ']' && left != '[') || (ch == ')' && left != '(') {
					return false
				}
			}
		}
	}
	return stack.IsEmpty()
}

func CheckBracketsV2(statement string) bool {
	stack := slicestack.NewSliceStack()
	for ch := range statement {
		if ch == '{' || ch == '[' || ch == '(' {
			stack.Push(ch)
		} else if ch == '}' || ch == ']' || ch == ')' {
			if stack.IsEmpty() {
				return false
			} else {
				left := stack.Pop()
				if (ch == '}' && left != '{') || (ch == ']' && left != '[') || (ch == ')' && left != '(') {
					return false
				}
			}
		}
	}
	return stack.IsEmpty()
}
