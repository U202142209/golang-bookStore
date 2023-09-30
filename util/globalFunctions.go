/*
@Project ：GolangProjects
@File    ：globalFunctions.go
@IDE     ：GoLand
@Author  ：@我不是大佬
@Email   ：2869210303@qq.com
@Date    ：2023/9/30 14:00
*/
package util

import (
	"fmt"
	"time"
)

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	t := time.Date(2023, 9, 28, 22, 8, 38, 0, time.UTC)
	formattedTime := FormatTime(t)
	fmt.Println(formattedTime)
}

func QuitInput(string2 string) bool {
	return string2 == "Q" || string2 == "q"
}
