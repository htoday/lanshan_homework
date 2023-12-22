package main

import (
	"redisBasic/api"
	"redisBasic/dao"
)

// Rdb
//go get -u gorm.io/gorm
//go get github.com/jmoiron/sqlx
//go get -u gorm.io/driver/mysql
//go get -u github.com/gin-gonic/gin
//go get -u github.com/go-redis/redis/v8/*

func main() {
	dao.InitRd() //初始化redis
	dao.InitDB() //初始化mysql
	api.InitRouter()
}
