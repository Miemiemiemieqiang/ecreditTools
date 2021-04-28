package main

import (
	"fmt"
	"time"
)

func main() {
	//ch := make(chan int, 100)
	//
	//go provider(0, ch)
	//go provider(100, ch)
	//go consumer(ch, 0)
	//
	//cha := make(chan int)
	//select {
	//case <-cha:
	//case <-time.After(10 * time.Second):
	//	fmt.Println("time out")
	//
	//}
	fmt.Println("Done")
}

func provider(start int, ch chan int) {
	for i := start; i < start+100; i++ {
		fmt.Printf("add %d into ch\n", i)
		ch <- i
		time.Sleep(10 * time.Millisecond)
	}
}

func consumer(ch chan int, index int) {
	for i := range ch {
		fmt.Printf("%d : ch out %d\n", index, i)
	}
}
