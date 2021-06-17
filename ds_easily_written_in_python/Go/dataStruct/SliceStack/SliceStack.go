package slicestack

type SliceStack []interface{}

func NewSliceStack() *SliceStack {
	return &SliceStack{}
}

func (s *SliceStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *SliceStack) Size() int {
	return len(*s)
}

func (s *SliceStack) Clear() {
	*s = (*s)[:0]
}

func (s *SliceStack) Push(item interface{}) *SliceStack {
	*s = append(*s, item)
	return s
}

func (s *SliceStack) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	popValue := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popValue
}

func (s *SliceStack) Peek() interface{} {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}
