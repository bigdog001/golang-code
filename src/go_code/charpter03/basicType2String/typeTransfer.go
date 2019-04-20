package main

import (
	"fmt"
)

func main(){
	
	// 必须显示的转换
	var a int = 32
	var b float32 = float32(a)
	fmt.Printf("b=%f",b)

	var c int = 32
	// var d float32 = 30.0
	var e bool = false
	// var f byte = 'h'

	// 调用api将基本类型转换成字符串
	var str string = fmt.Sprintf("%d",c)
	var str1 string = fmt.Sprintf("%t",e)
	
	fmt.Println("",str)
	fmt.Println("",str1)
	
	
	

	
	
}