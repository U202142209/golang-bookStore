/*
@Project ：GolangProjects
@File    ：BookMapper.go
@IDE     ：GoLand
@Author  ：@我不是大佬
@Email   ：2869210303@qq.com
@Date    ：2023/9/28 22:43
*/
package Dao

import (
	"bookStoreAdmin/entity"
	"bookStoreAdmin/util"
	"strings"
	"time"
)

func CreateNewBook(Bookname, Detail string, Price int) bool {
	DB := GetDBConnection()
	defer DB.Close()
	stmt, err := DB.Prepare("INSERT INTO books(bookname,detail,price) VALUES (?,?,?)")
	util.CheckErrorPanic(err)
	resilt, err := stmt.Exec(strings.TrimSpace(Bookname), strings.TrimSpace(Detail), Price)
	util.CheckErrorPanic(err)
	affectedRwos, _ := resilt.RowsAffected()
	return affectedRwos > 0
}

func GetAllBooks() []entity.Book {
	books := []entity.Book{}
	DB := GetDBConnection()
	defer DB.Close()
	rows, err := DB.Query("SELECT * FROM books")
	util.CheckErrorPanic(err)
	defer func() {
		if rows != nil {
			rows.Close() // 关闭未scan的sql连接
		}
	}()
	var (
		createTimeStr   string
		lastedittimestr string
	)
	for rows.Next() {
		book := new(entity.Book)
		err := rows.Scan(&book.ID, &book.Bookname, &book.Detail, &book.Price, &createTimeStr, &lastedittimestr)
		util.CheckErrorPanic(err)
		createTime, err := time.Parse("2006-01-02 15:04:05", createTimeStr)
		util.CheckErrorPanic(err)
		lastlOGINtIME, err := time.Parse("2006-01-02 15:04:05", lastedittimestr)
		util.CheckErrorPanic(err)
		book.Createtime = createTime
		book.Lastedittime = lastlOGINtIME
		books = append(books, *book)
	}
	return books
}

func GetBookById(id int) *entity.Book {
	DB := GetDBConnection()
	defer DB.Close()
	stmt, err := DB.Prepare("SELECT * FROM books where id = ?")
	util.CheckErrorPanic(err)
	rows, err := stmt.Query(id)
	util.CheckErrorPanic(err)
	defer func() {
		if rows != nil {
			rows.Close() // 关闭未scan的sql连接
		}
	}()
	var (
		createTimeStr   string
		lastedittimestr string
	)
	for rows.Next() {
		book := new(entity.Book)
		err := rows.Scan(&book.ID, &book.Bookname, &book.Detail, &book.Price, &createTimeStr, &lastedittimestr)
		util.CheckErrorPanic(err)
		createTime, err := time.Parse("2006-01-02 15:04:05", createTimeStr)
		util.CheckErrorPanic(err)
		lastlOGINtIME, err := time.Parse("2006-01-02 15:04:05", lastedittimestr)
		util.CheckErrorPanic(err)
		book.Createtime = createTime
		book.Lastedittime = lastlOGINtIME
		return book
	}
	return nil
}

func ChangeBook(bookname, bookdetail string, price, id int) bool {
	DB := GetDBConnection()
	defer DB.Close()
	stmt, err := DB.Prepare("UPDATE books SET bookname=? , detail=? ,price=? where id=?")
	util.CheckErrorPanic(err)
	result, err := stmt.Exec(bookname, bookdetail, price, id)
	util.CheckErrorPanic(err)
	count, err := result.RowsAffected()
	util.CheckErrorPanic(err)
	return count > 0
}

func DeleteBookByBookId(id int) bool {
	DB := GetDBConnection()
	defer DB.Close()
	stmt, err := DB.Prepare("DELETE FROM books where id-?")
	util.CheckErrorPanic(err)
	result, err := stmt.Exec(id)
	util.CheckErrorPanic(err)
	count, err := result.RowsAffected()
	util.CheckErrorPanic(err)
	return count > 0
}
