package utils

import "github.com/gin-gonic/gin"

// RespSuccess 返回成功的 JSON 响应
func RespSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"status": "success",
		"data":   data,
	})
}

// RespFail 返回失败的 JSON 响应
func RespFail(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"status":  "fail",
		"message": message,
	})
}
