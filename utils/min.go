package utils

func Min[T Number](nums ...T) T {
	if len(nums) == 0 {
		var zero T
		return zero
	}
	max := nums[0]
	for _, n := range nums[1:] {
		if n < max {

			max = n
		}
	}
	return max
}
