package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	//out := make(chan int)
	//go f1(out)
	//time.Sleep(1 * time.Second)
	//out <- 2
	//time.Sleep(2 * time.Second)

	var bo int = 111111111111111
	fmt.Println("data:%s", bo)
}
