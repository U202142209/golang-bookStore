/*
@Project ：GolangProjects
@File    ：errorUtil.go
@IDE     ：GoLand
@Author  ：@我不是大佬
@Email   ：2869210303@qq.com
@Date    ：2023/9/28 17:54
*/

package util

import "fmt"

func CheckError(err error) bool {
	if err != nil {
		fmt.Println("发生了错误:", err)
		return false
	}
	return true
}

func CheckErrorPanic(err error) {
	if err != nil {
		fmt.Println("发生了错误:", err)
		panic(err)
	}
}
