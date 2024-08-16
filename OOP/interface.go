package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("Color = ", animal.GetColor())
	fmt.Println("Type = ", animal.GetType())

}

func main() {
	var animal AnimalIF //接口的数据类型
	animal = &Cat{"Green"}
	animal.Sleep()

	animal = &Dog{"Yellow"}
	animal.Sleep()

	cat1 := Cat{"Red"}
	showAnimal(&cat1)
}
