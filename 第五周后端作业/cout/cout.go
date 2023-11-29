package main

import (
	"fmt"
	"time"
)

func Odd() {
	for {
		rec, ok := <-ch
		if ok == false {
			break
		}
		rec++
		fmt.Println(rec, "Odd")
		ch <- rec
	}

}
func Even() {
	for {
		rec := <-ch
		if rec > 100 {
			close(ch)
			break
		}
		rec++
		fmt.Println(rec, "Even")
		ch <- rec
	}

}

var ch chan int

func main() {
	ch = make(chan int, 1)
	go Odd()
	go Even()
	ch <- 0
	time.Sleep(time.Second)
}
