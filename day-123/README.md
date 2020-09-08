# gin框架的渲染

## json

`c.JSON(状态码, 能够被json序列化的数据gin.H{})`

## html

`c.HTML(状态码，模板文件，数据)`

### 加载模板文件的两个方法

1. `r.LoadHTMLGlob("templates/**/*")`：  相当于templates目录下的任意文件夹下的任意文件
2. `r.LoadHTMLGlob("templates/*")`		相当于存放于templates下的html文件，可以设置html存放的文件夹
3. `r.LoadHTMLFiles("templates/login.html")`  只能设置一个一个的html文件

### 加载静态文件

> html里用到的css,js,image等

1. ​	`r.Static("/static", "./static")` 第一个参数是代码里写的路径，后面那个是实际存放静态文件的路径。

## json渲染

### json序列化的两种方式

1. 使用内置的c.JSON
2. 使用内置的结构体

![kQHbdc](https://gitee.com/yirufeng/images/raw/master/uPic/kQHbdc.png)

## xml渲染

![NeD8Q8](https://gitee.com/yirufeng/images/raw/master/uPic/NeD8Q8.png)

结果：![f96mpP](https://gitee.com/yirufeng/images/raw/master/uPic/f96mpP.png)

## yaml渲染

![4dFQ6o](https://gitee.com/yirufeng/images/raw/master/uPic/4dFQ6o.png)

## protobuf渲染

Protobuf：做微服务用的二进制。

![zqoinD](https://gitee.com/yirufeng/images/raw/master/uPic/zqoinD.png)