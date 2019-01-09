package problem0055

func canJump(nums []int) bool {
	length := len(nums)
	maxPos := 0
	i := 0

	for ; i <= maxPos && i < length; i++ {
		maxPos = max(maxPos, nums[i] + i);
	}

	return i == length
}

func max(a, b int) int {
	if (a >= b) {
		return a;
	}

	return b;
}
