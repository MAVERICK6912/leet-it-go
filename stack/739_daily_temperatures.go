package stack

type temperature struct {
	temp  int
	index int
}

func dailyTemperatures(temperatures []int) []int {
	stack := New[temperature]()
	res := populateDefaultVal(0, make([]int, len(temperatures)))
	for index, temp := range temperatures {
		for stack.Length() > 0 && temp > stack.Peek().temp {
			stackTemp := stack.Pop()
			res[stackTemp.index] = index - stackTemp.index
		}
		stack.Push(temperature{temp: temp, index: index})
	}
	return res
}

func populateDefaultVal[T any](val T, in []T) []T {
	for index := range in {
		in[index] = val
	}
	return in
}
