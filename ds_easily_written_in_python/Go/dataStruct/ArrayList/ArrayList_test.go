package ArrayList

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	a := &Arraylist{}
	want := &Arraylist{10}
	got := a.Insert(0, 10)
	fmt.Println("want : ", want)
	fmt.Println("got : ", got)
	if got == want {
		t.Errorf("Insert(0, 10) = %v; want &[10]", got)
	}
}
