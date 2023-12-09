package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	middleware "webBasic/Middleware"
	"webBasic/dao"
	"webBasic/model"
	"webBasic/utils"
)

func InitRouter() {
	r := gin.Default()
	//r.Use(middleware.CORS())
	r2 := r.Group("/MsgBroad")
	{
		r2.GET("/message", getMessages)
		r2.POST("/message", addMessages)
	}
	r.POST("/register", register) // 注册
	r.POST("/login", login)       // 登录
	r.POST("/changePassword", changePassword)
	r.POST("/findPassword", findPassword)
	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}

	r.Run(":8008") // 跑在 8008 端口上
}
func addMessages(c *gin.Context) {
	var newMessage dao.Message
	if err := c.ShouldBindJSON(&newMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newMessage.Timestamp = time.Now()
	dao.Messages = append(dao.Messages, newMessage)
	dao.WriteMessagesToFile()
	c.JSON(http.StatusCreated, newMessage)
}

func getMessages(c *gin.Context) {
	c.JSON(http.StatusOK, dao.Messages)
}
func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}
func findPassword(c *gin.Context) {
	username := c.PostForm("username")
	flag := dao.CheckUser(username)
	if flag {
		password := dao.FindPasswordFormUserName(username)
		c.JSON(200, gin.H{
			"status":   120,
			"message":  "password is found",
			"password": password,
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  500,
		"message": "username is absent",
	})
}
func register(c *gin.Context) {
	// 传入用户名和密码
	username := c.DefaultPostForm("username", "LZH")
	password := c.PostForm("password")
	if username == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  300,
			"message": "username is blank",
		})
		return
	}
	// 验证用户名是否重复
	flag := dao.CheckUser(username)
	// 重复则退出
	if flag {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "username is existed",
		})
	}

	dao.AddUser(username, password)
	dao.SaveToJSON("dataHome", dao.Database) //数据保存
	// 以 JSON 格式返回信息
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "add user successful",
	})
}
func changePassword(c *gin.Context) {
	username := c.PostForm("username")
	rawPassword := c.PostForm("rawPassword")
	newPassword := c.PostForm("newPassword")
	// 验证用户名是否存在
	flag := dao.CheckUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}
	// 查找正确的密码
	selectPassword := dao.FindPasswordFormUserName(username)
	// 若不正确则传出错误
	if selectPassword != rawPassword {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "wrong password")
		return
	}
	dao.ChangePassword(username, newPassword)
	dao.SaveToJSON("dataHome", dao.Database)
}
func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 验证用户名是否存在
	flag := dao.CheckUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}
	// 查找正确的密码
	selectPassword := dao.FindPasswordFormUserName(username)
	// 若不正确则传出错误
	if selectPassword != password {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "wrong password")
		return
	}
	// 正确则登录成功
	// 创建一个我们自己的声明
	claim := model.MyClaims{
		Username: username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
			Issuer:    "LZH",                                // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.RespSuccess(c, tokenString)
}
