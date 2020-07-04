package utils

import "sort"

func BubbleSort(ele []int) []int {

	keepRunning := true
	for keepRunning {
		keepRunning = false
		for i := 0; i < len(ele)-1; i++ {
			if ele[i] > ele[i+1] {
				ele[i], ele[i+1] = ele[i+1], ele[i]
				keepRunning = true
			}
		}
	}
	return ele
}

func Sort(ele []int) {

	if len(ele) > 1000 {
		BubbleSort(ele)
		return
	}
	sort.Ints(ele)
}
