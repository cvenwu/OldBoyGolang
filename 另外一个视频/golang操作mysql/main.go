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
