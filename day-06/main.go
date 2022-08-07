package main

import "fmt"

func main() {

	//	input := []int{1, 1, 1, 1, 2, 1, 1, 4, 1, 4, 3, 1, 1, 1, 1, 1, 1, 1, 1, 4,
	//		1, 3, 1, 1, 1, 5, 1, 3, 1, 4, 1, 2, 1, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	//		3, 4, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 4, 1, 1, 1, 1, 3, 5, 1, 1, 2,
	//		1, 1, 1, 1, 4, 4, 1, 1, 1, 4, 1, 1, 4, 2, 4, 4, 5, 1, 1, 1, 1, 2, 3, 1, 1,
	//		4, 1, 5, 1, 1, 1, 3, 1, 1, 1, 1, 5, 5, 1, 2, 2, 2, 2, 1, 1, 2, 1, 1, 1, 1,
	//		1, 3, 1, 1, 1, 2, 3, 1, 5, 1, 1, 1, 2, 2, 1, 1, 1, 1, 1, 3, 2, 1, 1, 1, 4,
	//		3, 1, 1, 4, 1, 5, 4, 1, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 4, 5, 1, 1,
	//		1, 1, 5, 4, 1, 3, 1, 1, 1, 1, 4, 3, 3, 3, 1, 2, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	//		2, 1, 1, 1, 5, 1, 3, 1, 4, 3, 1, 3, 1, 5, 1, 1, 1, 1, 3, 1, 5, 1, 2, 4, 1,
	//		1, 4, 1, 4, 4, 2, 1, 2, 1, 3, 3, 1, 4, 4, 1, 1, 3, 4, 1, 1, 1, 2, 5, 2, 5,
	//		1, 1, 1, 4, 1, 1, 1, 1, 1, 1, 3, 1, 5, 1, 2, 1, 1, 1, 1, 1, 4, 4, 1, 1, 1,
	//		5, 1, 1, 5, 1, 2, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 2, 4, 1, 1,
	//		2, 1, 1, 3, 2}

	input := []int{3, 4, 3, 1, 2}

	m := make(map[int]int, 9)

	for i := 0; i < 9; i++ {
		m[i] = 0
	}
	for _, n := range input {
		m[n]++
	}

	var result int

	for i := 0; i < 1; i++ {
		m = nextDay(m)
		result = countFishes(m)
	}

	fmt.Println(m)

	fmt.Println(result)
}

func countFishes(m map[int]int) (r int) {
	for _, v := range m {
		r += v
	}
	return
}

// nextDay
func nextDay(m map[int]int) map[int]int {
	out := make(map[int]int, 9)
	for i := 0; i < 9; i++ {
		out[i] = 0
	}

	out[6] += m[0]
	out[8] += m[0]

	for i := 1; i < 9; i++ {
		out[i-1] += m[i]
	}

	return out

}