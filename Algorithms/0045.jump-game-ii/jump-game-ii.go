package problem0045

func jump(nums []int) int {
	cur := 0
	last := 0
	ret := 0

	for i := 0; i < len(nums); i++ {
		if (i > last) {
			last = cur
			ret++
		}

		cur = max(cur, nums[i] + i)
	}

	return ret
}

func max(a, b int) int {
	if (a > b) {
		return a
	}

	return b
}