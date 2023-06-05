package stack

import "sort"

type car struct {
	position float64
	speed    float64
}

func carFleet(target int, position []int, speed []int) int {
	cars := populateFleet(position, speed)

	sort.Slice(cars, func(i, j int) bool { return cars[i].position < cars[j].position })

	stack := NewStack[float64]()

	for i := len(cars) - 1; i >= 0; i-- {
		time := (float64(target) - cars[i].position) / cars[i].speed
		stack.Push(time)
		if stack.Length() >= 2 && stack.Get(1) <= stack.Get(2) {
			stack.Pop()
		}
	}
	return stack.Length()
}

func populateFleet(position, speed []int) []car {
	fleets := make([]car, 0)
	for index := range position {
		fleets = append(fleets, car{position: float64(position[index]), speed: float64(speed[index])})
	}
	return fleets
}
