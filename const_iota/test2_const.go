package main

import "fmt"

// const 来定义枚举类型
const (
	/* 可以在const() 添加一个关键字iota，每行的iota会自动加1
	第一行的iota默认为0
	BEIJING   = 0
	SHNAGHAI  = 1
	SHENAZHEN = 2

	*/

	BEIJING = iota
	SHANGHAI
	SHENAZHEN
)

// iota只能配合const() 一起实用
const (
	// iota 从0开始，每行自动加1
	a, b = iota + 1, iota + 2 // iota =0, a= iota+1= 1, b =iota+2 = 2
	c, d
	e, f

	g, h = iota * 2, iota * 3 // iota =3, g = iota*2 = 6,h=iota*3  9
)

func main() {
	const length int = 10
	fmt.Println("length = ", length)
	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)

}
