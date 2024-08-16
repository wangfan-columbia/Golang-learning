package main

import "fmt"

type book struct {
	name   string
	author string
}

// 方法名首字母大写，其他类可以访问，否则只能本包内使用
type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
	//fmt.Println("Name= " this.name)
}

func (this *Hero) Show() {
	fmt.Println("Hero = ", this.Name)
	fmt.Println("Hero's ad = ", this.Ad)
	fmt.Println("Hero's level = ", this.Level)
}

func main() {
	hero := Hero{Name: "fan", Ad: 100, Level: 2}
	hero.Show()
}
