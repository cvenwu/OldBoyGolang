
# gin框架基本示例

## beego与gin区别

1. beego主要用于公司内部系统开发，比如内部系统

2. gin框架主要应用于前后端分离开发，小程序，电商网站

[参考老师博客](https://www.liwenzhou.com/posts/Go/Gin_framework/)

![9LzB91](https://gitee.com/yirufeng/images/raw/master/uPic/9LzB91.png)

## gin下载与安装

下载gin框架：`go get -u github.com/gin-gonic/gin`

![zAWAYC](https://gitee.com/yirufeng/images/raw/master/uPic/zAWAYC.png)

## gin框架简单示例

```go
func indexHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK, gin.H{
			"msg": "这是index页面",
		})
}

func main() {
	//启动一个默认的路由
	router := gin.Default()
	//给/hello配置一个处理函数
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "你好中国",
		})
	})
	router.GET("/index", indexHandler)
	//启动webserver
	router.Run(":8080")
}

```
