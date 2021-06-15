package Sets

import "fmt"

type Set []interface{}

func (s *Set) Size() int {
	return len(*s)
}

func (s *Set) Display(msg string) {
	fmt.Println(msg, s.Size(), *s)
}

func (s *Set) Contains(item interface{}) bool {
	for _, v := range *s {
		if v == item {
			return true
		}
	}
	return false
}

func (s *Set) Insert(elem interface{}) *Set {
	if !s.Contains(elem) {
		*s = append(*s, elem)
	}
	return s
}

func (s *Set) Delete(elem interface{}) *Set {
	if s.Contains(elem) {
		for i := 0; i < len(*s); i++ {
			if (*s)[i] == fmt.Sprint(elem) {
				*s = append((*s)[:i], (*s)[i+1:]...)
			}
		}
	}
	return s
}

func (s *Set) Union(setB *Set) *Set {
	setC := make(Set, len(*s))
	for i, v := range *s {
		setC[i] = v
	}
	for _, elem := range *setB {
		if !s.Contains(elem) {
			setC = append(setC, elem)
		}
	}
	return &setC
}

func (s *Set) Intersect(setB *Set) *Set {
	setC := Set{}
	for _, elem := range *setB {
		if s.Contains(elem) {
			setC = append(setC, elem)
		}
	}
	return &setC
}

func (s *Set) Difference(setB *Set) *Set {
	setC := Set{}
	for _, elem := range *s {
		if !setB.Contains(elem) {
			setC = append(setC, elem)
		}
	}
	return &setC
}
