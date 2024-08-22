package main

import "fmt"

// interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called ....")
	fmt.Println(arg)
}

type Book struct {
	auth string
}

func main() {
	book := Book{"Golang"}
	myFunc(book)
	myFunc(100)
	myFunc("abc")
	// try git stash
	myFunc(200)

}
