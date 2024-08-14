package main

import "fmt"

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func main() {
	// 类似于stack 按顺序加入栈，在执行
	defer func1()
	defer func2()
	defer func3()

}
