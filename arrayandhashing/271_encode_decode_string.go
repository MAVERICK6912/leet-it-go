package arrayandhashing

import (
	"fmt"
	"strconv"
)

const delimiter byte = '~'

func encode(strs []string) string {
	res := ""
	for _, str := range strs {
		res += fmt.Sprintf("%d%s%s", len(str), string(delimiter), str)
	}
	return res
}

func decode(str string) []string {
	var (
		lenOfWord   int
		lengthFound bool
		length      []byte
		err         error
	)
	res := make([]string, 0)
	for i := 0; i < len(str); i++ {
		if str[i] == delimiter {
			lenOfWord, err = strconv.Atoi(string(length))
			if err != nil {
				continue
			}
			lengthFound = true
			length = make([]byte, 0)
		}
		if lengthFound {
			res = append(res, str[i+1:i+lenOfWord+1])
			i += lenOfWord + 1
			lengthFound = false
		}
		if i < len(str) {
			if str[i] >= 48 && str[i] <= 57 {
				length = append(length, str[i])
			}
		}
	}
	return res
}
