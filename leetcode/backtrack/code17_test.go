package backtrack

func letterCombinations(digits string) []string {
	if digits == "" || digits == "1" || digits == "*" || digits == "0" || digits == "#" {
		return nil
	}
	sets := map[string]string {
		"2":"abc",
		"3":"def",
		"4":"ghi",
		"5":"jkl",
		"6":"mno",
		"7":"pqrs",
		"8":"tuv",
		"9":"wxyz",
	}
	k := len(digits)
	oneResult := make([]byte, 0)
	results := make([]string, 0)
	var letterCombinationsCore func([]byte, int)
	letterCombinationsCore = func(in []byte, index int) {
		if len(oneResult) == k {
			results = append(results, string(oneResult))
			return
		}
		digitsBytes := []byte(sets[string(in[index])])
		for i := 0; i < len(digitsBytes); i ++ {
			oneResult = append(oneResult, digitsBytes[i])
			letterCombinationsCore(in, index + 1)
			// greedy
			oneResult = oneResult[:len(oneResult) - 1]
		}
	}
	letterCombinationsCore([]byte(digits), 0)
	return results
}
