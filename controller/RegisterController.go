/*
@Project ：GolangProjects
@File    ：RegisterController.go
@IDE     ：GoLand
@Author  ：@我不是大佬
@Email   ：2869210303@qq.com
@Date    ：2023/9/30 13:51
*/
package controller

import (
	"bookStoreAdmin/Dao"
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func RegistController() {
	availabe, username, password := ShowRegisterMenu()
	if !availabe {
		fmt.Println("\n两次输入的密码不一致，无法注册")
	} else if Dao.IsUsernameExists(username) {
		fmt.Print("\n用户名：", username, " 已经存在，无法注册")
	} else if Dao.CreateUser(username, password) {
		fmt.Println("\n注册成功！注册用户名:", username, "  注册密码:", password)
	}
}
func ShowRegisterMenu() (bool, string, string) {
	fmt.Print(" 请输入需要注册的用户名：")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()
	fmt.Print(" 请输入密码：")
	password, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Print(" \n请再次输入密码：")
	confirmPassword, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	if string(password) == string(confirmPassword) {
		return true, username, string(password)
	}
	return false, username, string(password)
}
