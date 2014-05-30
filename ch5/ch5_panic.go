package main

import (
	"fmt"
)

func main() {
	hello()
}

func hello() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Hello world, this func while panic")
	panic("test panic")
}
