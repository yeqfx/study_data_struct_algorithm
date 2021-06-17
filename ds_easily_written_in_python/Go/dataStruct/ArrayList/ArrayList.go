package ArrayList

import "fmt"

type Arraylist []interface{}

func NewArrayList() *Arraylist {
	return &Arraylist{}
}

func (a *Arraylist) Insert(pos int, elem interface{}) *Arraylist {
	result := append(*a, elem)
	copy(result[pos+1:], result[pos:])
	result[pos] = elem
	return &result
}

func (a *Arraylist) Delete(pos int) *Arraylist {
	*a = append((*a)[:pos], (*a)[pos+1:]...)
	return a
}

func (a *Arraylist) IsEmpty() bool {
	return len(*a) == 0
}

func (a *Arraylist) GetEntry(pos int) interface{} {
	return (*a)[pos]
}

func (a *Arraylist) Size() int {
	return len(*a)
}

func (a *Arraylist) Clear() *Arraylist {
	*a = (*a)[:0]
	return a
}

func (a *Arraylist) Find(item interface{}) int {
	for i, v := range *a {
		if v == item {
			return i
		}
	}
	return -1
}

func (a *Arraylist) Replace(pos int, elem interface{}) {
	(*a)[pos] = elem
}

func (a *Arraylist) Len() int      { return len(*a) }
func (a *Arraylist) Swap(i, j int) { (*a)[i], (*a)[j] = (*a)[j], (*a)[i] }
func (a *Arraylist) Less(i, j int) bool {
	str_i := fmt.Sprintf("%v", (*a)[i])
	str_j := fmt.Sprintf("%v", (*a)[j])
	return str_i < str_j
	// return fmt.Sprint((*a)[i]) < fmt.Sprint((*a)[j]) }
}

func (a *Arraylist) Merge(lst []interface{}) {
	*a = append(*a, lst...)
}

func (a *Arraylist) Display(msg string) {
	fmt.Println(msg, a.Size(), *a)
}
