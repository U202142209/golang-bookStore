/*
@Project ：GolangProjects
@File    ：book.go
@IDE     ：GoLand
@Author  ：@我不是大佬
@Email   ：2869210303@qq.com
@Date    ：2023/9/28 17:45
*/
package entity

import "time"

type Book struct {
	ID           int       `json:"id"`
	Bookname     string    `json:"name"`
	Detail       string    `json:"detail"`
	Price        int       `json:"price"`
	Createtime   time.Time `json:"createtime"`
	Lastedittime time.Time `json:"lastedittime"`
}
