package main

// Given an array of integers, return indices of the two numbers such that they add up to a specific
// target.
// You may assume that each input would have exactly one solution, and you may not use the same
// element twice.

//example:
//Given nums = [2, 7, 11, 15], target = 9,
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].

func twoSum(nums []int, target int) []int {
	len_num := len(nums)
	init := make([]int, 2)
	for i := 0; i < len_num; i++ {
		for j := i + 1; j < len_num; j++ {
			if nums[i]+nums[j] == target {
				init = []int{i, j}
				return init
			}
		}
	}
	return init
}

func main() {
	twoSum()
}
