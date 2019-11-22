

package main
import "fmt"

func sort(a []int) []int {
	l := len(a)
	for i := 0; i < l - 1; i++ {
		for j := 0; j < l - i - 1; j++ {
			if a[j] > a[j + 1] {
				a[j], a[j + 1] = a[j + 1], a[j]
			}
		}
	}
	return a
}
func quick(a []int, begin int, end int) []int {
	if end - begin <= 1 {
		return a
	}
	mid := begin
	for i := begin + 1; i <= end; i++ {
		if a[mid] > a[i] {
			if mid + 1 > i {
				a[mid + 1], a[i] = a[i], a[mid + 1]
			}
			mid++
		}
	}
	quick(a, begin, mid)
	quick(a, mid + 1, end)
	return a
}

func main() {
	a := []int{3, 1, 59, 31, 16, 22, 18, 93, 123, 76, 86, 12, 32, 23}
	fmt.Println(sort(a))
	fmt.Println(quick(a, 0, len(a) - 1))
}