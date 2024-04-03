package removeduplicates

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Run("Remove Duplicates from sorted array", func(t *testing.T) {
		input := []int{1, 1, 2}
		numberElems := removeDuplicates(input)

		assert.Equal(t, numberElems, 2)

		input2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		numberElems2 := removeDuplicates(input2)

		assert.Equal(t, numberElems2, 5)
	})
}
