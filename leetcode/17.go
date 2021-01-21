package leetcode

func letterCombinations(digits string) []string {
	keyBoard := map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	if digits == "" {
		return []string{}
	}
	results := []string{""}
	for _, key := range digits {
		letters := keyBoard[key]
		var tempArr []string
		for _, letter := range letters {
			for _, str := range results {
				tempArr = append(tempArr, str+string(letter))
			}
		}
		results = tempArr
	}
	return results
}
