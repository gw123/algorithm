package main

import (
	"fmt"
)

/***
  斐波那契数列
  递归解决
 */

func fn(n int) int {
	if (n <= 2) {
		return 1;
	}
	return fn(n-1) + fn(n-2)
}

/***
	动态规划
 */
func fn2(n int) int {
	if (n <= 2) {
		return 1;
	}
	table := make([]int, n+1)
	table[1], table[2] = 1, 1

	for i := 3; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}
	return table[n]
}

func main() {
	//t1 := time.Now()
	//result := fn(49)
	//t2 := time.Now()
	//fmt.Println(result, "used:", t2.Sub(t1))
	//2178309
	//
	res1 := fn2(20)
	result2 := fn2(10) * (fn2(9) + fn2(11))
	fmt.Println(res1, result2)
}
