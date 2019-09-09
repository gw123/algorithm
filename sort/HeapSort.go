package main

import "fmt"

func HeapSort(arr []int) {
	length := len(arr)
	startNode := length/2 -1
	for k := startNode; k >= 0; k-- {
		buildHeap(arr, k)
	}

	for lastNodePos := length - 1; lastNodePos > 0; lastNodePos -- {
		swap(arr, 0, lastNodePos)
		buildHeap(arr[0:lastNodePos], 0)
	}
}

func buildHeap(arr []int, currentPos int) {
	length := len(arr)
/***
结束条件
1. 当前节点是没有左右子节点,也就是当前节点是叶子节点
2. 当前节点大于左右节点. 这里需要结合HeapSort方法的第一个
循环在buildHeap过程中 采用从最深的的非叶子节点构造,所以除了
当前节点他的左右节点都已经是最大堆了,所以当当前节点大于左右
节点的没有必要在继续循环.
 */
	for currentPos*2+1 < length {
		leftChildPos := currentPos*2 + 1
		rightChildPos := leftChildPos + 1
		maxPos := leftChildPos
		// 选择左右子节点中最大的节点,若果没有右节点maxPos就是 左节点
		if rightChildPos < length {
			if arr[leftChildPos] < arr[rightChildPos] {
				maxPos = rightChildPos
			}
		}

		if arr[maxPos] > arr[currentPos] {
			swap(arr, maxPos, currentPos)
			currentPos = maxPos
		} else {
			break
		}
	}
}

func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	items := []int{8, 2, 9, 1, 3, 7, 22}
	HeapSort(items)
	fmt.Println(items)
}
