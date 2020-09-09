package main

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/9/8 9:32 下午
 * @Desc: 专门用来与数据库中数据对应的结构体
 */

type Book struct {
	Id    int     `db:"id"`
	Title string  `db:"title"`
	Price float64 `db:"price"`
}
