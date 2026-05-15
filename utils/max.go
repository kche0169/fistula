package utils

func Max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}
