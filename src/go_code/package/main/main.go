package main

import "fmt"
// 包路径默认是从src下面开始的
import "go_code/package/utils"


func main(){
	i := utils.Calculate(5,10)
	fmt.Printf("运算结果是%.2f",i)
}   