package main

func twoSum(nums []int, target int) []int {
	len_num := len(nums)
	for i := 0; i < len_num; i++{
		for j := i+1; j < len_num; j++ {
			if nums[i] + nums[j] == target {
				return i, j
			}
		}
	}
	return
}
