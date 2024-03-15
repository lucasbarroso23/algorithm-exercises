package twoSums

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

func twoSum1(arr []int, value int) []int {
	res := []int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if (arr[i]+arr[j]) == value && i != j {
				res = []int{arr[i], arr[j]}
				return res
			}
		}
	}

	return res
}

func twoSum2(arr []int, value int) []int {
	for i := 0; i < len(arr); i++ {
		diff := value - arr[i]
		for j := 0; j < len(arr); j++ {
			if arr[j] == diff {
				return []int{arr[i], arr[j]}
			}
		}
	}

	return nil
}
