package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
func copyFolderContents(src, dst string) error {
	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// 遍历文件列表并复制到目标文件夹
	for _, file := range files {
		srcFile := filepath.Join(src, file.Name()) //等同于srcFile := src + "/" + file.Name()
		dstFile := filepath.Join(dst, file.Name())

		if file.IsDir() { //如果文件是文件夹
			if _, err := os.Stat(dstFile); os.IsNotExist(err) { //判断文件夹是否存在，如果不存在就创建文件夹
				if err := os.MkdirAll(dstFile, os.ModePerm); err != nil {
					return err
				}
			}
			//利用递归复制文件夹
			if err := copyFolderContents(srcFile, dstFile); err != nil {
				return err
			}
		} else { //如果不是文件夹而是文件就复制
			if err := copyFile(srcFile, dstFile); err != nil {
				return err
			}
		}
	}
	return nil
}
func GetUpdateTime(FileName string) (time.Time, error) {
	UpdatedFile, err := os.Stat(FileName)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return UpdatedFile.ModTime(), err
}
func main() {
	FilePath1 := "example1"
	FilePath2 := "example2"

	InitialModtime1, err := GetUpdateTime(FilePath1)
	if err != nil {
		return
	}
	InitialModtime2, err := GetUpdateTime(FilePath2)
	if err != nil {
		return
	}

	ticker := time.NewTicker(5 * time.Second) //每五秒一次周期

	for range ticker.C { //C通道发送值就执行一次循环

		UpdatedModtime1, err := GetUpdateTime(FilePath1)
		if err != nil {
			return
		}

		if UpdatedModtime1 != InitialModtime1 {
			fmt.Println("检测到", FilePath1, "文件被修改")
			fmt.Println("正在将修改同步到", FilePath2)
			err = copyFolderContents(FilePath1, FilePath2)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("同步成功")
			InitialModtime1 = UpdatedModtime1
			InitialModtime2, err = GetUpdateTime(FilePath2)
			if err != nil {
				return
			}
		} else {
			fmt.Println("文件", FilePath1, "未被修改")
		}

		UpdatedModtime2, err := GetUpdateTime(FilePath2)
		if err != nil {
			return
		}
		if UpdatedModtime2 != InitialModtime2 {
			fmt.Println("检测到", FilePath2, "文件被修改")
			fmt.Println("正在将修改同步到", FilePath1)
			err = copyFolderContents(FilePath2, FilePath1)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("同步成功")
			InitialModtime2 = UpdatedModtime2
			InitialModtime1, err = GetUpdateTime(FilePath2)
			if err != nil {
				return
			}
		} else {
			fmt.Println("文件", FilePath2, "未被修改")
		}
	}
}
