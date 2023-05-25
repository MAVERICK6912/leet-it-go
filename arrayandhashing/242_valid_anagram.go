package arrayandhashing

func isAnagram(s string, t string) bool {
	charMap := make(map[rune]int)
	for _, r := range s {
		charMap[r] += 1
	}
	for _, r := range t {
		if _, ok := charMap[r]; !ok {
			return false
		} else {
			charMap[r] -= 1
		}
	}
	for _, count := range charMap {
		if count > 0 {
			return false
		}
	}
	return true
}
