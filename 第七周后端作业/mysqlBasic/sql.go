package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 定义一个全局对象db
var db *sqlx.DB

func initDB() {
	var err error

	dsn := "root:1234@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}

	fmt.Println("connecting to MySQL...")
	return
}

// 插入数据
func insertRowDemo(name1 string, age1 int, sex1 string) {
	sqlStr := "insert into student(sex,name, age) values (?,?,?)"
	ret, err := db.Exec(sqlStr, sex1, name1, age1)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}
func readStudents() error {
	// 执行查询语句
	rows, err := db.Query("SELECT * FROM student")
	if err != nil {
		return fmt.Errorf("查询数据失败: %v", err)
	}

	// 定义变量用于存储查询结果
	var id int
	var sex sql.NullString //不然遇到NULL会报错
	var name string
	var age int

	// 遍历查询结果并读取数据
	for rows.Next() {
		sex2 := "NULL"
		if sex.Valid {
			sex2 = sex.String
		}
		err := rows.Scan(&id, &sex, &name, &age)
		if err != nil {
			return fmt.Errorf("读取数据失败: %v", err)
		}
		fmt.Printf("ID: %d, Sex: %s, Name: %s, Age: %d\n", id, sex2, name, age)
	}

	return nil
}
func main() {
	initDB()
	//insertRowDemo("CHJ", 14, "boy")
	//insertRowDemo("MHY", 5, "girl")
	//insertRowDemo("WYF", 56, "boy")
	//insertRowDemo("WDNMD", 250, "girl")
	err := readStudents()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
