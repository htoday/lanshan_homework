package api

import (
	"github.com/gin-gonic/gin"
	"redisBasic/dao"
	"strconv"
)

func InitRouter() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "helloWorld",
		})
	}) //get测试,一般返回200为成功，404为未找到，500为服务器内部错误

	r.POST("/register", register) // 注册
	//r.POST("/login", login)       // 登录
	r.POST("/select", selectFromID)
	r.POST("/like", like)
	err := r.Run(":8088") //端口前的冒号不能省略
	if err != nil {
		println("Error:", err)
	}
}
func like(c *gin.Context) {
	idStr := c.PostForm("id")
	PassageIDStr := c.PostForm("passageID")
	err := dao.AddLike(idStr, PassageIDStr)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "succeed liking",
	})
}
func selectFromID(c *gin.Context) {
	idStr := c.PostForm("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot translate id into intValue",
		})
		return
	}
	username, password, err := dao.GetFromID(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to find",
			"Error":   err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message":  "Succeed to find",
		"username": username,
		"password": password,
	})
}
func register(c *gin.Context) {
	// 传入用户名和密码
	username := c.DefaultPostForm("username", "LZH")
	password := c.PostForm("password")
	if username == "" {
		c.JSON(500, gin.H{
			"message": "username is blank",
		})
		return
	}
	// 验证用户名是否重复
	flag := dao.CheckUser(username)
	// 重复则退出
	if flag {
		// 以 JSON 格式返回信息
		c.JSON(500, gin.H{
			"message": "username is existed",
		})
		return
	}
	dao.AddUser(username, password)
	c.JSON(200, gin.H{
		"message": "add user successful",
	})
}
