/*
@Project ：GolangProjects
@File    ：user.go
@IDE     ：GoLand
@Author  ：@我不是大佬
@Email   ：2869210303@qq.com
@Date    ：2023/9/28 17:45
*/
package entity

import "time"

type User struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Createtime    time.Time `json:"createtime"`
	LoginCount    int       `json:"logincount"`
	LastLoginTime time.Time `json:"lastlogintime"`
}
