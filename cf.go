package main

import (
	"fmt"
)

func iterMask(n int) []int {
	i := n
	ti := n
	arr :=  []int {}
	for i >= 0{
		fmt.Println(i,n, i&n)
		if n & i == ti {
			arr = append(arr, i)
		}
		i--
	}
	return arr
}

func countSumOfTwoRepresentations2(n int, l int, r int) int {
	cnt := 0
	a := l
	b := r
	if (a == b) {
		cnt = 1
	}
	for a != r{
		if a + r == n {
			cnt ++
		}
		a++
	}
	for b != l{
		if b + l == n {
			cnt ++
		}
		b--
	}
	return cnt
}


func main (){

	fmt.Println(countSumOfTwoRepresentations2(6,2,4))

}
