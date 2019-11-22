

package main
import "fmt"

func mergeSort(a []int) []int {
	sort(a, 0, len(a) - 1)
	return a
}
func sort(a []int, l int, r int) {
	if (l >= r) {
		return
	}
	mid :=  int (l + r) / 2
	sort(a, l, mid)
	sort(a, mid + 1, r)
	merge(a, l, mid, r)
}
func merge (a []int, l int, mid int, r int) {
	tmp := []int{}
	currL := l
	currR := mid + 1
	for ; currL <= mid && currR <= r; {
		if a[currL] <= a[currR] {
			tmp = append(tmp, a[currL])
			currL++
		} else {
			tmp = append(tmp, a[currR])
			currR++
		}
	}

	if currL <= mid {
		for i := currL; i <= mid; i++ {
			tmp = append(tmp, a[i])
		}
	}
	if  currR <= r {
		for i := currR; i <= r; i++ {
			tmp = append(tmp, a[i])
		}
	}
	for i := 0; i < len(tmp); i++ {
		a[l + i] = tmp[i]
	}
}

func main() {
	a := []int{3, 1, 59, 31, 16, 22, 18, 93, 123, 76, 86, 12, 32, 23}
	fmt.Println(mergeSort(a))
}