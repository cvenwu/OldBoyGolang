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
