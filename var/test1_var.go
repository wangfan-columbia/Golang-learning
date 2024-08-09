// 四种变量的声明
package main

import (
	"fmt"
)

// 声明全局变量 方法一，二，三 是可以的
var gA int = 1
var gB = 2

// 方法四不行， := 只能在函数内，不能声明全局变量

func main() {
	// 方法一： 声明一个变量，默认是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	// 方法二： 声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("type of bb = %T\n", bb)
	// 方法三： 省去数据类型，通过值自动匹配
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	//方法四：省略var，自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("Type of e = %T", e)
	//fmt.Println("e = ", e)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx= ", xx, ", yy= ", yy)

	var kk, ll = 10, "aced"
	fmt.Println("kk= ", kk, ", ll= ", ll)

	var (
		nn = 123
		cc = "string"
	)
	fmt.Println("nn= ", nn, ", cc= ", cc)

}
