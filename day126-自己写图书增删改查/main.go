package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/9/8 9:21 下午
 * @Desc:
 */

func init() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>连接数据库")
	//程序启动就应该连接数据库
	err := initDB()
	if err != nil {
		panic(err)
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>数据库连接成功")

}

func main() {
	r := gin.Default()

	//查看所有书籍
	r.GET("/book/list", bookListHandler)

	_ = r.Run(":8080")
}


//查看所有书籍
func bookListHandler(context *gin.Context) {
	//1. 连数据库，
	//2. 查数据
	bookList, err := queryAllBook()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": err,
		})
		return
	}

	//3. 返回给浏览器
	context.JSON(http.StatusOK, gin.H{
		"code": 0,  //约定的状态码
		"data": bookList,
	})
}