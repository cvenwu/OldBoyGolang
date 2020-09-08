package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/9/8 5:30 下午
 * @Desc: 参数相关示例
 */

func queryStringHandler(c *gin.Context) {
	//获取query string参数
	name := c.Query("name") //查不到就是空字符串
	//如果我们想查不到的时候指定默认值可以使用DefaultQuery
	city := c.DefaultQuery("city", "北京") //查不到就用北京填充
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"city": city,
	})
}

func formHandler(c *gin.Context) {
	username := c.PostForm("username")
	pwd := c.DefaultPostForm("password", "123456")  //如果post请求没有传递password我们就设置值为123456
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"password": pwd,
	})
}

func main() {
	r := gin.Default()
	//query string 也就是url后面跟着的参数
	r.GET("/index", queryStringHandler)
	//form params: HTMl页面上form表单提交的数据
	r.POST("/form", formHandler)
	//提取url后面路径 例如：/book/list查看书籍，/book/new 新建一本书，/book/delete 删除一本书
	r.GET("/book/:action", paramsHandler)
	r.Run(":8080")
}

func paramsHandler(context *gin.Context) {
	//提取路径参数
	actionVal := context.Param("action")
	context.JSON(http.StatusOK, gin.H{
		"action": actionVal,
	})
}
