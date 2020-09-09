

# 另外一个Bilibili视频介绍
[参考](https://www.bilibili.com/video/BV14C4y147y8?p=130)

10-102

103-111

112-121

122-129

130-142: 128 139 140, 142没看

144-148

149-180

# 04 配置GOPATH


gopath表示我们写的go的项目存放路径，所有的项目代码都存放到gopath的src目录下

![1CZyAM](https://gitee.com/yirufeng/images/raw/master/uPic/1CZyAM.png)

## 130 看完

![1uH90U](https://gitee.com/yirufeng/images/raw/master/uPic/1uH90U.png)

## 131 看完

![osLaE7](https://gitee.com/yirufeng/images/raw/master/uPic/osLaE7.png)

![dJ1wv3](https://gitee.com/yirufeng/images/raw/master/uPic/dJ1wv3.png)

![PSkUPS](https://gitee.com/yirufeng/images/raw/master/uPic/PSkUPS.png)

![3sVe0I](https://gitee.com/yirufeng/images/raw/master/uPic/3sVe0I.png)

读未提交：一个事务读取了另一个事务还未提交的数据

## 132 Go操作Mysql

[参考老师博客](https://www.liwenzhou.com/posts/Go/go_mysql/)

![U89pXB](https://gitee.com/yirufeng/images/raw/master/uPic/U89pXB.png)

### database/sql

**go官方提供的一个原生的支持连接池，是并发安全的**

**标准库没有具体实现，只是列出了第三方库需要具体实现的内容**

### 下载驱动

下载依赖：` go get -u github.com/go-sql-driver/mysql`

![tnYxOa](https://gitee.com/yirufeng/images/raw/master/uPic/tnYxOa.png)

### 使用驱动

![dz7m3L](https://gitee.com/yirufeng/images/raw/master/uPic/dz7m3L.png)

```go
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //init()
	"log"
)

//go连接mysql示例

func main() {
	//打开数据库，返回一个数据库连接池，因为建立连接很耗时
	//指定数据源的名字
	dsn := "root:1018222wxw@tcp(127.0.0.1:3306)/myemployees"

	//连接数据库
	//第一个参数指定使用哪个数据库驱动，如mysql oracle
	//如何使用第1个字符串去找到对应的数据库校验dsn呢？因为我们导入mysql的时候执行对应的Init函数，里面注册了mysql
	db, err := sql.Open("mysql", dsn)  //这里只是检测数据源的格式是否正确，并不会校验用户名密码是否正确
	if err != nil {
		log.Println(err)
		log.Fatalln("数据源dsn格式不正确！！！")
	}

	//尝试与数据库进行连接
	err = db.Ping()
	if err != nil {
		log.Println(err)
		log.Fatalln("数据库连接失败")
	}

	log.Println("数据库连接成功--------------------------")

	db.Close()

}

```

### 查询单条记录以及查询sql中带参数的单条记录

注意点1：![j5wbd0](https://gitee.com/yirufeng/images/raw/master/uPic/j5wbd0.png)

注意点2：![nQEcuG](https://gitee.com/yirufeng/images/raw/master/uPic/nQEcuG.png)

```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}


func init() {
	dsn := "root:1018222wxw@tcp(127.0.0.1:3306)/sql_test"

	//校验dsn格式是否正确
	//注意：这里不能使用:=，因为db是全局变量，如果使用:=将会成为一个局部变量
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("dsn格式不正确")
		log.Fatalln(err)
	}

	//测试连接是否通
	err = db.Ping()
	if err != nil {
		log.Println("建立连接失败！！")
		log.Fatalln(err)
	}

	//设置数据库的最大连接数
	db.SetMaxOpenConns(20)
	//设置数据库的最大空闲连接数
	db.SetMaxIdleConns(20)

	log.Println("建立连接成功----------------------->")
}

func main() {
	//defer db.Close()
	fmt.Println("123")

	//1.查询单条记录
	sqlStr := `select id, name, age from user where id = 1`
	//2.执行
	rowObj := db.QueryRow(sqlStr) //其实就是从连接池中拿到一个连接去数据库查询单条记录
	//3. 拿到结果
	var u user
	rowObj.Scan(&u.id, &u.name, &u.age)
	//4. 打印结果
	fmt.Println(u.age, u.name)


	//带参数的sql语句
	//1.查询单条记录
	sqlStr = `select id, name, age from user where id = ?`
	//2.执行
	rowObj = db.QueryRow(sqlStr, 2) //其实就是从连接池中拿到一个连接去数据库查询单条记录
	//3. 拿到结果,注意这里要修改user对象必须传入指针，同时我们必须对rowObj调用scan方法，因为会释放数据库连接到连接池中
	rowObj.Scan(&u.id, &u.name, &u.age)
	//4. 打印结果
	fmt.Println(u.age, u.name)
}

```

![ZRSRay](https://gitee.com/yirufeng/images/raw/master/uPic/ZRSRay.png)

代码的改进：

```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func init() {
	dsn := "root:1018222wxw@tcp(127.0.0.1:3306)/sql_test"

	//校验dsn格式是否正确
	//注意：这里不能使用:=，因为db是全局变量，如果使用:=将会成为一个局部变量
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("dsn格式不正确")
		log.Fatalln(err)
	}

	//测试连接是否通
	err = db.Ping()
	if err != nil {
		log.Println("建立连接失败！！")
		log.Fatalln(err)
	}

	//设置数据库的最大连接数
	db.SetMaxOpenConns(20)
	//设置数据库的最大空闲连接数
	db.SetMaxIdleConns(20)

	log.Println("建立连接成功----------------------->")
}

//查询单个记录
func queryOne(id int) {
	//1.查询单条记录
	sqlStr := `select id, name, age from user where id = ?`
	//2.执行
	var u user
	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age) //其实就是从连接池中拿到一个连接去数据库查询单条记录
	if err != nil {
		log.Println("查询失败------------>", err)
	}
	//3. 拿到结果
	log.Println("查询成功：--------->", u.age, u.name)
}

func main() {
	//defer db.Close()
	fmt.Println("123")
	queryOne(2)
	queryOne(123)
}

```

## 134 查询多条记录



```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func init() {
	dsn := "root:1018222wxw@tcp(127.0.0.1:3306)/sql_test"

	//校验dsn格式是否正确
	//注意：这里不能使用:=，因为db是全局变量，如果使用:=将会成为一个局部变量
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("dsn格式不正确")
		log.Fatalln(err)
	}

	//测试连接是否通
	err = db.Ping()
	if err != nil {
		log.Println("建立连接失败！！")
		log.Fatalln(err)
	}

	//设置数据库的最大连接数
	db.SetMaxOpenConns(20)
	//设置数据库的最大空闲连接数
	db.SetMaxIdleConns(20)

	log.Println("建立连接成功----------------------->")
}

//查询单个记录
func queryOne(id int) {
	//1.查询单条记录
	sqlStr := `select id, name, age from user where id = ?`
	//2.执行
	var u user
	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age) //其实就是从连接池中拿到一个连接去数据库查询单条记录
	if err != nil {
		log.Println("查询失败------------>", err)
		return
	}
	//3. 拿到结果
	log.Println("查询成功：--------->", u.age, u.name)
}

//查询多条记录
func queryMore(id int) {
	//1. sql语句
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		log.Println("多行查询失败------------->", err)
		return
	}
	defer rows.Close()
	//循环取值
	var u user
	for rows.Next() {
		err = rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			log.Println("多行查询失败----------->", err)
		}
		log.Println(u)
	}
}

func main() {
	//defer db.Close()
	fmt.Println("123")
	queryOne(2)
	queryOne(123)
	queryMore(1)
}

```

## 135 插入更新和删除数据

### 插入操作

通过`db.Exec(命令，参数)`来执行对应的插入sql语句，注意我们可以根据返回的第1个参数来获取**最后一条插入记录的id值**以及**影响的行数**





### 插入，删除，更新对应的代码

```go
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//golang操作mysql进行增加和删除

var db *sql.DB

type User struct {
	id   int
	name string
	age  int
}

func init() {
	var err error
	//指定数据源
	dataSource := "root:1018222wxw@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatalln("数据库参数检验失败-------------")
	}

	//尝试和数据库通信
	err = db.Ping()
	if err != nil {
		log.Fatalln("数据库连接错误---------------")
	}

	log.Println("数据库连接成功--------------------")
}

func queryOne(id int) {
	sqlStr := `select id, name, age from user where id = ?`
	var user User
	err := db.QueryRow(sqlStr, id).Scan(&user.id, &user.name, &user.age)
	if err != nil {
		log.Fatalln("查询失败--------------------》")
	}
	log.Println("查询成功-----------------------》", user.id, user.name, user.age)
}

func queryMore(id int) {
	sqlStr := `select id, name, age from user where id > ?`
	var user User
	ret, err := db.Query(sqlStr, id)
	defer ret.Close()
	if err != nil {
		log.Fatalln("查询失败--------------------》")
	}

	for ret.Next() {
		err = ret.Scan(&user.id, &user.name, &user.age)
		if err != nil {
			log.Fatalln("多条查询失败------------------")
		}
		log.Println("查询成功-----------------------》", user.id, user.name, user.age)
	}

}

//插入数据操作
func insert() {
	//1.写sql语句
	sqlStr := `insert into user(name,age) values(?, ?)`
	//2.exec执行sql语句
	ret, err := db.Exec(sqlStr, "吴晓文", "22")
	if err != nil {
		log.Println("插入失败-----------------", ret)
	}
	//3. 如果是插入数据的操作，能够拿到插入数据的id值
	id, err := ret.LastInsertId()
	if err != nil {
		log.Println("get id failed----------------")
	}
	log.Println("id：", id)
	//ret.RowsAffected() //影响的行数
}

//更新数据操作
func update() {
	//1.写sql语句
	sqlStr := `update  user set age = 13 where id = 7`
	//2.exec执行sql语句
	ret, err := db.Exec(sqlStr)
	if err != nil {
		log.Println("更新失败-----------------", ret)
	}
	//3. 如果是插入数据的操作，能够拿到插入数据的id值
	affectRows, err := ret.RowsAffected()
	if err != nil {
		log.Println("get id failed----------------")
	}
	log.Println("受影响的行数：", affectRows)
	//ret.RowsAffected() //影响的行数
}

//删除数据操作
func delete(id int) {
	//1.写sql语句
	sqlStr := `delete from user where id = ?`
	//2.exec执行sql语句
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		log.Println("删除失败-----------------", ret)
	}
	//3. 如果是插入数据的操作，能够拿到插入数据的id值
	affectRows, err := ret.RowsAffected()
	if err != nil {
		log.Println("delete failed----------------")
	}
	log.Println("受影响的行数：", affectRows)
	//ret.RowsAffected() //影响的行数
}

func main() {
	//insert()
	update()
	delete(7)
	queryOne(2)
	queryMore(1)
}

```



## 136 sql预处理

![HndvSE](https://gitee.com/yirufeng/images/raw/master/uPic/HndvSE.png)



如果我们要**大量执行同一个sql语句（也就是批量插入，更新，删除的时候）**，只是参数不同，我们可以让mysql对其进行预处理之后我们传递参数过去即可，这样**可以提高效率。同时还可以避免sql注入**

![jysbbP](https://gitee.com/yirufeng/images/raw/master/uPic/jysbbP.png)

![p3Ny2q](https://gitee.com/yirufeng/images/raw/master/uPic/p3Ny2q.png)

![Rkx1wQ](https://gitee.com/yirufeng/images/raw/master/uPic/Rkx1wQ.png)

![WfeAn7](https://gitee.com/yirufeng/images/raw/master/uPic/WfeAn7.png)







## 137 go语言实现事务

事务相关方法：![sPUOL2](https://gitee.com/yirufeng/images/raw/master/uPic/sPUOL2.png)

```go
package main

//golang操作事务

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//golang操作mysql进行增加和删除

var db *sql.DB

type User struct {
	id   int
	name string
	age  int
}

func init() {
	var err error
	//指定数据源
	dataSource := "root:1018222wxw@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatalln("数据库参数检验失败-------------")
	}

	//尝试和数据库通信
	err = db.Ping()
	if err != nil {
		log.Fatalln("数据库连接错误---------------")
	}

	log.Println("数据库连接成功--------------------")
}

func transactionDemo() {
	//1. 开启事务
	tx, err := db.Begin()
	if err != nil {
		log.Fatalln("事务开始失败------------------>", err)
	}

	//执行多个sql操作
	sqlStr := "update user set age = age - 2 where id = 2"
	sqlStr2 := "update e32 set age = age + 2 where id = 3"

	_, err = tx.Exec(sqlStr)
	if err != nil {
		// 要回滚
		tx.Rollback()
		log.Fatalln("语句1执行失败，要回滚-------------------")
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		// 要回滚
		tx.Rollback()
		log.Fatalln("语句2执行失败，要回滚-------------------")
	}
	err = tx.Commit()
	if err != nil {
		// 要回滚
		tx.Rollback()
		log.Println("事务提交失败，要回滚--------------------")
	}
	fmt.Println("事务执行成功----------------------")
}

func queryMore(id int) {
	sqlStr := `select id, name, age from user where id > ?`
	var user User
	ret, err := db.Query(sqlStr, id)
	defer ret.Close()
	if err != nil {
		log.Fatalln("查询失败--------------------》")
	}

	for ret.Next() {
		err = ret.Scan(&user.id, &user.name, &user.age)
		if err != nil {
			log.Fatalln("多条查询失败------------------")
		}
		log.Println("查询成功-----------------------》", user.id, user.name, user.age)
	}

}

func main() {
	fmt.Println("----------------------------------事务提交前----------------------------------")
	queryMore(1)
	transactionDemo()
	fmt.Println("----------------------------------事务提交后----------------------------------")
	queryMore(1)
}

```



## 138 sqlx的使用

优点：在查询出结果之后不需要一个一个字段的给变量赋值





## 139





## 140 



## 141 NSQ

Golang开发的一个分布式的消息队列

[参考老师博客](https://www.liwenzhou.com/posts/Go/go_nsq/)

![pbAxqu](https://gitee.com/yirufeng/images/raw/master/uPic/pbAxqu.png)



使用消息队列好处：为了将同步的操作变成异步的

常用消息队列：rabbitMQ



## 142 NSQ使用

### 参考

1. [官网](https://nsq.io/)

2. [老师博客](https://www.liwenzhou.com/posts/Go/go_nsq/)

