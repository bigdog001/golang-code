package main

import (
	"fmt"
)

type Person struct{
	Name string
}
// 给Person绑定一个方法
func (p Person)doing(){
	fmt.Println("名字是",p.Name);
}
// 给Person绑定一个方法
func (p Person)calculate(a int,b int){
	res := a+b
	fmt.Println("和是",res);
}

func main(){
	// 方法是作用在指定的数据类型上面的! 自定义类型(type)也可以有方法,而不仅仅是结构体才能有方法!
	//func (p Person)doing()表示Person结构体有一个方法doing.
	var p Person
	p.doing()
	p.calculate(5,10)
}