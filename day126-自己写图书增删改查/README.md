




## 建立数据表

```sql
create database go_test;
use go_test;
CREATE TABLE book(
	id BIGINT(20) auto_increment PRIMARY KEY,
	title VARCHAR(20) not null,
	price DOUBLE(10,2) UNSIGNED
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4
```


## 

安装sqlx：`go get github.com/jmoiron/sqlx`