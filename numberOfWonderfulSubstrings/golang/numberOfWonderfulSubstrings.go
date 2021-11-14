package numberOfWonderfulSubstrings

func wonderfulSubstrings(word string) int64 {

	mask := 0
	cnt := make([]int64, 1024)
	cnt[0] = 1
	count := int64(0)
	for i := 0; i < len(word); i++ {
		mask ^= (1 << (word[i] - 'a'))
		count += cnt[mask]
		for j := 0; j < 10; j++ {
			count += cnt[mask^(1<<j)]
		}
		cnt[mask]++
	}
	return count
}
