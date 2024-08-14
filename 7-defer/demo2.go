package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0

}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0

}

// return 后的语句先执行， defer的语句后执行
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	returnAndDefer()
}
