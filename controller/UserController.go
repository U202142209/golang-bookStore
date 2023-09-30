/*
@Project ：GolangProjects
@File    ：UserController.go
@IDE     ：GoLand
@Author  ：@我不是大佬
@Email   ：2869210303@qq.com
@Date    ：2023/9/28 21:54
*/
package controller

import (
	"bookStoreAdmin/Dao"
	"bookStoreAdmin/util"
	"fmt"
)

func ShowAllUsers() {
	// 获取用户信息
	fmt.Print("\n已注册用户信息如下")
	fmt.Println("\n序号\tID\t用户名\t密码\t注册时间\t登录次数\t上次登录时间")
	users := Dao.GetAllUsers()
	for count, user := range users {
		fmt.Println(count+1, "\t", user.ID, "\t", user.Username, "\t", user.Password, "\t", util.FormatTime(user.Createtime), "\t", user.LoginCount, "\t", util.FormatTime(user.LastLoginTime))
	}
}
