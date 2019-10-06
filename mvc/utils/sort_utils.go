package utils

import "sort"

//BubbleSort export order function
func BubbleSort(elements []int) []int {
	keepRunning := true

	for keepRunning {
		keepRunning = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}
	return elements
}

//MySort piola
func MySort(els []int) {
	if len(els) < 1000 {
		BubbleSort(els)
		return
	}
	sort.Ints(els)
}
