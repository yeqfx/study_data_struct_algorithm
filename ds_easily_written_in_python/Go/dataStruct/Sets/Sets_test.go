package Sets

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	s := &Set{"foo", "bar"}
	got := s.Contains("foo")
	want := true
	if got != want {
		t.Errorf("foo should contains Set")
	}
	got = s.Contains("baz")
	want = false
	if got != want {
		t.Errorf("baz should not contains Set")
	}
}

func TestInsert(t *testing.T) {
	s := &Set{}
	got := s.Insert("foo")
	want := &Set{"foo"}
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("foo should be inserted in Set")
	}

}
