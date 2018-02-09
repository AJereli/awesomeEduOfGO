/*
 * Copyright (c) 2018. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package Trash

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


//func main (){
//
//	fmt.Println(countSumOfTwoRepresentations2(6,2,4))
//
//}
