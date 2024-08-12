package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100

	return c
}

// 返回多个返回值，匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777
}

// 返回多个返回值，有形参名称
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("-------function 3--------")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = 100
	r2 = 200
	return
}

// 返回多个返回值，有形参名称
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("-------function 4--------")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = 100
	r2 = 200
	return
}

func main() {
	c := foo1("abc", 55)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("hhh", 999)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)

	r1, r2 := foo3("foo3", 333)
	fmt.Println("ret1 = ", r1, "ret2 = ", r2)

	r3, r4 := foo4("foo4", 444)
	fmt.Println("r3 = ", r3, "r4 = ", r4)
}
