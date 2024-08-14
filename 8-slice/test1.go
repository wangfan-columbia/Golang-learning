package main

import "fmt"

// 值拷贝
func printArray1(myArray [4]int) {

	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}

}

// 引用传递， 不同元素长度的动态数组。他们的形参一致
func printArrayDynamic(myArray []int) {
	// _ 表示匿名的变量
	for _, value := range myArray {
		fmt.Println(" value = ", value)
	}

	myArray[0] = 100

}

func main() {
	//固定长度的数据
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}

	myArray3 := [4]int{11, 22, 33, 44}

	for i := 10; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	//查看数组的数据类型
	fmt.Printf("myArray1 type = %T\n", myArray1)
	fmt.Printf("myArray2 type = %T\n", myArray2)
	fmt.Printf("myArray3 type = %T\n", myArray3)

	// 动态数组，切片 slice
	myArray4 := []int{11, 22, 33, 44}
	printArrayDynamic(myArray4)
}
