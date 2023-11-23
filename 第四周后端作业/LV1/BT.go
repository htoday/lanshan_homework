package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	newFile1, err := os.Create("example1.txt")
	if err != nil {
		fmt.Println("创建文件出错:", err)
	} else {
		defer newFile1.Close()
		fmt.Println("成功创建文件:", newFile1.Name())
	}

	newFile2, err := os.Create("example2.txt")
	if err != nil {
		fmt.Println("创建文件出错:", err)
	} else {
		defer newFile2.Close()
		fmt.Println("成功创建文件:", newFile2.Name())
	}

	// 打开文件（使用 OpenFile）
	File1, err := os.OpenFile("example1.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件出错:", err)
	} else {
		defer File1.Close()
		fmt.Println("成功打开文件:", File1.Name())
	}
	File2, err := os.OpenFile("example2.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件出错:", err)
	} else {
		defer File2.Close()
		fmt.Println("成功打开文件:", File2.Name())
	}
	now1 := time.Now()
	// 写入文件
	for i := 1; i <= 10000; i++ {
		dataToWrite := "Hello, Golang!\n"
		_, err := File1.WriteString(dataToWrite)
		if err != nil {
			fmt.Println("写入文件出错:", err)
			return
		}
		//fmt.Printf("成功写入 %d 字节到文件\n", bytesWritten)
	}
	WriteTime1 := time.Since(now1)

	//用bufio写入文件

	now2 := time.Now()

	writer := bufio.NewWriter(File2)
	defer writer.Flush() // 确保在程序结束前将缓冲区中的数据写入文件

	for i := 1; i <= 10000; i++ {
		dataToWrite := "Hello, Golang!\n"
		_, err := writer.WriteString(dataToWrite)
		if err != nil {
			fmt.Println("写入文件出错:", err)
			return
		}
		//fmt.Printf("成功写入 %d 字节到文件\n", bytesWritten)
	}
	WriteTime2 := time.Since(now2)
	fmt.Println(" 缓冲写入花费了：", WriteTime2)
	fmt.Println(" 不缓冲写入花费了：", WriteTime1)
}
