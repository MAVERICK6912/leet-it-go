package stack

type stack[T any] []T

func New[T any]() stack[T] {
	return make(stack[T], 0)
}

func (s *stack[T]) Push(n ...T) {
	for _, v := range n {
		*s = append(*s, v)
	}
}

func (s *stack[T]) Pop() T {
	var res T
	if len(*s) == 0 {
		return res
	}
	res = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *stack[T]) Peek() T {
	var res T
	if len(*s) > 0 {
		return (*s)[len(*s)-1]
	}
	return res
}

func (s *stack[T]) GetAll() []T {
	res := make([]T, 0)
	for _, v := range *s {
		res = append(res, v)
	}
	return res
}

func (s stack[T]) Length() int {
	return len(s)
}
