package main

import (
	"sort"

	"../../dataStruct/ArrayList"
)

func main() {
	var s = &ArrayList.Arraylist{}
	s.Display("Golang Interface 슬라이스로 구현한 ArrayList 테스트")
	// a = make(Arraylist, 10)
	s = s.Insert(0, 10)
	s = s.Insert(0, 20)
	s = s.Insert(1, 30)
	s = s.Insert(s.Size(), 40)
	s = s.Insert(2, 50)
	s.Display("Golang Interface 슬라이스로 구현한 ArrayList(삽입x5)")
	sort.Sort(s)
	s.Display("Golang Interface 슬라이스로 구현한 ArrayList(정렬후)")
	s.Replace(2, 90)
	s.Display("Golang Interface 슬라이스로 구현한 ArrayList(교체x1)")
	s.Delete(2)
	s.Delete(s.Size() - 1)
	s.Delete(0)
	s.Display("Golang Interface 슬라이스로 구현한 ArrayList(삭제x3)")
	lst := []interface{}{1, 2, 3}
	s.Merge(lst)
	s.Display("Golang Interface 슬라이스로 구현한 ArrayList(병합x3)")
	s.Clear()
	s.Display("Golang Interface 슬라이스로 구현한 ArrayList(정리후)")
}
