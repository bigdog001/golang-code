package main

import (
	"fmt"
	"strconv"
)

func main(){
	


	var c int = 32
	// var d float32 = 30.0
	// var e bool = false
	// var f byte = 'h'

	// 调用api将基本类型转换成字符串
	var str string = strconv.FormatInt(int64(c),10)
	var str1 string = strconv.FormatFloat(float64(c),'f',10,64)
	
	fmt.Println("",str)
	fmt.Println("",str1)
	
	// strconv.itoa
	
	

	
	
}