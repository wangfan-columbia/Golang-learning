package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human eat...")
}

func (this *Human) Walk() {
	fmt.Println("Human walk...")
}

type SuperMan struct {
	Human
	level int
}

// 重定义父类的方法
func (this *SuperMan) Eat() {
	fmt.Println("Super Man eat...")
}

// 子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("Super Man fly...")
}

func main() {
	h := Human{"fan", "m"}
	h.Eat()
	h.Walk()

	s := SuperMan{Human{"li4", "female"}, 88}
	s.Fly()
	s.Walk()

}
