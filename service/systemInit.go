/*
@Project ：GolangProjects
@File    ：systemInit.go
@IDE     ：GoLand
@Author  ：@我不是大佬
@Email   ：2869210303@qq.com
@Date    ：2023/9/28 14:47
*/
package service

import (
	"bookStoreAdmin/Dao"
)

// 系统初始化
func SystemInit() (result bool) {
	DB := Dao.GetDBConnection()
	result = Dao.Migrate(DB)
	defer DB.Close()
	return result
}
