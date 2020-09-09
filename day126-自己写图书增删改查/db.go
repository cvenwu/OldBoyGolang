package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/9/8 9:23 下午
 * @Desc: 与数据库进行交互
 */

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:1018222wxw@tcp(127.0.0.1:3306)/go_test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	//最大连接数
	db.SetMaxOpenConns(100)
	//最大空闲连接数
	db.SetMaxIdleConns(16)
	return nil
}

//查所有书籍
func queryAllBook() (bookList []*Book, err error) {
	sqlStr := "select id, title, price from book;"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("查询所有书籍信息失败！")
		return
	}
	return
}
