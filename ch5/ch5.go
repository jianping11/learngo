package ch5

import (
	"fmt"
)

//
// func main() {
// 	counterA := createCounter(2)
// 	counterB := createCounter(102)
// 	for i := 0; i < 5; i++ {
// 		a := <-counterA                                    //从通道A接收一个值，并保存到a中
// 		fmt.Printf("(A -> %d, B -> %d) \n", a, <-counterB) //直接将
// 	}
// }

func createCounter(start int) chan int {
	next := make(chan int)
	go func(i int) {
		for {
			next <- i
			i++
		}
	}(start)
	fmt.Println(1)

	return next
}
