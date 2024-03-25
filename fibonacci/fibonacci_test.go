package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	t.Run("fibonacci1", func(t *testing.T) {
		f := fibonacci1()
		for i := 0; i < 10; i++ {
			fmt.Print(f())
		}
	})

	t.Run("fibonacci2", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			fmt.Print(fibonacci2(i))
		}

	})

	t.Run("fibonacciGoRoutine", func(t *testing.T) {
		c := make(chan int, 10)
		go fibonacciGoRoutine(c)
		for i := range c {
			fmt.Print(i)
		}
	})

	t.Run("fibonacciGoRoutineSelect", func(t *testing.T) {
		c := make(chan int)
		quit := make(chan int)

		go func() {
			for i := 0; i < 10; i++ {
				fmt.Print(<-c)
			}
			quit <- 0
		}()
		fibonacciGoRoutineSelect(c, quit)
	})
}
