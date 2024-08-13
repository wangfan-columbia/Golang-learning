package lib2

import "fmt"

// 函数首字母大写 对外开放
func Lib2Test() {
	fmt.Println("lib2Test()...")
}

func init() {
	fmt.Println("lib2. init()...")
}
