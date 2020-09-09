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
