package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/9/8 4:26 下午
 * @Desc:
 */

func loginHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "hello",
	})
	//context.HTML(http.StatusOK, "shopping/login.html", )
}

//func indexHandler(context *gin.Context) {
//	context.HTML(http.StatusOK, "index/index.html", gin.H{
//		"标题": "index",
//	})
//}

func main() {
	r := gin.Default()
	//加载模板文件
	//r.LoadHTMLGlob("templates/**/*")  相当于templates目录下的任意文件夹下的任意文件
	//r.LoadHTMLGlob("templates/*")		//可以设置html存放的文件夹
	r.LoadHTMLFiles("templates/login.html")  //只能设置一个一个的html文件
	//r.LoadHTMLGlob("templates/**/*")

	//设置静态文件的目录
	//第一个参数是代码里使用的路径，第二个参数是实际保存静态文件的路径
	r.Static("/static", "./static")
	r.GET("/login", loginHandler)
	r.Run(":8080")
}