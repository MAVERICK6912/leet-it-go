package stack

type Histogram struct {
	height int
	index  int
}

func largestRectangleArea(heights []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	heightStack := NewStack[Histogram]()
	maxArea := 0
	for index, height := range heights {
		start := index
		for heightStack.Length() > 0 && heightStack.Peek().height > height {
			hist := heightStack.Pop()
			maxArea = max(maxArea, hist.height*(index-hist.index))
			start = hist.index
		}
		heightStack.Push(Histogram{height: height, index: start})
	}
	if heightStack.Length() > 0 {
		for heightStack.Length() > 0 {
			hist := heightStack.Pop()
			maxArea = max(maxArea, hist.height*(len(heights)-hist.index))
		}
	}
	return maxArea
}
