/*
@Project ：GolangProjects
@File    ：main.go
@IDE     ：GoLand
@Author  ：@我不是大佬
@Email   ：2869210303@qq.com
@Date    ：2023/9/28 14:00
*/
package main

import (
	"bookStoreAdmin/controller"
	"bookStoreAdmin/service"
	"fmt"
)

func main() {
	if service.SystemInit() {
		fmt.Print("----------欢迎使用bookStroe管理系统后台----------")
		for {
			var userChioce string = controller.ShowMainMenu()
			if userChioce == "1" {
				controller.LoginController()
			} else if userChioce == "2" {
				controller.RegistController()
			} else if userChioce == "3" {
				fmt.Println("退出系统成功，欢迎下次使用")
				break
			} else {
				fmt.Print("你输入的选项不正确，请重新选择")
			}
		}
	}
}
