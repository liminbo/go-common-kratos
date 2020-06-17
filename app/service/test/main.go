package main

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	go f1(out)
	time.Sleep(1 * time.Second)
	out <- 2
	time.Sleep(2 * time.Second)
}
