package main

import "fmt"

func myMap() {
	// key是string， value也是string
	var myMap1 map[string]string

	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}

	//在使用map前，需要先用make给map分配数据空间
	myMap1 = make(map[string]string, 10)

	// map
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	fmt.Println("myMap2 : ", myMap2)

	//myMap3
	myMap3 := map[string]string{
		"1": "php",
		"2": "java",
		"3": "python",
	}
	fmt.Println("myMap3 : ", myMap3)

	// cityMap是引用传递
	cityMap := make(map[string]string)
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}

	//删除
	delete(cityMap, "Japan")

	//修改
	cityMap["USA"] = "DC"

}
