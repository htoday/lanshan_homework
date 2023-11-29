package main

import (
	"fmt"
	"time"
)

func FindMax() int {
	Max := 0
	for i := 1; i <= n; i++ {
		Max = max(Max, number[i])
	}
	return Max
}
func Lock(id int) {
	choosing[id] = true //进入选号状态
	number[id] = FindMax() + 1
	choosing[id] = false
	for i := 1; i <= n; i++ {
		for choosing[i] == true {
		} //如果其他进程在选号，等待
		for (number[i] != 0) && (number[i] < number[id]) || ((number[i] == number[id]) && (i < id)) {
		} //如果看到其他进程的号码比自己小（靠前）就等待
	}
}
func Unlock(id int) {
	number[id] = 0
}
func routine(id int) {

	Lock(id)
	for i := 1; i <= 100000; i++ {
		sum++
	}
	fmt.Println(id, "执行完毕")
	Unlock(id)

}

var choosing [1000]bool
var number [1000]int
var sum, n int //n= number of all routines

func main() {
	n = 4
	for {
		sum = 0
		go routine(1)
		go routine(2)
		go routine(3)
		go routine(4)
		//go routine(4)
		//go routine(5)
		time.Sleep(time.Second)
		fmt.Println(sum)
	}
}
