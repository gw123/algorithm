package main

import "fmt"

/***
	n 当前递归盘子数量
	x 当前递归的初始盘子所在的柱子
	y 当前递归要借用的柱子
	z 要移动到的柱子
 */
func hannuota(n int, x, y, z byte) {
	//将最后一个盘子上面的n-1盘子借助z移动到y
	if (n == 1) {
		fmt.Printf("将第%d个盘子从%c->%c\n", n, x, z)
		return
	}
	hannuota(n-1, x, z, y)
	//最下面的正好可以移动到z
	fmt.Printf("将第%d个盘子从%c->%c\n", n, x, z)
	//将当前在y上面的n-1盘子借助x移动到z
	hannuota(n-1, y, x, z)
}

func main() {
	hannuota(4, 'x', 'y', 'z')
}
