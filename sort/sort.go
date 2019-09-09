package main

import "fmt"

func select_sort(items []int) {
	length := len(items)
	for i := 0; i < length-1; i++ {
		pos := i
		for j := i + 1; j < length-1; j++ {
			if items[j] < items[pos] {
				pos = j
			}
		}
		if pos != i {
			items[pos], items[i] = items[i], items[pos]
		}
	}
}

func bubble(items []int) {
	length := len(items)

	for i := 0; i < length-1; i++ {
		flag := false
		for j := length - 1; j > i; j-- {
			if items[j] < items[j-1] {
				items[j], items[j-1] = items[j-1], items[j]
				flag = true
			}
		}
		if flag == false {
			return
		}
	}
}

func bubble1(items []int) {
	length := len(items)
	for i := 0; i < length-1; i++ {
		flag := false
		for j := 0; j<length -1; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
				flag = true
			}
		}
		if flag == false {
			return
		}
	}
}

func test1() {
	items := []int{8, 2, 9, 4, 8, 1, 7, 22}
	select_sort(items)
	fmt.Println(items)
}

func test2() {
	items := []int{8, 2, 9, 4, 8, 1, 7, 22}
	bubble1(items)
	fmt.Println(items)
}

func main() {
	test2()
}
