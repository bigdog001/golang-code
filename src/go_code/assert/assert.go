package main

import (
	"fmt"
	
)
type Point struct{
	x int
	y int
}

func main(){
		// 空接口
		var a interface{}
		var point Point = Point{1,2}
		a = point

		var b Point
		// b = a  //报错
		// 必须用类型断言,才不会报错
		b = a.(Point)
		fmt.Println("",b)

	}