package main

//golang操作mysql进行单行和多行查询

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
