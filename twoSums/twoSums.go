package twoSums

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
