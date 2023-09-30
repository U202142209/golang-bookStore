/*
@Project ：GolangProjects
@File    ：UserMapper.go
@IDE     ：GoLand
@Author  ：@我不是大佬
@Email   ：2869210303@qq.com
@Date    ：2023/9/28 22:43
*/
package Dao

import (
	"bookStoreAdmin/entity"
	"bookStoreAdmin/util"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func GetAllUsers() []entity.User {
	users := []entity.User{}
	DB := GetDBConnection()
	defer DB.Close()
	rows, err := DB.Query("SELECT * FROM bookusers")
	util.CheckErrorPanic(err)
	defer func() {
		if rows != nil {
			rows.Close() // 关闭未scan的sql连接
		}
	}()
	var (
		createTimeStr    string
		lastlogintimestr string
	)
	for rows.Next() {
		user := new(entity.User)
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &createTimeStr, &user.LoginCount, &lastlogintimestr)
		util.CheckErrorPanic(err)
		createTime, err := time.Parse("2006-01-02 15:04:05", createTimeStr)
		util.CheckErrorPanic(err)
		lastlOGINtIME, err := time.Parse("2006-01-02 15:04:05", lastlogintimestr)
		util.CheckErrorPanic(err)
		user.Createtime = createTime
		user.LastLoginTime = lastlOGINtIME
		users = append(users, *user)
	}
	return users
}

// 新增用户，注册
func CreateUser(username, password string) bool {
	DB := GetDBConnection()
	defer DB.Close()
	stmt, err := DB.Prepare("INSERT INTO bookusers(username,password) VALUES (?,?)")
	util.CheckErrorPanic(err)
	resilt, err := stmt.Exec(strings.TrimSpace(username), strings.TrimSpace(password))
	util.CheckErrorPanic(err)
	affectedRwos, _ := resilt.RowsAffected()
	return affectedRwos > 0
}

// 登录
func IsExists(username, password string) (bool, *entity.User) {
	DB := GetDBConnection()
	defer DB.Close()
	stmt, err := DB.Prepare("SELECT * FROM bookusers where username=? AND password = ?")
	util.CheckErrorPanic(err)
	rows, err := stmt.Query(strings.TrimSpace(username), strings.TrimSpace(password))
	util.CheckErrorPanic(err)
	user := new(entity.User)
	for rows.Next() {
		var (
			createTimeStr    string
			lastlogintimestr string
		)
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &createTimeStr, &user.LoginCount, &lastlogintimestr)
		util.CheckErrorPanic(err)
		createTime, err := time.Parse("2006-01-02 15:04:05", createTimeStr)
		util.CheckErrorPanic(err)
		lastlOGINtIME, err := time.Parse("2006-01-02 15:04:05", lastlogintimestr)
		util.CheckErrorPanic(err)
		user.Createtime = createTime
		user.LastLoginTime = lastlOGINtIME
		if AddLoginCount(user.ID, DB) {
			user.ID += 1
			return true, user
		} else {
			fmt.Println("增加login count失败")
		}
	}
	return false, nil
}

// 增加登录次数
func AddLoginCount(ID int, DB *sql.DB) bool {
	// 事务处理
	tx, err := DB.Begin()
	util.CheckErrorPanic(err)
	stmt, err := tx.Prepare("UPDATE bookusers SET logincount = logincount + 1 WHERE id = ?;")
	util.CheckErrorPanic(err)
	result, err := stmt.Exec(ID)
	util.CheckErrorPanic(err)
	count, _ := result.RowsAffected()
	if count > 0 {
		tx.Commit()
		return true
	}
	tx.Rollback()
	return false
}

// 判断用户名是否存在
func IsUsernameExists(username string) bool {
	DB := GetDBConnection()
	defer DB.Close()
	stmt, err := DB.Prepare("SELECT * FROM bookusers where username=?")
	util.CheckErrorPanic(err)
	rows, err := stmt.Query(strings.TrimSpace(username))
	util.CheckErrorPanic(err)
	for rows.Next() {
		return true
	}
	return false
}
