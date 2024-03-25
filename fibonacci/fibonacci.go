package fibonacci

import "fmt"

/*
The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence,
such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
Given n, calculate F(n).



Example 1:

Input: n = 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
*/

func fibonacci1() func() int {
	x, y := 0, 1

	return func() int {
		f := x
		x, y = y, f+y
		return f
	}
}

func fibonacci2(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci2(n-1) + fibonacci2(n-2)
}

func fibonacciGoRoutine(c chan int) {
	n := cap(c)
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacciGoRoutineSelect(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
