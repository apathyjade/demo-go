

package main

import "fmt"

func selectSort(a []int) []int {
	l := len(a)
	for i := 0; i < l - 1; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if a[min] > a[j] {
				a[min], a[j] = a[j], a[min]
			}
		}
		if min != i {
			min = i
		}
	}
	return a
}

func getMaxNode(a []int, node int, l int) {
	left := 2 * node + 1
	right := 2 * node + 2
	temp := node
	if (l >= left && a[left] > a[temp]) {
		temp = left
	}
	if (l >= right && a[right] > a[temp]) {
		temp = right
	}
	if node != temp {
	 a[node], a[temp] = a[temp], a[node]
	}
}
func heapSort(a []int) []int {
	l := len(a)
	for i := l - 1; i > 0; i-- {
		var b int = len(a) / 2 - 1
		for j := b; j >= 0; j-- {
			getMaxNode(a, j, i)
		}
		a[0], a[i] = a[i], a[0]
	}
	return a
}

func main() {
	a := []int{3, 1, 59, 31, 16, 22, 18, 93, 123, 76, 86, 12, 32, 23}
	fmt.Println(selectSort(a))
	fmt.Println(heapSort(a))
}