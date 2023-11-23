package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 打开一个日志文件，如果文件不存在则创建，追加写入
	File1, err := os.OpenFile("example1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错:", err)
	} else {
		defer File1.Close()
		fmt.Println("成功打开文件:", File1.Name())
	}

	dataToWrite := "Hello, Golang!\n"
	_, err1 := File1.WriteString(dataToWrite)
	if err1 != nil {
		fmt.Println("写入文件出错:", err1)
		return
	}
	//fmt.Printf("成功写入 %d 字节到文件\n", bytesWritten)

	// 创建一个带时间戳的写入器
	logWriter := &timestampWriter{timestamp: time.Now(), writer: File1}

	// 模拟用户操作并记录日志
	fmt.Fprintln(logWriter, " 用户登录")
	time.Sleep(2 * time.Second)
	fmt.Fprintln(logWriter, " 用户执行操作A")
	time.Sleep(1 * time.Second)
	fmt.Fprintln(logWriter, " 用户执行操作B")
}

// timestampWriter 是一个实现 io.Writer 接口的结构体，它在写入数据前添加时间戳
type timestampWriter struct {
	timestamp time.Time
	writer    *os.File
}

func (tw *timestampWriter) Write(p []byte) (n int, err error) {
	// 添加时间戳和时间
	tw.timestamp = time.Now()
	timestamp := tw.timestamp.Format("2006-01-02 15:04:05")
	// 输出到文件
	dataToWrite := timestamp + string(p)
	fmt.Printf("%s", dataToWrite)
	dataToWrite += "\n"
	_, err1 := tw.writer.WriteString(dataToWrite)
	if err1 != nil {
		fmt.Println("写入文件出错:", err1)
		return
	}
	return 0, err1
}
