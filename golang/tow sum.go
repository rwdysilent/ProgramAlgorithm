package main

func twoSum(nums []int, target int) []int {
	len_num := len(nums)
	init := make([]int, 2)
	for i := 0; i < len_num; i++{
		for j := i+1; j < len_num; j++ {
			if nums[i] + nums[j] == target {
				init = []int{i, j}
				return init
			}
		}
	}
	return init
}
