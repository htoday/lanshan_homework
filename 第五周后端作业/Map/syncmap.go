package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type YourMap struct {
	data    sync.Map
	waiting map[int]chan string
}

func (m *YourMap) Get(k int, maxWaitingTime time.Duration) (string, error) {
	val, ok := m.data.Load(k)
	if !ok {
		ch := make(chan string)
		m.waiting[k] = ch
		select {
		case rec := <-ch:
			//delete(m.waiting, k)
			return rec, nil
		case <-time.After(maxWaitingTime):
			return "", errors.New("超时了")
		}
	}
	return val.(string), nil
}

func (m *YourMap) Put(k int, v string) {
	m.data.Store(k, v)
	if ch, ok := m.waiting[k]; ok {
		ch <- v
		close(ch) // 关闭通道，因为此时等待的 goroutine 已经得到了数据
	}
}

func main() {
	myMap := YourMap{
		waiting: make(map[int]chan string),
	}

	go myMap.Put(2, "崩铁")
	go func() {
		s, err := myMap.Get(1, time.Second)
		if err == nil {
			fmt.Println("找到", s)
		} else {
			fmt.Println(err)
		}
	}()
	go func() {
		s, err := myMap.Get(2, time.Second)
		if err == nil {
			fmt.Println("找到", s)
		} else {
			fmt.Println(err)
		}
	}()
	go func() {
		s, err := myMap.Get(3, time.Second)
		if err == nil {
			fmt.Println("找到", s)
		} else {
			fmt.Println(err)
		}
	}()
	go myMap.Put(1, "原神")
	time.Sleep(time.Second * 2)
}
