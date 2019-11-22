

package main
import "fmt"

func sort(a []int) []int {
	count := len(a)
	for i := 1; i < count; i++ {
		currIndex := i
		for j := i - 1; j >= 0; j-- {
			if a[currIndex] < a[j] {
				a[currIndex], a[j] = a[j], a[currIndex]
				currIndex = j
			} else {
				break
			}
		}
	}
	return a
}

func shellSort(a []int, step int) []int {
	count := len(a)
	for i := step; i < count; i += step {
		currIndex := i
		for j := i - step; j >= 0; j-- {
			if a[currIndex] < a[j] {
				a[currIndex], a[j] = a[j], a[currIndex]
				currIndex = j
			} else {
				break
			}
		}
	}
	return a
}
func sSort(a []int) []int {
	steps := []int {8, 4, 2, 1}
	for index := 0; index < len(steps); index++ {
		a = shellSort(a, steps[index])
	}
	return a
}

func main() {
	a := []int{3, 1, 59, 31, 16, 22, 18, 93, 123, 76, 86, 12, 32, 23}
	fmt.Println(sort(a))
	fmt.Println(sSort(a))
}