/*
@Project ：GolangProjects
@File    ：BookController.go
@IDE     ：GoLand
@Author  ：@我不是大佬
@Email   ：2869210303@qq.com
@Date    ：2023/9/28 21:55
*/
package controller

import (
	"bookStoreAdmin/Dao"
	"bookStoreAdmin/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CreateNewBook() {
	scanner := bufio.NewScanner(os.Stdin)
	var bookname string
	var bookdetail string
	var bookprice int
	for {
		fmt.Print("\n请输入新增的书籍名称:(Q退出)")
		scanner.Scan()
		bookname = scanner.Text()
		if util.QuitInput(bookname) {
			fmt.Print("取消新增书籍操作成功")
			return
		}
		if len(bookname) > 0 {
			break
		} else {
			fmt.Print("\n出错了，书籍名称不能为空")
		}
	}
	for {
		fmt.Print("\n请输入新增的书籍简介:(Q退出)")
		scanner.Scan()
		bookdetail = scanner.Text()
		if util.QuitInput(bookdetail) {
			fmt.Print("取消新增书籍操作成功")
			return
		}
		if len(bookdetail) > 0 {
			break
		} else {
			fmt.Print("\n出错了，书籍简介不能为空")
		}
	}
	for {
		fmt.Print("\n请输入新增的书籍价格:(Q退出)")
		scanner.Scan()
		input := scanner.Text()
		if util.QuitInput(input) {
			fmt.Print("取消新增书籍操作成功")
			return
		}
		var err error
		bookprice, err = strconv.Atoi(input)
		if err != nil {
			fmt.Print("\n输入的书籍价格必须为整数")
		} else {
			if bookprice <= 0 {
				fmt.Print("\n书籍的价格必须大于0")
			} else {
				break
			}
		}
	}
	if Dao.CreateNewBook(bookname, bookdetail, bookprice) {
		fmt.Print("新增书籍成功，书籍名称:", bookname, " 书籍简介:", bookdetail, " 书籍价格:", bookprice)
	}
}

func ShowAllBooks() {
	// 获取书籍信息
	fmt.Print("\n已添加的书籍息如下")
	// fmt.Println("\n序号 \tID \t书籍名称 \t书籍简介 \t书籍价格 \t新增时间 \t上次编辑时间")
	books := Dao.GetAllBooks()
	for count, book := range books {
		fmt.Println("\n序号:", count+1, " 书籍ID:", book.ID, " 书籍名称:《", book.Bookname, "》 书籍价格:", book.Price, "\n 书籍简介:", book.Detail, "\n 添加时间:", util.FormatTime(book.Createtime), " 上次编辑时间:", util.FormatTime(book.Lastedittime))
	}
}

func ChangeBookData() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\n请输入需要修改的书籍ID:")
	scanner.Scan()
	changeId, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Print("\n  出错了，输入的书籍ID必须为整数")
		return
	}
	if changeId <= 0 {
		fmt.Print("\n  出错了，书籍ID必须大于0")
		return
	}
	currentBook := Dao.GetBookById(changeId)
	if currentBook == nil {
		fmt.Print("\n  出错了，输入的书籍ID无效，没有查询到指定的书籍")
		return
	}
	fmt.Print("\n当前被修改的书籍信息如下：书籍名称：<", currentBook.Bookname, "> 书籍简介：", currentBook.Detail, " 书籍价格：", currentBook.Price)
	fmt.Print("\n\t请输入修改后的书籍名称:(Q退出)")
	scanner.Scan()
	bookname := scanner.Text()
	if util.QuitInput(bookname) {
		fmt.Print("取消修改书籍操作成功")
		return
	}
	if len(bookname) <= 0 {
		fmt.Print("\n出错了，书籍名称不能为空")
		return
	}
	fmt.Print("\t请输入修改后的书籍简介:(Q退出)")
	scanner.Scan()
	bookdetail := scanner.Text()
	if util.QuitInput(bookdetail) {
		fmt.Print("取消新增书籍操作成功")
		return
	}
	if len(bookdetail) <= 0 {
		fmt.Print("\n出错了，书籍简介不能为空")
		return
	}
	fmt.Print("\t请输入修改后的书籍价格:(Q退出)")
	scanner.Scan()
	input := scanner.Text()
	if util.QuitInput(input) {
		fmt.Print("取消新增书籍操作成功")
		return
	}
	bookprice, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print("\n输入的书籍价格必须为整数")
		return
	}
	if bookprice <= 0 {
		fmt.Print("\n书籍的价格必须大于0")
		return
	}
	// 修改书籍
	if Dao.ChangeBook(bookname, bookdetail, bookprice, changeId) {
		fmt.Print("修改数据信息成功！")
	} else {
		fmt.Print("出错了，修改数据信息失败，没有数据受到影响")
	}
}

func deleteBook() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\n请输入需要删除的书籍ID:")
	scanner.Scan()
	deleteId, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Print("\n  出错了，输入的书籍ID必须为整数")
		return
	}
	if deleteId <= 0 {
		fmt.Print("\n  出错了，书籍ID必须大于0")
		return
	}
	if Dao.GetBookById(deleteId) == nil {
		fmt.Print("\n  出错了，书籍ID：", deleteId, "不存在数据库中")
	}
	if Dao.DeleteBookByBookId(deleteId) {
		fmt.Print("\n删除书籍信息成功！")
	} else {
		fmt.Print("出错了，删除书籍信息失败，没有数据受到影响")
	}
}
