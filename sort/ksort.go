package main

import "fmt"


func KsortInt(items []int, left, right int) {
	key := items[left]
	pos := left
	if left < right {
		i := left
		j := right

		for i < j {
			for items[j] >= key && i < j {
				j--
			}
			items[pos] = items[j]
			pos = j
			for items[i] <= key && i < j {
				i++
			}
			items[pos] = items[i]
			pos = i
		}
		items[i] = key
		KsortInt(items, left, i-1)
		KsortInt(items, i+1, right)
	}
}

func main() {
	items := []int{8, 2, 9, 1, 3, 7, 22}
	KsortInt(items, 0, len(items)-1)
	fmt.Println(items)
}
