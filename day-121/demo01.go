package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/9/8 3:18 下午
 * @Desc:
 */

func main() {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"姓名": "yirufeng",
			"年龄": "22",
		})
	})
	r.Run(":8080")
}
