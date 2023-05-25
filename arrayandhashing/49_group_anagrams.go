package arrayandhashing

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]rune][]string)

	for _, str := range strs {
		var count [26]rune
		for _, char := range str {
			count[char-'a'] += 1
		}
		anagramMap[count] = append(anagramMap[count], str)
	}

	var res [][]string
	for _, val := range anagramMap {
		res = append(res, val)
	}
	return res
}
