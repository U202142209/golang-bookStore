/*
@Project ：GolangProjects
@File    ：AdminController.go
@IDE     ：GoLand
@Author  ：@我不是大佬
@Email   ：2869210303@qq.com
@Date    ：2023/9/30 13:47
*/
package controller

import (
	"bufio"
	"fmt"
	"os"
)

func AdminController() {
	for {
		adminChioce := ShowAdminMenu()
		if adminChioce == "1" {
			ShowAllUsers()
		} else if adminChioce == "2" {
			ShowAllBooks()
		} else if adminChioce == "3" {
			CreateNewBook()
		} else if adminChioce == "4" {
			ChangeBookData()
		} else if adminChioce == "5" {
			deleteBook()
		} else if adminChioce == "6" {
			fmt.Print("\n退出登录成功！")
			break
		} else {
			fmt.Print("你输入的选项不正确，请重新选择")
		}
	}
}

func ShowAdminMenu() string {
	fmt.Print("\n 1.查看所有用户信息\n 2.查看所有书籍信息\n 3.新增书籍信息\n 4.修改书籍信息\n 5.删除指定书籍信息\n 6.退出登录\n情选择您的操作:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
