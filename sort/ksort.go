package main

import (
	"fmt"
	"time"
)

func sort(items []interface{}, compore func(a, b interface{}) uint8) {

}

/***
   10,2,9,1,3,7,22
 */

type Compare func(a, b interface{}) uint8

func Ksort(items []interface{}, compore Compare) {
	key := items[0]
	left := 0
	right := len(items) - 1
	fmt.Print(items, "\t")
	fmt.Println("left ", left, "\tright ", right, "\t", items[left:right+1])
	time.Sleep(time.Second)
	if left < right {
		i := left
		j := right
		pos := left
		for i < j {
			for compore(items[j], key) >= 1 && i < j {
				j--
			}

			if j > pos {
				items[pos] = items[j]
				pos = j
			}

			for compore(items[i], key) <= 1 && i < j {
				i++
			}

			if i < pos {
				items[pos] = items[i]
				pos = i
			}
		}

		items[pos] = key
		if left < i-1 {
			Ksort(items[left:i], compore)
		}
		if right > i+1 {
			Ksort(items[i+1:right+1], compore)
		}
	}
}

func KsortInt(items []int, left, right int) {
	key := items[left]
	fmt.Print(items, "\t")
	fmt.Println("left ", left, "\tright ", right, "\t", items[left:right+1])
	time.Sleep(time.Second)
	if left < right {
		i := left
		j := right
		pos := left
		for i < j {
			for items[j] >= key && i < j {
				j--
			}

			if j > pos {
				items[pos] = items[j]
				pos = j
			}

			for items[i] <= key && i < j {
				i++
			}

			if i < pos {
				items[pos] = items[i]
				pos = i
			}
		}

		items[pos] = key
		if left < i-1 {
			KsortInt(items, left, i-1)
		}
		if right > i+1 {
			KsortInt(items, i+1, right)
		}
	}
}

func sortInt(items []int) {
	arrayLen := len(items)
	KsortInt(items, 0, arrayLen-1)

}

func main() {
	items := []int{8, 2, 9, 1, 3, 7, 22, 4, 12, 11}
	sortInt(items)
	fmt.Println(items)
	//items := make([]int_type,0)
	//Ksort( items , func(a, b interface{}) uint8 {
	//	aa := a.(int)
	//	bb := b.(int)
	//	if aa > bb {
	//		return 1
	//	} else if aa == bb {
	//		return 0
	//	} else {
	//		return -1
	//	}
	//})
	//KsortInt(items, 0, len(items)-1)
}
