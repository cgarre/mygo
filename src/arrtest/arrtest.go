package main

import "fmt"

func printDupes(a []int) {
	//using map
	/*var m map[int]int
	m = make(map[int]int)
	for _, s := range a {
		m[s]++
	}

	for k, v := range m {
		fmt.Println(k, " = ", v)
	}*/

	N := len(a)

	for i := 0; i < N; i++ {
		currVal := a[i] % N
		a[currVal] += N
	}

	for i := 0; i < N; i++ {
		times := a[i] / N
		if times > 0 {
			fmt.Println(i, " appeared ", times, " time(s).")
		}
		a[i] = a[i] % N
	}

}

func printArr(a []int) {
	for _, s := range a {
		fmt.Println(s)
	}
}

func main() {

	a := []int{1, 2, 2}

	printDupes(a)
	printArr(a)
}
