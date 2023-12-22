package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var Rdb *redis.Client
var db *sqlx.DB

func AddLike(id string, passageID string) error {
	ctx := context.Background()
	err := Rdb.SAdd(ctx, passageID, id).Err()
	if err != nil {
		return err
	}
	return nil
}
func GetFromID(id int) (interface{}, interface{}, error) {
	ctx := context.Background()
	userID := strconv.Itoa(id)
	userData, err := Rdb.HMGet(ctx, userID, "username", "password").Result()
	if err != nil {
		println("Failed to find userData in redis")
		return "", "", err
	}
	if userData[0] == nil && userData[1] == nil {
		println("Not find in redis")

		sqlstr := "SELECT * FROM users where id = ?"
		var user1 User
		err = db.Get(&user1, sqlstr, id)
		if err != nil {
			return "", "", err
		}
		//存入redis
		err = Rdb.HMSet(ctx, userID, "username", user1.Username, "password", user1.Password).Err()
		if err != nil {
			return user1.Username, user1.Password, err
		}
		//设置过期时间
		/*err = Rdb.Expire(ctx, userID, 1*time.Hour).Err()
		if err != nil {
			fmt.Println("Failed to set expiration for key:", err)
			return user1.Username, user1.Password, err
		}*/
		return user1.Username, user1.Password, nil
	}
	return userData[0], userData[1], nil
}
func CheckUser(username1 string) bool {
	var userExists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?) AS user_exists"
	err := db.Get(&userExists, query, username1)

	if err != nil {
		log.Fatal(err)
	}
	return userExists
}

func AddUser(username1 string, password1 string) {
	sqlStr := "insert into users(username,password) values (?,?)"
	ret, err := db.Exec(sqlStr, username1, password1)
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
func InitRd() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		//Password: "123456",
		DB: 0,
	})

}
func InitDB() {
	var err error

	dsn := "root:1234@tcp(127.0.0.1:3306)/testDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	fmt.Println("connecting to MySQL...")
	return
}
func Update(id int, oldPassword string, newPassword string) error {
	_, realPassword, err := GetFromID(id)
	if err != nil {
		return err
	}
	if realPassword == oldPassword {
		sqlStr := "update users set password=? where id = ?"
		ret, err := db.Exec(sqlStr, newPassword, id)
		if err != nil {
			return err
		}
		n, err := ret.RowsAffected() // 操作影响的行数
		if err != nil {
			fmt.Printf("get RowsAffected failed, err:%v\n", err)
			return err
		}
		fmt.Printf("update success, affected rows:%d\n", n)
		ctx := context.Background()
		idStr := strconv.Itoa(id)
		err = Rdb.HDel(ctx, idStr, "username", "password").Err()
		if err != nil {
			fmt.Println("Failed to delete hash field:", err)
			return err
		}
	}
	return nil
}
