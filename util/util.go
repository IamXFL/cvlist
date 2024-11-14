package util

func generatePermutations(n int) []string {
	var result []string
	var path []byte

	var backtrack func(int)
	backtrack = func(index int) {
		if index == n {
			result = append(result, string(path))
			return
		}

		for i := 0; i < 26; i++ {
			path[index] = byte('a' + i)
			backtrack(index + 1)
		}
	}

	path = make([]byte, n)
	backtrack(0)
	return result
}
