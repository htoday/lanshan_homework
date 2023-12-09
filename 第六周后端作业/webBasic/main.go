package main

import (
	"fmt"
	"webBasic/api"
	"webBasic/dao"
)

func main() {
	dao.ReadMessagesFromFile()
	err := dao.LoadFormJSON("dataHome")
	if err != nil {
		fmt.Println("读取错误：", err)
	}
	defer dao.SaveToJSON("dataHome", dao.Database) //数据保存
	defer func() {
		for v, k := range dao.Database {
			fmt.Println("username:", k, "password", v)
		}
	}() //不知道写了有什么用的defer。。。因为并不会执行。。。
	api.InitRouter()
}
