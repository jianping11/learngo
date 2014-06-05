package main

import (
	"fmt"
	"math"
)

func main() {
	// xs := []int{2, 3, 4, 5, 6, 7, 8, 9}
	// ys := []string{"a", "b", "张", "dd"}
	// fmt.Println(SliceIndex(len(xs), func(i int) bool { return xs[i] == 8 }))
	// fmt.Println(SliceIndex(len(xs), func(i int) bool { return xs[i] == 2 }))
	// fmt.Println(SliceIndex(len(xs), func(i int) bool { return xs[i] == 100 }))
	// fmt.Println(SliceIndex(len(ys), func(i int) bool { return ys[i] == "dd" }))
	// fmt.Println(SliceIndex(len(ys), func(i int) bool { return ys[i] == "张" }))

	fmt.Println(SliceIndex(math.MaxInt32, func(i int) bool {
		// if i != 0 && i%5 == 0 {
		// 	fmt.Printf("now i = %d \n", i)
		// 	if i%5 == 0 {
		// 	}
		// } else {
		// 	fmt.Printf("now i = %d \t", i)
		// 	if i%5 == 0 {
		// 	}
		// }

		return i > 0 && i%27 == 0 && i%51 == 0
	}))

	readings := []int{4, -3, 2, -7, 8, 19, -11, 7, 18, -6}
	even := make([]int, 0, len(readings))
	SliceFilter(len(readings),
		func(i int) bool {
			return readings[i]%2 == 0
		},
		func(i int) {
			even = append(even, readings[i])
		})
	fmt.Println(even)

	parts := []string{"X15", "T14", "X23", "A41", "L19", "X57", "A63"}
	var xparts []string
	SliceFilter(len(parts), func(i int) bool { return parts[i][0] == 'X' },
		func(i int) { xparts = append(xparts, parts[i]) })
	fmt.Println("xparts", xparts)
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

// func IntFilter(slice []int, predicate func(int) bool) []int {
// 	filtered := make([]int, 0, len(slice))
// 	for i := 0; i < len(slice); i++ {
// 		if predicate(slice[i]) {
// 			filtered = append(filtered, slice[i])
// 		}
// 	}
//
// 	return filtered
// }

func SliceFilter(limit int, predicate func(int) bool, appender func(int)) {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			appender(i)
		}
	}
}
