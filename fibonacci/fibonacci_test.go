package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	t.Run("fibonacci1", func(t *testing.T) {
		f := fibonacci1()
		for i := 0; i < 40; i++ {
			fmt.Print(f())
		}
	})

	t.Run("fibonacci2", func(t *testing.T) {
		for i := 0; i < 40; i++ {
			fmt.Print(fibonacci2(i))
		}

	})
}
