package main

import "fmt"

func insertSort(list []int) {
	count := 0
	for i := 1; i < len(list); i++ {
		for j := i; j > 0 && list[j] < list[j-1]; j-- {
			list[j], list[j-1] = list[j-1], list[j]
			count++
		}
	}
	fmt.Println("insertSort交换次数", count)
}

func shellSort(list []int) {
	count := 0
	length := len(list)
	for gap := length / 2; gap >= 1; gap = gap / 2 {
		for i := gap; i < length; i += gap {
			for j := i; j > 0 && list[j] < list[j-gap]; j -= gap {
				list[j], list[j-gap] = list[j-gap], list[j]
				count++
			}
		}
	}
	fmt.Println("shellSort交换次数", count)
}

func insertSortTest() {
	items := []int{92, 28, 8, 2, 16, 75, 9, 48, 1, 63, 3, 78, 7, 9, 49, 22, 4, 12, 11, 20, 10, 14, 18, 31, 22, 32, 40, 50, 39, 42, 46, 98, 78, 93, 58, 84, 93, 85, 67, 57, 58, 38}
	insertSort(items)
	fmt.Println(items)
}

func shellSortTest() {
	items := []int{92, 28, 8, 2, 16, 75, 9, 48, 1, 63, 3, 78, 7, 9, 49, 22, 4, 12, 11, 20, 10, 14, 18, 31, 22, 32, 40, 50, 39, 42, 46, 98, 78, 93, 58, 84, 93, 85, 67, 57, 58, 38}
	shellSort(items)
	fmt.Println(items)
}

func main() {
	insertSortTest()
	shellSortTest()
}
