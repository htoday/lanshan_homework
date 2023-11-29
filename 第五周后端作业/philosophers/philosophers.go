package main

import (
	"fmt"
	"time"
)

func Eat(id int) {
	//for {
	personNum <- struct{}{}
	m[id-1] <- 1 //占用筷子
	m[id%5] <- 1
	fmt.Println(id, "号哲学家吃了")
	_ = <-m[id-1] //释放筷子
	_ = <-m[id%5]
	_ = <-personNum
	Think()
	//}

}
func Think() {
	time.Sleep(time.Second)
}

var m [6]chan int
var personNum chan struct{}

func main() {
	personNum = make(chan struct{}, 4) //缓冲量为4
	for i := 0; i <= 4; i++ {
		m[i] = make(chan int, 1)
	}
	for i := 1; i <= 5; i++ {
		go Eat(i)
	}
	time.Sleep(time.Second * 3)
}
