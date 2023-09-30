/*
@Project ：GolangProjects
@File    ：Connection.go
@IDE     ：GoLand
@Author  ：@我不是大佬
@Email   ：2869210303@qq.com
@Date    ：2023/9/28 22:44
*/
package Dao

import (
	"bookStoreAdmin/config"
	"bookStoreAdmin/util"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func GetDBConnection() *sql.DB {
	// 连接数据库
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", config.USERNAME, config.PASSWORD, config.NETWORK, config.SERVER, config.PORT, config.DATABASE)
	DB, err := sql.Open("mysql", conn)
	util.CheckErrorPanic(err)
	// 设置最大连接周期。超时就close()
	DB.SetConnMaxLifetime(100 * time.Second)
	// 设置最大连接数
	DB.SetMaxOpenConns(100)
	return DB
}

func Migrate(DB *sql.DB) bool {
	var createTablesSql = `CREATE TABLE IF NOT EXISTS bookusers (
	  id INT AUTO_INCREMENT PRIMARY KEY,
	  username VARCHAR(255) NOT NULL,
	  password VARCHAR(255) NOT NULL,
	  createtime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  logincount INT UNSIGNED DEFAULT 1,
	  lastlogintime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`

	var createBooksTableSql = `CREATE TABLE IF NOT EXISTS books (
	  id INT AUTO_INCREMENT PRIMARY KEY,
	  bookname CHAR(100) NOT NULL,
	  detail TEXT,
	  price INT UNSIGNED,
	  createtime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  lastedittime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`

	if _, err := DB.Exec(createTablesSql); err != nil {
		fmt.Println("bookusers表创建失败：", err)
		return false
	}
	if _, err := DB.Exec(createBooksTableSql); err != nil {
		fmt.Println("books表创建失败：", err)
		return false
	}
	fmt.Println("数据表创建成功！")
	return true
}
