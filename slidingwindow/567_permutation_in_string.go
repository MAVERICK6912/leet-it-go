package slidingwindow

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	countS1, countS2 := make([]int, 26), make([]int, 26)
	fillArrayWithDefaultValue(0, countS1)
	fillArrayWithDefaultValue(0, countS2)

	for i := 0; i < len(s1); i++ {
		countS1[int(s1[i]-'a')] += 1
		countS2[int(s2[i]-'a')] += 1
	}
	matches := 0
	for i := 0; i < 26; i++ {
		if countS1[i] == countS2[i] {
			matches += 1
		}
	}
	left := 0
	for right := len(s1); right < len(s2); right++ {
		if matches == 26 {
			return true
		}
		index := int(s2[right] - 'a')
		countS2[index] += 1
		if countS1[index] == countS2[index] {
			matches += 1
		} else if countS1[index]+1 == countS2[index] {
			matches -= 1
		}

		index = int(s2[left] - 'a')
		countS2[index] -= 1
		if countS1[index] == countS2[index] {
			matches += 1
		} else if countS1[index]-1 == countS2[index] {
			matches -= 1
		}
		left += 1

	}
	return matches == 26
}

func fillArrayWithDefaultValue[T any](v T, arr []T) []T {
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return arr
}
