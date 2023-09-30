/*
@Project ：GolangProjects
@File    ：loginController.go
@IDE     ：GoLand
@Author  ：@我不是大佬
@Email   ：2869210303@qq.com
@Date    ：2023/9/28 14:07
*/
package controller

import (
	"bookStoreAdmin/Dao"
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func ShowLoginMenu() (string, string) {
	fmt.Print(" 请输入用户名：")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()
	fmt.Print(" 请输入密码：")
	//scanner.Scan()
	//password = scanner.Text()
	password, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	return username, string(password)
}

func ShowMainMenu() string {
	// 展示用户登录菜单，并获取用户的输入值
	fmt.Print("\n 1.登录\n 2.新用户注册\n 3.退出系统\n情选择您的操作:")
	scanner := bufio.NewScanner(os.Stdin)
	// 读取用户输入的字符串
	scanner.Scan()
	return scanner.Text()
}

func LoginController() {
	// 登录
	userInputUserName, userInputPassword := ShowLoginMenu()
	flag, _ := Dao.IsExists(userInputUserName, userInputPassword)
	if flag {
		fmt.Print("登录成功！")
		AdminController()
	} else {
		fmt.Print("\n用户名或密码不正确，登录失败")
	}
}
