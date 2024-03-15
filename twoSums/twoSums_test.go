package twoSums

import (
	"testing"
)

func TestTwoSums(t *testing.T) {
	t.Run("twoSums1", func(t *testing.T) {
		data := []int{3, 4, 1, 5}
		value := 8

		result := twoSum1(data, value)
		t.Logf("two sums result: %+v", result)
	})

	t.Run("twoSums2", func(t *testing.T) {
		data := []int{3, 4, 1, 5}
		value := 8

		result := twoSum2(data, value)
		t.Logf("two sums result: %+v", result)
	})
}
